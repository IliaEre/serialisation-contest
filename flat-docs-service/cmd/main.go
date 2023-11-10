package main

import (
	"context"
	"flat--docs-service/internal/service"
	"flat--docs-service/pkg/handler"
	"fmt"
	"github.com/gin-gonic/gin"
	ginprometheus "github.com/zsais/go-gin-prometheus"
	"net/http"
	"os"
	"os/signal"
	"time"
)

const (
	post = "/save"
	get  = "/find"
	TTL  = 5
	addr = "0.0.0.0:50051"
)

func main() {
	router := gin.Default()
	p := ginprometheus.NewPrometheus("gin")
	p.Use(router)
	sv := service.NewReportService()
	gw := handler.NewHandler(*sv)

	router.POST(post, gw.Save)
	router.GET(get, gw.FindByParams)

	srv := &http.Server{
		Addr:    addr,
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
