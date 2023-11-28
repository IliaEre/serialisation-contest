package main

import (
	"context"
	"errors"
	"flat-docs-service/pkg/db"
	"flat-docs-service/pkg/handler"
	"flat-docs-service/pkg/service"
	"fmt"
	"github.com/gin-gonic/gin"
	ginprometheus "github.com/zsais/go-gin-prometheus"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

const (
	post         = "/report"
	validate     = "/report/validate"
	get          = "/reports"
	TTL          = 5
	addr         = ":9093"
	subsystem    = "gin"
	mongoAddress = "mongodb://mongo:27017"
	collection   = "flatReports"
)

func main() {
	router := gin.Default()
	p := ginprometheus.NewPrometheus(subsystem)
	p.Use(router)

	mc, err := db.NewMongoRepository(mongoAddress, collection)
	if err != nil {
		log.Println("err", err)
		os.Exit(1)
	}
	sv := service.NewReportService(mc)
	gw := handler.NewHandler(sv)

	router.POST(post, gw.Save)
	router.POST(get, gw.FindByParams)
	router.POST(validate, gw.ValidateDepartment)

	srv := &http.Server{
		Addr:    addr,
		Handler: router,
	}

	go func() {
		// Service connections
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Print(err)
		}
	}()

	// Create a channel to listen for the OS interrupt signal
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	fmt.Println("Shutdown Server...")

	// Create a context with a timeout
	ctx, cancel := context.WithTimeout(context.Background(), TTL*time.Second)
	defer cancel()

	// Shutdown the server with the created context
	if err := srv.Shutdown(ctx); err != nil {
		panic(err)
	}

	fmt.Println("Server exiting")
}
