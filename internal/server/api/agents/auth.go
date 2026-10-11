package agents

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/ColleBoll/LabControl/internal/server/database"
	"github.com/ColleBoll/LabControl/pkg/protocol"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func HandleAgentAuth(c *gin.Context) {
	var handshake protocol.AgentHandshake

	if err := c.ShouldBindJSON(&handshake); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid body request",
		})
	}

	var agent database.Agent

	// check of agent al bestaat
	result := database.GetServerDB().First(&agent, 1)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {

			// agent does not exist
			database.GetServerDB().Create(database.Agent{
				ID:        uint(handshake.ID),
				Hostname:  handshake.Hostname,
				UpdatedAt: time.Now(),
				CreatedAt: time.Now(),
			})

			c.Status(http.StatusCreated)
			fmt.Println("Registered and created new agent: ", handshake.ID)
		} else {
			// andere erro
		}
	}

	// zo ja >>
}
