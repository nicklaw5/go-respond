package respond

import (
	"encoding/json"
	"net/http"
)

// Response is the HTTP response
type Response struct {
	Writer     http.ResponseWriter
	Headers    map[string]string
	DefMessage bool
}

// DefaultMessageResponse is for transporting a default http message
type DefaultMessageResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

// NewResponse creates and returns a new response
func NewResponse(w http.ResponseWriter) *Response {
	return &Response{
		Writer: w,
		Headers: map[string]string{
			"Content-Type": "application/json; charset=utf-8",
		},
	}
}

func (resp *Response) DefaultMessage() *Response {
	resp.DefMessage = true
	return resp
}

// DeleteHeader deletes a single header from the response
func (resp *Response) DeleteHeader(key string) *Response {
	resp.Writer.Header().Del(key)
	return resp
}

// AddHeader adds a single header to the response
func (resp *Response) AddHeader(key string, value string) *Response {
	resp.Writer.Header().Add(key, value)
	return resp
}

// WriteResponse writes the HTTP response status, headers and body
func (resp *Response) writeResponse(code int, v interface{}) error {
	if len(resp.Headers) > 0 {
		resp.writeHeaders()
	}

	resp.writeStatusCode(code)

	if v == nil && resp.DefMessage {
		v = DefaultMessageResponse{
			Status:  code,
			Message: http.StatusText(code),
		}
	}

	if v != nil {
		body, err := json.Marshal(v)
		if err != nil {
			panic(err)
		}
		// can just return an error when connection is hijacked or content-size is longer then declared.
		if _, err := resp.Writer.Write(body); err != nil {
			panic(err)
		}
	}

	return nil
}

func (resp *Response) writeHeaders() {
	for key, value := range resp.Headers {
		resp.Writer.Header().Set(key, value)
	}
}

func (resp *Response) writeStatusCode(code int) {
	resp.Writer.WriteHeader(code)
}
