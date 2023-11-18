package handler

import (
	"flat-docs-service/flat/docs/sample"
	"flat-docs-service/internal/builder"
	"flat-docs-service/pkg/service"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
)

const (
	defaultBufferOffset = 0
	capacity            = 100
)

type Handler struct {
	service service.ReportServiceInterface
	bb      *builder.Pool
}

func NewHandler(service service.ReportServiceInterface) *Handler {
	return &Handler{service: service, bb: builder.NewBuilderPool(capacity)}
}

func (h *Handler) Save(c *gin.Context) {
	request := c.Request
	requestBytes := getBytes(request)
	flatRequest := sample.GetRootAsSaveRequest(requestBytes, defaultBufferOffset)
	doc := new(sample.Document)
	flatDoc := flatRequest.Document(doc)

	message := "ok"
	err := h.service.Save(flatDoc)
	if err != nil {
		log.Println("Error while saving doc:", err)
		message = err.Error()
	}

	b := h.bb.Get()
	defer h.bb.Put(b)

	responseString := b.CreateString(message)
	sample.SaveResponseStart(b)
	sample.SaveResponseAddMessage(b, responseString)
	response := sample.SaveResponseEnd(b)
	b.Finish(response)
	fb := b.FinishedBytes()
	c.Header("Content-Type", "application/octet-stream")
	c.Data(200, "application/octet-stream", fb)
}

func (h *Handler) FindByParams(c *gin.Context) {
	request := c.Request
	buf := getBytes(request)
	flatRequest := sample.GetRootAsFindRequest(buf, defaultBufferOffset)
	response, err := h.service.Find(int(flatRequest.Limit()), int(flatRequest.Offset()))

	c.Header("Content-Type", "application/octet-stream")
	if err != nil {
		log.Println("Error while finding docs:", err)
		c.AbortWithError(500, err)
		return
	}
	c.Data(200, "application/octet-stream", *response)
}

func getBytes(request *http.Request) []byte {
	bytes, err := io.ReadAll(request.Body)
	if err != nil || len(bytes) == 0 {
		log.Fatal("Error while processioning request", err)
	}
	return bytes
}
