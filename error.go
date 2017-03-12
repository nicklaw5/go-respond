package respond

import "net/http"

// ErrorResponse is a standard HTTP JSON error
type ErrorResponse struct {
	Success bool   `json:"success"`
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// UnprocessableEntity returns a 422 Unprocessable Entity json response
func (resp *HTTPResponse) UnprocessableEntity(message string) {
	resp.SetStatusCode(http.StatusUnprocessableEntity).
		SetBody(resp.MarshallJSON(&ErrorResponse{false, resp.StatusCode, message})).
		WriteResponse()
}

// Conflict returns a 409 Conflict json response
func (resp *HTTPResponse) Conflict(message string) {
	resp.SetStatusCode(http.StatusConflict).
		SetBody(resp.MarshallJSON(&ErrorResponse{false, resp.StatusCode, message})).
		WriteResponse()
}
