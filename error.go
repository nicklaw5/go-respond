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
	resp.WriteErrorResponse(http.StatusUnprocessableEntity, message)
}

// Conflict returns a 409 Conflict json response
func (resp *HTTPResponse) Conflict(message string) {
	resp.WriteErrorResponse(http.StatusConflict, message)
}

// WriteErrorResponse is the error response writer
func (resp *HTTPResponse) WriteErrorResponse(code int, message string) {
	resp.SetStatusCode(code)
	body := &ErrorResponse{false, resp.StatusCode, message}
	resp.SetBody(resp.MarshallJSON(body)).
		WriteResponse()
}
