package handle

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"json-docs-service/pkg/model"
	s "json-docs-service/pkg/service"
	"net/http"
	"strconv"
)

type GatewayInterface interface {
	Find(c *gin.Context)
	Save(c *gin.Context)
}

type HttpGateway struct {
	GatewayInterface
	reportService s.ReportServiceInterface
}

func NewHttpGateway(rs *s.ReportService) *HttpGateway {
	return &HttpGateway{reportService: rs}
}

func (h HttpGateway) Find(c *gin.Context) {
	limitParam := c.DefaultQuery("limit", "10")
	offsetParam := c.DefaultQuery("offset", "0")
	limit, _ := strconv.Atoi(limitParam)
	offset, _ := strconv.Atoi(offsetParam)

	result, err := h.reportService.Find(limit, offset)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		return
	}

	if len(result) != 0 {
		c.IndentedJSON(http.StatusOK, result)
	} else {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "reports not found"})
	}
}

func (h HttpGateway) Save(c *gin.Context) {
	var document model.Document
	if err := c.BindJSON(&document); err != nil {
		fmt.Println("Error while parsing the doc", err)
		sendError(c)
		return
	}

	err := h.reportService.Save(document)
	if err != nil {
		fmt.Println("Error while saving the doc", err)
		sendError(c)
		return
	}

	c.IndentedJSON(http.StatusCreated, gin.H{"status": "ok"})
}

func (h HttpGateway) Validate(c *gin.Context) {
	var document model.Document
	if err := c.BindJSON(&document); err != nil {
		fmt.Println("Error while parsing the doc", err)
		sendError(c)
		return
	}

	err := h.reportService.Validate(document)
	if err != nil {
		fmt.Println("Error while validating the doc", err)
		sendError(c)
		return
	}

	c.IndentedJSON(http.StatusCreated, gin.H{"status": "ok"})
}

func sendError(c *gin.Context) {
	c.IndentedJSON(http.StatusInternalServerError, gin.H{"status": "error"})
}
