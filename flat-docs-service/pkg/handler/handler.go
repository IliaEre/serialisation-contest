package handler

import (
	"flat--docs-service/flat/docs/sample"
	"flat--docs-service/internal/service"
	"io"
	"log"
	"net/http"
)

type Handler struct {
	s service.ReportServiceInterface
}

func NewHandler(service service.ReportServiceInterface) *Handler {
	return &Handler{s: service}
}

func (h Handler) Save(w http.ResponseWriter, request *http.Request) {
	buf, err := io.ReadAll(request.Body)
	if err != nil {
		log.Fatal("request", err)
	}
	docs := sample.GetRootAsSaveRequest(buf, 0)
	h.s.Save(*docs)

	w.Header().Set("Content-Type", "application/octet-stream")
	w.WriteHeader(200)
}

func (h Handler) FindByParams(w http.ResponseWriter, request *http.Request) {
	print(request)
}
