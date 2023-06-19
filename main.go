package main

import (
	"github.com/Spiralzix/LinemanAssignment/handler"
	"github.com/Spiralzix/LinemanAssignment/internal/core/repositorys"
	"github.com/Spiralzix/LinemanAssignment/internal/core/service"
	"github.com/gin-gonic/gin"
)

func main() {
	repo := repositorys.NewCOVIDRepository()
	service := service.NewCOVIDService(repo)
	handler := handler.NewCOVIDHandler(service)
	r := gin.Default()
	r.GET("/covid/summary", handler.GetSummary)
	r.Run()
}
