package main

import (
	"fmt"
	"time"

	"github.com/Spiralzix/LinemanAssignment/handler"
	"github.com/Spiralzix/LinemanAssignment/internal/core/repositorys"
	"github.com/Spiralzix/LinemanAssignment/internal/core/service"
	"github.com/gin-gonic/gin"
)

func main() {
	start := time.Now()
	repo := repositorys.NewCOVIDRepository()
	service := service.NewCOVIDService(repo)
	handler := handler.NewCOVIDHandler(service)

	r := gin.Default()
	r.GET("/covid/summary", handler.GetSummary)
	fmt.Println("ใช้เวลาในการ Run ทั้งสิ้น : ", time.Since(start), " วินาที")
	r.Run()

}
