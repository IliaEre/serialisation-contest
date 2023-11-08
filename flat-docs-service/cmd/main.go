package main

import (
	"flat--docs-service/internal/service"
	"flat--docs-service/pkg/handler"
	"net/http"
)

var addr = "0.0.0.0:50051"

func main() {
	s := service.NewReportService()
	h := handler.NewHandler(s)
	http.HandleFunc("/save", h.Save)
	http.HandleFunc("/find", h.FindByParams)
	http.ListenAndServe(addr, nil)
}
