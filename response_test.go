package respond_test

import (
	"errors"
	"fmt"
	resp "github.com/nicklaw5/go-respond"
	"net/http"
	"net/http/httptest"
	"testing"
)

func newRequest(t *testing.T, method string) *http.Request {
	req, err := http.NewRequest(method, "/", nil)
	if err != nil {
		t.Fatal(err)
	}
	return req
}

func validateStatusCode(responseStatusCode int, expectedStatusCode int) error {
	if responseStatusCode != expectedStatusCode {
		return errors.New(fmt.Sprintf("Handler returned wrong status code: got %v wanted %v",
			responseStatusCode, expectedStatusCode))
	}
	return nil
}

func validateResponseBody(responseBody string, expectedBody string) error {
	if responseBody != expectedBody {
		return errors.New(fmt.Sprintf("Handler returned unexpected body: got %v wanted %v",
			responseBody, expectedBody))
	}
	return nil
}

func validateResponseHeader(responseHeaderValue string, expectedHeaderValue string) error {
	if responseHeaderValue != expectedHeaderValue {
		return errors.New(fmt.Sprintf("Handler returned unexpected body: got %v wanted %v",
			responseHeaderValue, expectedHeaderValue))
	}
	return nil
}

func TestDefaultMessage(t *testing.T) {
	t.Parallel()

	req := newRequest(t, "GET")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp.NewResponse(w).DefaultMessage().
			Unauthorized(nil)
	})
	handler.ServeHTTP(rr, req)

	if err := validateResponseBody(rr.Body.String(), "{\"status\":401,\"message\":\"Unauthorized\"}"); err != nil {
		t.Fatal(err)
	}
}

func TestRespondInvalidType(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("responding with invalid type (channel) should have caused a panic")
		}
	}()

	t.Parallel()

	req := newRequest(t, "GET")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp.NewResponse(w).DefaultMessage().
			Unauthorized(make(chan int))
	})
	handler.ServeHTTP(rr, req)
}

func TestContentTypeHeader(t *testing.T) {
	t.Parallel()

	req := newRequest(t, "POST")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp.NewResponse(w).
			Ok(nil)
	})
	handler.ServeHTTP(rr, req)

	contentType := "application/json; charset=utf-8"
	if err := validateResponseHeader(rr.Header().Get("Content-Type"), contentType); err != nil {
		t.Fatal(err.Error())
	}
}

func TestAddHeader(t *testing.T) {
	t.Parallel()

	req := newRequest(t, "POST")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp.NewResponse(w).
			AddHeader("foo", "bar").
			AddHeader("ping", "pong").
			Ok(nil)
	})
	handler.ServeHTTP(rr, req)

	if err := validateResponseHeader(rr.Header().Get("ping"), "pong"); err != nil {
		t.Fatal(err.Error())
	}

	if err := validateResponseHeader(rr.Header().Get("foo"), "bar"); err != nil {
		t.Fatal(err.Error())
	}
}

func TestDeleteHeader(t *testing.T) {
	t.Parallel()

	req := newRequest(t, "GET")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		res := resp.NewResponse(w)

		res.AddHeader("foo", "bar").
			Ok(nil)

		res.DeleteHeader("foo")
	})
	handler.ServeHTTP(rr, req)

	if err := validateResponseBody(rr.Header().Get("foo"), ""); err != nil {
		t.Fatal(err.Error())
	}
}
