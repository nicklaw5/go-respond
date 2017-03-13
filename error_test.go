package respond_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	resp "github.com/nicklaw5/go-respond"
)

func TestBadRequest(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		res := resp.NewResponse(w)
		res.BadRequest("Bad request")
	})
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("Handler returned wrong status code: got %v wanted %v",
			status, http.StatusBadRequest)
	}

	expected := `{"success":false,"code":400,"message":"Bad request"}`
	if rr.Body.String() != expected {
		t.Errorf("Handler returned unexpected body: got %v wanted %v",
			rr.Body.String(), expected)
	}
}

func TestUnauthorized(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		res := resp.NewResponse(w)
		res.Unauthorized("Unauthorized")
	})
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusUnauthorized {
		t.Errorf("Handler returned wrong status code: got %v wanted %v",
			status, http.StatusUnauthorized)
	}

	expected := `{"success":false,"code":401,"message":"Unauthorized"}`
	if rr.Body.String() != expected {
		t.Errorf("Handler returned unexpected body: got %v wanted %v",
			rr.Body.String(), expected)
	}
}

func TestForbidden(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		res := resp.NewResponse(w)
		res.Forbidden("Forbidden")
	})
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusForbidden {
		t.Errorf("Handler returned wrong status code: got %v wanted %v",
			status, http.StatusForbidden)
	}

	expected := `{"success":false,"code":403,"message":"Forbidden"}`
	if rr.Body.String() != expected {
		t.Errorf("Handler returned unexpected body: got %v wanted %v",
			rr.Body.String(), expected)
	}
}

func TestNotFound(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		res := resp.NewResponse(w)
		res.NotFound("Not found")
	})
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf("Handler returned wrong status code: got %v wanted %v",
			status, http.StatusNotFound)
	}

	expected := `{"success":false,"code":404,"message":"Not found"}`
	if rr.Body.String() != expected {
		t.Errorf("Handler returned unexpected body: got %v wanted %v",
			rr.Body.String(), expected)
	}
}

func TestMethodNotAllowed(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		res := resp.NewResponse(w)
		res.MethodNotAllowed("Method not allowed")
	})
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusMethodNotAllowed {
		t.Errorf("Handler returned wrong status code: got %v wanted %v",
			status, http.StatusMethodNotAllowed)
	}

	expected := `{"success":false,"code":405,"message":"Method not allowed"}`
	if rr.Body.String() != expected {
		t.Errorf("Handler returned unexpected body: got %v wanted %v",
			rr.Body.String(), expected)
	}
}

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
			status, http.StatusUnprocessableEntity)
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
			status, http.StatusConflict)
	}

	expected := `{"success":false,"code":409,"message":"An error occured"}`
	if rr.Body.String() != expected {
		t.Errorf("Handler returned unexpected body: got %v wanted %v",
			rr.Body.String(), expected)
	}
}

func TestInternalServerError(t *testing.T) {
	req, err := http.NewRequest("POST", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		res := resp.NewResponse(w)
		res.InternalServerError("An unexpected error occured")
	})
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusInternalServerError {
		t.Errorf("Handler returned wrong status code: got %v wanted %v",
			status, http.StatusInternalServerError)
	}

	expected := `{"success":false,"code":500,"message":"An unexpected error occured"}`
	if rr.Body.String() != expected {
		t.Errorf("Handler returned unexpected body: got %v wanted %v",
			rr.Body.String(), expected)
	}
}
