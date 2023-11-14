package middle

import (
	"github.com/gin-gonic/gin"
	"json-docs-service/internal/model"
	service2 "json-docs-service/pkg/service"
	"net/http"
	"strconv"
)

type GatewayInterface interface {
	Find(c *gin.Context)
	Save(c *gin.Context)
}

type HttpGateway struct {
	GatewayInterface
	reportService service2.ReportServiceInterface
}

func NewHttpGateway(rs service2.ReportService) *HttpGateway {
	return &HttpGateway{reportService: rs}
}

func (h HttpGateway) Find(c *gin.Context) {
	limitParam := c.DefaultQuery("limit", "10")
	offsetParam := c.DefaultQuery("offset", "0")
	limit, _ := strconv.Atoi(limitParam)
	offset, _ := strconv.Atoi(offsetParam)

	result := h.reportService.Find(limit, offset)
	if len(result) != 0 {
		c.IndentedJSON(http.StatusOK, result)
	} else {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "reports not found"})
	}
}

func (h HttpGateway) Save(c *gin.Context) {
	var document model.Document
	if err := c.BindJSON(&document); err != nil {
		return
	}

	h.reportService.Save(document)
	c.IndentedJSON(http.StatusCreated, gin.H{"status": "ok"})
}
