package main

import (
	"context"
	"log"
	"runtime/debug"
	"strconv"
	"time"

	_ "github.com/golang/protobuf/ptypes/timestamp"
	_ "github.com/satori/go.uuid"
	"github.com/spf13/afero"
	"google.golang.org/grpc"
	pb "grpc-load-test/docs"

	"github.com/yandex/pandora/cli"
	"github.com/yandex/pandora/components/phttp/import"
	"github.com/yandex/pandora/core"
	"github.com/yandex/pandora/core/aggregator/netsample"
	"github.com/yandex/pandora/core/import"
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
	client grpc.ClientConn
	conf   GunConfig
	// Configured on Bind, before shooting
	aggr core.Aggregator // Maybe your custom Aggregator.
	core.GunDeps
}

func NewGun(conf GunConfig) *Gun {
	return &Gun{conf: conf}
}

func (g *Gun) Bind(aggr core.Aggregator, deps core.GunDeps) error {
	// create gRPC stub at gun initialization
	conn, err := grpc.Dial(
		g.conf.Target,
		grpc.WithInsecure(),
		grpc.WithTimeout(time.Second),
		grpc.WithUserAgent("load test, pandora custom shooter"))
	if err != nil {
		log.Fatalf("FATAL: %s", err)
	}
	g.client = *conn
	g.aggr = aggr
	g.GunDeps = deps
	return nil
}

func (g *Gun) Shoot(ammo core.Ammo) {
	customAmmo := ammo.(*Ammo)
	g.shoot(customAmmo)
}

func (g *Gun) caseSave(client pb.DocumentServiceClient, ammo *Ammo) int {
	code := 0
	docName := ammo.Param1

	out, err := client.Save(
		context.TODO(), &pb.SaveRequest{Document: createDoc(docName)},
	)

	if err != nil {
		log.Printf("FATAL: %s", err)
		code = 500
	}

	if out != nil {
		code = 200
	}
	return code
}

func (g *Gun) caseFindAll(client pb.DocumentServiceClient, ammo *Ammo) int {
	code := 0
	// prepare item_id and warehouse_id
	limit, err := strconv.ParseInt(ammo.Param1, 10, 0)
	if err != nil {
		log.Printf("Failed to parse ammo FATAL", err)
		code = 314
	}
	offset, err2 := strconv.ParseInt(ammo.Param2, 10, 0)
	if err2 != nil {
		log.Printf("Failed to parse ammo FATAL", err2)
		code = 314
	}

	out, err3 := client.GetAllByLimitAndOffset(
		context.TODO(), &pb.GetAllRequest{
			Limit:  int32(limit),
			Offset: int32(offset),
		})

	if err3 != nil {
		log.Printf("FATAL", err3)
		code = 316
	}

	if out != nil {
		code = 200
	}

	return code
}

func (g *Gun) shoot(ammo *Ammo) {
	code := 0
	sample := netsample.Acquire(ammo.Tag)

	client := pb.NewDocumentServiceClient(&g.client)

	switch ammo.Tag {
	case "/SaveCase":
		code = g.caseSave(client, ammo)
	case "/FindCase":
		code = g.caseFindAll(client, ammo)
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
	// Standard imports.
	fs := afero.NewOsFs()
	coreimport.Import(fs)
	// May not be imported, if you don't need http guns and etc.
	phttp.Import(fs)

	// Custom imports. Integrate your custom types into configuration system.
	coreimport.RegisterCustomJSONProvider("custom_provider", func() core.Ammo { return &Ammo{} })

	register.Gun("gprc_gun", NewGun, func() GunConfig {
		return GunConfig{
			Target: "grpc target",
		}
	})

	cli.Run()
}
