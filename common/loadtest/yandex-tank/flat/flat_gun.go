package main

import (
	"bytes"
	phttp "github.com/yandex/pandora/components/phttp/import"
	coreimport "github.com/yandex/pandora/core/import"
	"log"
	"net/http"
	"os"
	"runtime/debug"
	"time"

	"github.com/spf13/afero"

	"github.com/yandex/pandora/cli"
	"github.com/yandex/pandora/core"
	"github.com/yandex/pandora/core/aggregator/netsample"
	"github.com/yandex/pandora/core/register"
)

type Ammo struct {
	Tag    string
	Param1 string
	Param2 string
	Param3 string
}

type GunConfig struct {
	Target string `validate:"required"` // Configuration will fail, without target defined
}

type Gun struct {
	// Configured on construction.
	client http.Client
	conf   GunConfig
	// Configured on Bind, before shooting
	aggr core.Aggregator // Maybe your custom Aggregator.
	core.GunDeps
	files [][]byte
}

func NewGun(conf GunConfig) *Gun {
	return &Gun{conf: conf}
}

func (g *Gun) Bind(aggr core.Aggregator, deps core.GunDeps) error {
	c := http.Client{Timeout: time.Duration(1) * time.Second}
	g.client = c

	g.GunDeps = deps
	g.aggr = aggr

	saveFile, err := os.ReadFile("save.bin")
	if err != nil {
		log.Printf("FATAL: %s", err)
		panic(err)
	}

	findFile, err := os.ReadFile("request.binary")
	if err != nil {
		log.Printf("FATAL: %s", err)
		panic(err)
	}

	validateFile, err := os.ReadFile("validate.bin")
	if err != nil {
		log.Printf("FATAL: %s", err)
		panic(err)
	}

	g.files = append(g.files, saveFile, findFile, validateFile)
	return nil
}

func (g *Gun) Shoot(ammo core.Ammo) {
	customAmmo := ammo.(*Ammo)
	g.shoot(customAmmo)
}

func (g *Gun) caseSave(client *http.Client) int {
	host := g.conf.Target

	response, err := client.Post(host+"/report", "application/octet-stream", bytes.NewBuffer(g.files[0]))
	if err != nil {
		log.Printf("FATAL: %s", err)
		return 500
	}

	return response.StatusCode
}

func (g *Gun) caseFindAll(client *http.Client) int {
	host := g.conf.Target
	response, err := client.Post(host+"/reports", "application/octet-stream", bytes.NewBuffer(g.files[1]))
	if err != nil {
		log.Printf("FATAL: %s", err)
		return 500
	}

	return response.StatusCode
}

func (g *Gun) caseValidate(client *http.Client) int {
	host := g.conf.Target

	response, err := client.Post(host+"/report/validate", "application/octet-stream", bytes.NewBuffer(g.files[2]))
	if err != nil {
		log.Printf("FATAL: %s", err)
		return 500
	}

	return response.StatusCode
}

func (g *Gun) shoot(ammo *Ammo) {
	code := 0
	sample := netsample.Acquire(ammo.Tag)

	client := &g.client

	switch ammo.Tag {
	case "/SaveCase":
		code = g.caseSave(client)
	case "/FindCase":
		code = g.caseFindAll(client)
	case "/ValidateCase":
		code = g.caseValidate(client)
	default:
		code = 404
	}

	defer func() {
		sample.SetProtoCode(code)
		g.aggr.Report(sample)
	}()
}

func main() {
	debug.SetGCPercent(-1)
	fs := afero.NewOsFs()
	coreimport.Import(fs)
	phttp.Import(fs)

	coreimport.RegisterCustomJSONProvider("flat_provider", func() core.Ammo { return &Ammo{} })
	register.Gun("flat_gun", NewGun, func() GunConfig {
		return GunConfig{
			Target: "flat_target",
		}
	})

	cli.Run()
}
