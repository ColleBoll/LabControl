package response

import "net/http"

type Accept struct {
	Message string `json:"message"`
}

func ReportAccept(w http.ResponseWriter, message string) {
	JSON(w, http.StatusAccepted, Accept{
		Message: message,
	})
}
