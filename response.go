package respond

import (
	"encoding/json"
	"net/http"
)

// HTTPResponse is our HTTP response
type HTTPResponse struct {
	Writer     http.ResponseWriter
	StatusCode int
	Body       []byte
	Headers    map[string]string
	JSON       bool
}

// NewResponse sets the HTTP response status code
func NewResponse(w http.ResponseWriter) *HTTPResponse {
	return &HTTPResponse{Writer: w}
}

// IsJSON set the response type to json
func (resp *HTTPResponse) IsJSON(isJSON bool) *HTTPResponse {
	if isJSON {
		resp.JSON = true
		resp.SetJSONHeader()
	}
	return resp
}

// SetJSONHeader set the response type to json and attached the json header
func (resp *HTTPResponse) SetJSONHeader() *HTTPResponse {
	resp.AddHeader("Content-Type", "application/json; charset=utf-8")
	return resp
}

// SetStatusCode sets the HTTP response status code
func (resp *HTTPResponse) SetStatusCode(code int) *HTTPResponse {
	resp.StatusCode = code
	return resp
}

// SetBody sets the HTTP response body
func (resp *HTTPResponse) SetBody(body []byte) *HTTPResponse {
	resp.Body = body
	return resp
}

// SetHeaders sets the HTTP response headers. Any existing set headers that
// exists before this method is called will removed. Use AddHeader to append
// any existing headers
func (resp *HTTPResponse) SetHeaders(headers map[string]string) *HTTPResponse {

	// Remove any existing headers
	if len(resp.Headers) > 0 {
		for key := range resp.Headers {
			resp.DeleteHeader(key)
		}
	}

	resp.Headers = headers
	for key, value := range headers {
		resp.Writer.Header().Set(key, value)
	}
	return resp
}

// DeleteHeader deletes a single provided header
func (resp *HTTPResponse) DeleteHeader(key string) *HTTPResponse {
	resp.Writer.Header().Del(key)
	return resp
}

// AddHeader adds a single header to the repsonse
func (resp *HTTPResponse) AddHeader(key string, value string) *HTTPResponse {
	resp.Writer.Header().Add(key, value)
	return resp
}

// WriteResponse write the HTTP response headers and body
func (resp *HTTPResponse) WriteResponse() {
	// TODO: check status code, headers and body are set
	// TODO: set headers separately
	// resp.Writer.Header().Set("Content-Type", "application/json")
	resp.Writer.WriteHeader(resp.StatusCode)
	resp.Writer.Write(resp.Body)
}

// MarshallJSON transforms an interface into json
func (resp *HTTPResponse) MarshallJSON(v interface{}) []byte {
	data, _ := json.Marshal(v)
	return data
}
