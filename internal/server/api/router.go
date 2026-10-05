package api

import (
	"net/http"

	"github.com/ColleBoll/LabControl/internal/server/api/agents"
)

func NewRouter() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /api/v1/health", HealthHandler)

	mux.HandleFunc("POST /api/v1/agents/report", agents.ReportHandler)

	return mux
}
