package agents

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ColleBoll/LabControl/internal/server/response"
	"github.com/ColleBoll/LabControl/pkg/protocol"
)

func ReportHandler(w http.ResponseWriter, r *http.Request) {
	var report protocol.AgentReport

	if err := json.NewDecoder(r.Body).Decode(&report); err != nil {
		response.ErrorResponseJSON(w, http.StatusBadRequest, "invalid_request", "invalid request")
		return
	}

	response.ReportAccept(w, "Thx agent for the great report!")

	fmt.Println(report)
}
