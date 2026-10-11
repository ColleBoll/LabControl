package agents

import (
	"fmt"
	"net/http"

	"github.com/ColleBoll/LabControl/pkg/protocol"
	"github.com/gin-gonic/gin"
)

func ReportHandler(c *gin.Context) {
	var report protocol.AgentReport

	if err := c.ShouldBindJSON(&report); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request body",
		})
	}

	fmt.Printf("Agent report recieved: %+v\n", report)

	c.Status(http.StatusAccepted)
}
