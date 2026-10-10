package api

import (
	"github.com/ColleBoll/LabControl/internal/server/api/agents"
	"github.com/gin-gonic/gin"
)

func NewRouter() {
	router := gin.Default()

	router.GET("/api/v1/health", HealthHandler)
	router.POST("/api/v1/agents/report", agents.ReportHandler)

	router.Run(":8080")
}
