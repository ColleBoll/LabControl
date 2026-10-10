package api

import (
	"github.com/ColleBoll/LabControl/internal/server/api/agents"
	"github.com/gin-gonic/gin"
)

func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Next()
	}
}

func NewRouter() {
	router := gin.New()
	router.Use(gin.Logger(), gin.Recovery())

	v1 := router.Group("/api/v1")
	//v1.Use(AuthRequired())
	{
		v1.GET("/health", HealthHandler)
		v1.POST("/agents/report", agents.ReportHandler)
	}

	router.Run(":8080")
}
