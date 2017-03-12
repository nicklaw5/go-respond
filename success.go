package respond

import "net/http"

// SuccessResponse is a standard HTTP JSON response
type SuccessResponse struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
}

// Created returns a 201 Created response
func (resp *HTTPResponse) Created(v interface{}) {
	resp.SetStatusCode(http.StatusCreated).
		SetBody(resp.MarshallJSON(&SuccessResponse{true, v})).
		WriteResponse()
}
