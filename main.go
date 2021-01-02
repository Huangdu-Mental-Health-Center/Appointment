package main

import (
	"Huangdu_HMC_Appointment/src/handler"
	"Huangdu_HMC_Appointment/src/logger"
	"io"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func setupLogs() {

	gin.DisableConsoleColor()
	f, err := os.Create("gin.log")
	if err != nil {
		logger.Error.Printf("create gin.log failed: %v;\n", err)
	}
	gin.DefaultWriter = io.MultiWriter(f)
}

func main() {
	logger.Init()
	logger.Info.Println("Huangdu HMC Appointment module.")
	setupLogs()
	gin.SetMode(gin.DebugMode)
	router := gin.Default()
	h := handler.Handler{}
	router.Use(cors.Default())
	router.GET("/api/appointment/get", h.GetAppoint)
	router.POST("/api/appointment/new", h.NewAppoint)
	router.DELETE("/api/appointment/del", h.DelAppoint)
	_ = router.Run(":5700")
}
