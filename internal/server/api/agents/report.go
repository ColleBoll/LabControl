package agents

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ColleBoll/LabControl/pkg/protocol"
	"github.com/gin-gonic/gin"
)

func ReportHandler(c *gin.Context) {
	var report protocol.AgentReport

	if err := json.NewDecoder(c.Request.Body).Decode(&report); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"invalid_request": "invalid request"})
		return
	}

	c.JSON(http.StatusAccepted, "Thx agent for the great report!")

	fmt.Println(report)
}
