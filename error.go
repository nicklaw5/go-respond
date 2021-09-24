package respond

import "net/http"

// BadRequest returns a 400 Bad Request JSON response
func (resp *Response) BadRequest(v interface{}) {
	resp.writeResponse(http.StatusBadRequest, v)
}

// Unauthorized returns a 401 Unauthorized JSON response
func (resp *Response) Unauthorized(v interface{}) {
	resp.writeResponse(http.StatusUnauthorized, v)
}

// Forbidden returns a 403 Forbidden JSON response
func (resp *Response) Forbidden(v interface{}) {
	resp.writeResponse(http.StatusForbidden, v)
}

// NotFound returns a 404 Not Found JSON response
func (resp *Response) NotFound(v interface{}) {
	resp.writeResponse(http.StatusNotFound, v)
}

// MethodNotAllowed returns a 405 Method Not Allowed JSON response
func (resp *Response) MethodNotAllowed(v interface{}) {
	resp.writeResponse(http.StatusMethodNotAllowed, v)
}

// NotAcceptable returns a 406 Not Acceptable JSON response
func (resp *Response) NotAcceptable(v interface{}) {
	resp.writeResponse(http.StatusNotAcceptable, v)
}

// Conflict returns a 409 Conflict JSON response
func (resp *Response) Conflict(v interface{}) {
	resp.writeResponse(http.StatusConflict, v)
}

// Gone returns a 410 Gone JSON response
func (resp *Response) Gone(v interface{}) {
	resp.writeResponse(http.StatusGone, v)
}

// LengthRequired returns a 411 Length Required JSON response
func (resp *Response) LengthRequired(v interface{}) {
	resp.writeResponse(http.StatusLengthRequired, v)
}

// PreconditionFailed returns a 412 Precondition Failed JSON response
func (resp *Response) PreconditionFailed(v interface{}) {
	resp.writeResponse(http.StatusPreconditionFailed, v)
}

// RequestEntityTooLarge returns a 413 Request Entity Too Large JSON response
func (resp *Response) RequestEntityTooLarge(v interface{}) {
	resp.writeResponse(http.StatusRequestEntityTooLarge, v)
}

// UnsupportedMediaType returns a 415 Unsupported Media Type JSON response
func (resp *Response) UnsupportedMediaType(v interface{}) {
	resp.writeResponse(http.StatusUnsupportedMediaType, v)
}

// UnprocessableEntity returns a 422 Unprocessable Entity JSON response
func (resp *Response) UnprocessableEntity(v interface{}) {
	resp.writeResponse(http.StatusUnprocessableEntity, v)
}

// InternalServerError returns a 500 Internal Server Error JSON response
func (resp *Response) InternalServerError(v interface{}) {
	resp.writeResponse(http.StatusInternalServerError, v)
}

// NotImplemented returns a 501 Not Implemented JSON response
func (resp *Response) NotImplemented(v interface{}) {
	resp.writeResponse(http.StatusNotImplemented, v)
}

// BadGateway returns a 502 Bad Gateway JSON response
func (resp *Response) BadGateway(v interface{}) {
	resp.writeResponse(http.StatusBadGateway, v)
}

// ServiceUnavailable returns a 503 Service Unavailable JSON response
func (resp *Response) ServiceUnavailable(v interface{}) {
	resp.writeResponse(http.StatusServiceUnavailable, v)
}

// GatewayTimeout returns a 504 Gateway Timeout JSON response
func (resp *Response) GatewayTimeout(v interface{}) {
	resp.writeResponse(http.StatusGatewayTimeout, v)
}
