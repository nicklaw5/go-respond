package respond_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	resp "github.com/nicklaw5/go-respond"
)

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func TestBadRequest(t *testing.T) {
	req := newRequest(t, "GET")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp.NewResponse(w).
			BadRequest(&Error{400, "An error occured"})
	})
	handler.ServeHTTP(rr, req)

	if err := validateStatusCode(rr.Code, http.StatusBadRequest); err != nil {
		t.Fatal(err.Error())
	}

	expected := `{"code":400,"message":"An error occured"}`
	if err := validateResponseBody(rr.Body.String(), expected); err != nil {
		t.Fatal(err.Error())
	}
}

func TestUnauthorized(t *testing.T) {
	req := newRequest(t, "GET")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp.NewResponse(w).
			Unauthorized(&Error{401, "Unauthorized"})
	})
	handler.ServeHTTP(rr, req)

	if err := validateStatusCode(rr.Code, http.StatusUnauthorized); err != nil {
		t.Fatal(err.Error())
	}

	expected := `{"code":401,"message":"Unauthorized"}`
	if err := validateResponseBody(rr.Body.String(), expected); err != nil {
		t.Fatal(err.Error())
	}
}

func TestForbidden(t *testing.T) {
	req := newRequest(t, "GET")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp.NewResponse(w).
			Forbidden(&Error{401, "Forbidden"})
	})
	handler.ServeHTTP(rr, req)

	if err := validateStatusCode(rr.Code, http.StatusForbidden); err != nil {
		t.Fatal(err.Error())
	}

	expected := `{"code":401,"message":"Forbidden"}`
	if err := validateResponseBody(rr.Body.String(), expected); err != nil {
		t.Fatal(err.Error())
	}
}

func TestNotFound(t *testing.T) {
	req := newRequest(t, "GET")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp.NewResponse(w).
			NotFound(&Error{404, "Not found"})
	})
	handler.ServeHTTP(rr, req)

	if err := validateStatusCode(rr.Code, http.StatusNotFound); err != nil {
		t.Fatal(err.Error())
	}

	expected := `{"code":404,"message":"Not found"}`
	if err := validateResponseBody(rr.Body.String(), expected); err != nil {
		t.Fatal(err.Error())
	}
}

func TestMethodNotAllowed(t *testing.T) {
	req := newRequest(t, "GET")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp.NewResponse(w).
			MethodNotAllowed(&Error{405, "Method not allowed"})
	})
	handler.ServeHTTP(rr, req)

	if err := validateStatusCode(rr.Code, http.StatusMethodNotAllowed); err != nil {
		t.Fatal(err.Error())
	}

	expected := `{"code":405,"message":"Method not allowed"}`
	if err := validateResponseBody(rr.Body.String(), expected); err != nil {
		t.Fatal(err.Error())
	}
}

func TestConflict(t *testing.T) {
	req := newRequest(t, "POST")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		res := resp.NewResponse(w)
		res.Conflict(&Error{409, "Username already take"})
	})
	handler.ServeHTTP(rr, req)

	if err := validateStatusCode(rr.Code, http.StatusConflict); err != nil {
		t.Fatal(err.Error())
	}

	expected := `{"code":409,"message":"Username already take"}`
	if err := validateResponseBody(rr.Body.String(), expected); err != nil {
		t.Fatal(err.Error())
	}
}

func TestUnprocessableEntity(t *testing.T) {
	req := newRequest(t, "POST")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp.NewResponse(w).
			UnprocessableEntity(&Error{422, "Unprocessable entity"})
	})
	handler.ServeHTTP(rr, req)

	if err := validateStatusCode(rr.Code, http.StatusUnprocessableEntity); err != nil {
		t.Fatal(err.Error())
	}

	expected := `{"code":422,"message":"Unprocessable entity"}`
	if err := validateResponseBody(rr.Body.String(), expected); err != nil {
		t.Fatal(err.Error())
	}
}

func TestInternalServerError(t *testing.T) {
	req := newRequest(t, "POST")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp.NewResponse(w).
			InternalServerError(&Error{500, "An unexpected error occured"})
	})
	handler.ServeHTTP(rr, req)

	if err := validateStatusCode(rr.Code, http.StatusInternalServerError); err != nil {
		t.Fatal(err.Error())
	}

	expected := `{"code":500,"message":"An unexpected error occured"}`
	if err := validateResponseBody(rr.Body.String(), expected); err != nil {
		t.Fatal(err.Error())
	}
}
