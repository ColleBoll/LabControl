package main

import (
	"fmt"
	"net/http"

	"github.com/ColleBoll/LabControl/internal/server/api"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /api/health", api.HealthHandler)

	fmt.Println("LabControl running on http://localhost:8080")

	if err := http.ListenAndServe(":8080", mux); err != nil {
		panic(err)
	}
}
