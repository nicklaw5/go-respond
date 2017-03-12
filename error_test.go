package respond_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	resp "github.com/nicklaw5/go-respond"
)

func TestUnprocessableEntity(t *testing.T) {
	req, err := http.NewRequest("POST", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		res := resp.NewResponse(w)
		res.UnprocessableEntity("An error occured")
	})
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusUnprocessableEntity {
		t.Errorf("Handler returned wrong status code: got %v wanted %v",
			status, http.StatusOK)
	}

	expected := `{"success":false,"code":422,"message":"An error occured"}`
	if rr.Body.String() != expected {
		t.Errorf("Handler returned unexpected body: got %v wanted %v",
			rr.Body.String(), expected)
	}
}

func TestConflict(t *testing.T) {
	req, err := http.NewRequest("POST", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		res := resp.NewResponse(w)
		res.Conflict("An error occured")
	})
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusConflict {
		t.Errorf("Handler returned wrong status code: got %v wanted %v",
			status, http.StatusOK)
	}

	expected := `{"success":false,"code":409,"message":"An error occured"}`
	if rr.Body.String() != expected {
		t.Errorf("Handler returned unexpected body: got %v wanted %v",
			rr.Body.String(), expected)
	}
}
