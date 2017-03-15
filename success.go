package respond

import (
	"net/http"
)

// Ok returns a 200 OK JSON response
func (resp *Response) Ok(v interface{}) {
	resp.writeResponse(http.StatusOK, v)
}

// Created returns a 201 Created JSON response
func (resp *Response) Created(v interface{}) {
	resp.writeResponse(http.StatusCreated, v)
}

// Accepted returns a 202 Accepted JSON response
func (resp *Response) Accepted(v interface{}) {
	resp.writeResponse(http.StatusAccepted, v)
}

// NoContent returns a 204 No Content JSON response
func (resp *Response) NoContent() {
	resp.writeResponse(http.StatusNoContent, nil)
}
