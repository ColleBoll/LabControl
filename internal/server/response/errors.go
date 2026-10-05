package response

import "net/http"

type ErrorResponse struct {
	Error Error `json:"error"`
}

type Error struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func ErrorResponseJSON(w http.ResponseWriter, status int, code, message string) {
	JSON(w, status, ErrorResponse{
		Error: Error{
			Code:    code,
			Message: message,
		},
	})
}
