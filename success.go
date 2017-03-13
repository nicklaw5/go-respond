package respond

import "net/http"

// SuccessResponse is a standard HTTP JSON response
type SuccessResponse struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
}

// Ok returns a 200 OK JSON response
func (resp *HTTPResponse) Ok(v interface{}) {
	resp.WriteSuccessResponse(http.StatusOK, v)
}

// Created returns a 201 Created JSON response
func (resp *HTTPResponse) Created(v interface{}) {
	resp.WriteSuccessResponse(http.StatusCreated, v)
}

// Accepted returns a 202 Accepted JSON response
func (resp *HTTPResponse) Accepted(v interface{}) {
	resp.WriteSuccessResponse(http.StatusAccepted, v)
}

// NoContent returns a 204 No Content JSON response
func (resp *HTTPResponse) NoContent() {
	resp.WriteSuccessResponse(http.StatusNoContent, nil)
}

// WriteSuccessResponse is the error response writer
func (resp *HTTPResponse) WriteSuccessResponse(code int, v interface{}) {
	resp.SetStatusCode(code)
	body := &SuccessResponse{true, v}
	resp.SetBody(resp.MarshallJSON(body)).
		WriteResponse()
}
