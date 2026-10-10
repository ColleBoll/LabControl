package main

import (
	"log"
	"time"

	"github.com/ColleBoll/LabControl/internal/agent/client"
	"github.com/ColleBoll/LabControl/internal/agent/collectors"
	"github.com/ColleBoll/LabControl/pkg/protocol"
)

func main() {
	server := client.New("http://localhost:8080")

	system, err := collectors.SystemMetrics()
	if err != nil {
		log.Fatal(err)
	}

	report := protocol.AgentReport{
		AgentID:   system.Hostname,
		Timestamp: time.Now(),
		System:    system,
	}

	if err := server.SendReport(report); err != nil {
		log.Fatal(err)
	}

	log.Println("Report sent")
}
