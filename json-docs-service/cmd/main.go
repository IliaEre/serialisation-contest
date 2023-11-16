package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	ginprometheus "github.com/zsais/go-gin-prometheus"
	"go.mongodb.org/mongo-driver/mongo"
	"json-docs-service/pkg/db"
	"json-docs-service/pkg/middle"
	"json-docs-service/pkg/service"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

const (
	post         = "/report"
	get          = "/reports"
	TTL          = 5
	address      = ":9091"
	mongoAddress = "mongodb://localhost:27017"
)

func main() {
	router := gin.Default()
	p := ginprometheus.NewPrometheus("gin")
	p.Use(router)

	rc, err := db.NewMongoRepository(mongoAddress)
	if err != nil {
		log.Fatal("Could create mongo:", err)
	}

	defer func(Client *mongo.Client, ctx context.Context) {
		err := Client.Disconnect(ctx)
		if err != nil {
			log.Fatal("Could create mongo:", err)

		}
	}(&rc.Client, context.Background())

	sv := service.NewReportService(*rc)
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
