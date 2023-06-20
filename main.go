package main

import (
	"fmt"
	"log"

	"github.com/Spiralzix/LinemanAssignment/config"
	"github.com/Spiralzix/LinemanAssignment/external"
	"github.com/Spiralzix/LinemanAssignment/handler"
	"github.com/Spiralzix/LinemanAssignment/internal/core/service"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg, err := config.NewConfig()
	fmt.Println("cfg", cfg)
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}
	record := external.NewCOVIDRecord(cfg)
	service := service.NewCOVIDService(record)
	handler := handler.NewCOVIDHandler(service)
	r := gin.Default()
	r.GET("/covid/summary", handler.GetSummary)
	r.Run(cfg.AppPort.Port)
}
