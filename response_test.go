package respond_test

import (
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	resp "github.com/nicklaw5/go-respond"
)

func newRequest(t *testing.T, method string) *http.Request {
	req, err := http.NewRequest(method, "/", nil)
	if err != nil {
		t.Fatal(err)
	}
	return req
}

func validateStatusCode(responseStatusCode int, expectedStatusCode int) error {
	if status := responseStatusCode; status != expectedStatusCode {
		return errors.New(fmt.Sprintf("Handler returned wrong status code: got %v wanted %v",
			status, expectedStatusCode))
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

func TestSetHeaders(t *testing.T) {
	req := newRequest(t, "GET")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		res := resp.NewResponse(w)
		headers := map[string]string{
			"X-One-1": "headers-one",
			"X-One-2": "headers-one",
		}
		res.SetHeaders(headers)

		headers = map[string]string{
			"X-Two-1": "headers-two",
			"X-Two-2": "headers-two",
		}
		res.SetHeaders(headers)
		res.Created(nil)
	})
	handler.ServeHTTP(rr, req)

	expected := "headers-two"
	if err := validateResponseBody(rr.Header().Get("X-Two-1"), expected); err != nil {
		t.Fatal(err.Error())
	}

	if err := validateResponseBody(rr.Header().Get("X-Two-2"), expected); err != nil {
		t.Fatal(err.Error())
	}
}

func TestAddHeader(t *testing.T) {
	req := newRequest(t, "GET")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		res := resp.NewResponse(w)
		res.AddHeader("Ping", "Pong")
		res.Created(nil)
	})
	handler.ServeHTTP(rr, req)

	expected := "Pong"
	if err := validateResponseBody(rr.Header().Get("Ping"), expected); err != nil {
		t.Fatal(err.Error())
	}
}

func TestDeleteHeader(t *testing.T) {
	req := newRequest(t, "GET")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		res := resp.NewResponse(w)
		res.Created(nil)
		res.DeleteHeader("Content-Type")
	})
	handler.ServeHTTP(rr, req)

	expected := ""
	if err := validateResponseBody(rr.Header().Get("Content-Type"), expected); err != nil {
		t.Fatal(err.Error())
	}
}
