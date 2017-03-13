package respond

import "net/http"

// ErrorResponse is a standard HTTP JSON error
type ErrorResponse struct {
	Success bool   `json:"success"`
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// BadRequest returns a 400 Bad Request JSON response
func (resp *HTTPResponse) BadRequest(message string) {
	resp.WriteErrorResponse(http.StatusBadRequest, message)
}

// Unauthorized returns a 401 Unauthorized JSON response
func (resp *HTTPResponse) Unauthorized(message string) {
	resp.WriteErrorResponse(http.StatusUnauthorized, message)
}

// Forbidden returns a 401 Forbidden JSON response
func (resp *HTTPResponse) Forbidden(message string) {
	resp.WriteErrorResponse(http.StatusForbidden, message)
}

// NotFound returns a 404 Not Found JSON response
func (resp *HTTPResponse) NotFound(message string) {
	resp.WriteErrorResponse(http.StatusNotFound, message)
}

// MethodNotAllowed returns a 405 Method Not Allowed JSON response
func (resp *HTTPResponse) MethodNotAllowed(message string) {
	resp.WriteErrorResponse(http.StatusMethodNotAllowed, message)
}

// UnprocessableEntity returns a 422 Unprocessable Entity JSON response
func (resp *HTTPResponse) UnprocessableEntity(message string) {
	resp.WriteErrorResponse(http.StatusUnprocessableEntity, message)
}

// Conflict returns a 409 Conflict JSON response
func (resp *HTTPResponse) Conflict(message string) {
	resp.WriteErrorResponse(http.StatusConflict, message)
}

// InternalServerError returns a 500 Internal Server Error JSON response
func (resp *HTTPResponse) InternalServerError(message string) {
	resp.WriteErrorResponse(http.StatusInternalServerError, message)
}

// WriteErrorResponse is the error response writer
func (resp *HTTPResponse) WriteErrorResponse(code int, message string) {
	resp.SetStatusCode(code)
	body := &ErrorResponse{false, resp.StatusCode, message}
	resp.SetBody(resp.MarshallJSON(body)).
		WriteResponse()
}
