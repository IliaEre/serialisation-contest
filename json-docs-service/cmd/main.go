package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"json-docs-service/internal/middle"
	"json-docs-service/internal/service"
	"net/http"
	"os"
	"os/signal"
	"time"
)

const (
	post = "/report"
	get  = "/reports"
	TTL  = 5
)

func main() {
	router := gin.Default()
	sv := service.NewReportService()
	gw := middle.NewHttpGateway(*sv)

	router.POST(post, gw.Save)
	router.GET(get, gw.Find)

	srv := &http.Server{
		Addr:    "localhost:8080",
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
