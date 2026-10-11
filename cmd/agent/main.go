package main

import (
	"log"
	"time"

	"github.com/ColleBoll/LabControl/internal/agent/client"
	"github.com/ColleBoll/LabControl/internal/agent/collectors"
	"github.com/ColleBoll/LabControl/pkg/protocol"
)

func main() {
	api := client.NewAPIClient("http://localhost:8080")

	system, err := collectors.SystemMetrics()
	if err != nil {
		log.Fatal(err)
	}

	report := protocol.AgentReport{
		AgentID:   system.Hostname,
		Timestamp: time.Now(),
		System:    system,
	}

	err = client.Send(api, "/api/v1/agents/report", report)
	if err != nil {
		log.Println(err)
	}

	log.Println("Report sent")
}
