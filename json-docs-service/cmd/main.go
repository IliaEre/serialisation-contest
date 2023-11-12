package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	ginprometheus "github.com/zsais/go-gin-prometheus"
	"json-docs-service/pkg/middle"
	"json-docs-service/pkg/service"
	"net/http"
	"os"
	"os/signal"
	"time"
)

const (
	post    = "/report"
	get     = "/reports"
	TTL     = 5
	address = ":9091"
)

func main() {
	router := gin.Default()
	p := ginprometheus.NewPrometheus("gin")
	p.Use(router)
	sv := service.NewReportService()
	gw := middle.NewHttpGateway(*sv)

	router.POST(post, gw.Save)
	router.GET(get, gw.Find)

	srv := &http.Server{
		Addr:    address,
		Handler: router,
	}

	go func() {
		// Service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			panic(err)
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
