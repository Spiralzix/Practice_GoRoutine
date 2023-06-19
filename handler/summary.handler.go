package handler

import (
	"log"
	"net/http"

	"github.com/Spiralzix/LinemanAssignment/internal/core/service"
	"github.com/gin-gonic/gin"
)

type CovidHandler struct {
	service service.ICovidService
}

func NewCOVIDHandler(service service.ICovidService) *CovidHandler {
	return &CovidHandler{
		service: service,
	}
}

func (h *CovidHandler) GetSummary(c *gin.Context) {
	summaryReport, err := h.service.GetReport()
	if err != nil {
		log.Println("Failed to get a summary report", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get a summary report"})
		return
	}
	c.JSON(http.StatusOK, summaryReport)
}
