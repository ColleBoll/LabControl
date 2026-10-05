package api

import (
	"net/http"

	"github.com/ColleBoll/LabControl/internal/server/response"
)

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	response.JSON(w, http.StatusOK, map[string]string{
		"status": "ok",
	})
}
