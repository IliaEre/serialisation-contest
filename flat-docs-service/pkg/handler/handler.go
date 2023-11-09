package handler

import (
	"flat--docs-service/flat/docs/sample"
	"flat--docs-service/internal/service"
	"flat--docs-service/pkg/mapper"
	flatbuffers "github.com/google/flatbuffers/go"
	"io"
	"log"
	"net/http"
)

const buffSize = 256

type Handler struct {
	s service.ReportServiceInterface
}

func NewHandler(service service.ReportServiceInterface) *Handler {
	return &Handler{s: service}
}

func (h Handler) Save(w http.ResponseWriter, request *http.Request) {
	buf := getBytes(request)
	flatRequest := sample.GetRootAsSaveRequest(buf, 0)
	doc := new(sample.Document)
	flatDoc := flatRequest.Document(doc)

	err := h.s.Save(*flatDoc)

	w.Header().Set("Content-Type", "application/octet-stream")
	builder := flatbuffers.NewBuilder(buffSize)

	var message flatbuffers.UOffsetT
	if err != nil {
		message = builder.CreateString(err.Error())
	} else {
		message = builder.CreateString("ok")
	}

	sample.SaveResponseStart(builder)
	sample.SaveResponseAddMessage(builder, message)
	response := sample.SaveResponseEnd(builder)
	builder.Finish(response)
	fb := builder.FinishedBytes()
	_, err = w.Write(fb)
	if err != nil {
		log.Fatal("parse", err)
	}
}

func (h Handler) FindByParams(w http.ResponseWriter, request *http.Request) {
	buf := getBytes(request)
	flatRequest := sample.GetRootAsFindRequest(buf, 0)
	docs := h.s.Find(int(flatRequest.Limit()), int(flatRequest.Offset()))

	builder := flatbuffers.NewBuilder(1024)
	off := make([]flatbuffers.UOffsetT, len(docs))
	for i := 0; i < len(docs); i++ {
		fb := sample.GetRootAsDocument(docs[i].Table().Bytes, 0)
		doc := mapper.CreateDocument(builder, fb)
		off = append(off, doc)
	}

	sample.FindResponseStartDocsVector(builder, len(off))
	for i := len(off) - 1; i >= 0; i-- {
		builder.PrependUOffsetT(off[i])
	}
	docsVector := builder.EndVector(len(off))

	w.Header().Set("Content-Type", "application/octet-stream")

	sample.FindResponseStart(builder)
	sample.FindResponseAddDocs(builder, docsVector)
	response := sample.FindResponseEnd(builder)
	builder.Finish(response)
	fb := builder.FinishedBytes()
	_, err := w.Write(fb)
	if err != nil {
		log.Fatal("parse", err)
	}
}

func getBytes(request *http.Request) []byte {
	buf, err := io.ReadAll(request.Body)
	if err != nil || len(buf) == 0 {
		log.Fatal("request", err)
	}
	return buf
}
