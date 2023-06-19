package handler

import (
	"log"
	"net/http"

	"github.com/Spiralzix/LinemanAssignment/internal/core/service"
	"github.com/gin-gonic/gin"
)

type covidHandler struct {
	service service.ICovidService
}

func NewCOVIDHandler(service service.ICovidService) *covidHandler {
	return &covidHandler{
		service: service,
	}
}

func (h *covidHandler) GetSummary(c *gin.Context) {
	summaryReport, err := h.service.GetReport()
	if err != nil {
		log.Println("Failed to fetch COVID data:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch COVID data"})
		return
	}
	c.JSON(http.StatusOK, summaryReport)
}
