package respond_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	resp "github.com/nicklaw5/go-respond"
)

func TestBadRequest(t *testing.T) {
	req := newRequest(t, "GET")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		res := resp.NewResponse(w)
		res.BadRequest("Bad request")
	})
	handler.ServeHTTP(rr, req)

	if err := validateStatusCode(rr.Code, http.StatusBadRequest); err != nil {
		t.Fatal(err.Error())
	}

	expected := `{"success":false,"code":400,"message":"Bad request"}`
	if err := validateResponseBody(rr.Body.String(), expected); err != nil {
		t.Fatal(err.Error())
	}
}

func TestUnauthorized(t *testing.T) {
	req := newRequest(t, "GET")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		res := resp.NewResponse(w)
		res.Unauthorized("Unauthorized")
	})
	handler.ServeHTTP(rr, req)

	if err := validateStatusCode(rr.Code, http.StatusUnauthorized); err != nil {
		t.Fatal(err.Error())
	}

	expected := `{"success":false,"code":401,"message":"Unauthorized"}`
	if err := validateResponseBody(rr.Body.String(), expected); err != nil {
		t.Fatal(err.Error())
	}
}

func TestForbidden(t *testing.T) {
	req := newRequest(t, "GET")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		res := resp.NewResponse(w)
		res.Forbidden("Forbidden")
	})
	handler.ServeHTTP(rr, req)

	if err := validateStatusCode(rr.Code, http.StatusForbidden); err != nil {
		t.Fatal(err.Error())
	}

	expected := `{"success":false,"code":403,"message":"Forbidden"}`
	if err := validateResponseBody(rr.Body.String(), expected); err != nil {
		t.Fatal(err.Error())
	}
}

func TestNotFound(t *testing.T) {
	req := newRequest(t, "GET")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		res := resp.NewResponse(w)
		res.NotFound("Not found")
	})
	handler.ServeHTTP(rr, req)

	if err := validateStatusCode(rr.Code, http.StatusNotFound); err != nil {
		t.Fatal(err.Error())
	}

	expected := `{"success":false,"code":404,"message":"Not found"}`
	if err := validateResponseBody(rr.Body.String(), expected); err != nil {
		t.Fatal(err.Error())
	}
}

func TestMethodNotAllowed(t *testing.T) {
	req := newRequest(t, "GET")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		res := resp.NewResponse(w)
		res.MethodNotAllowed("Method not allowed")
	})
	handler.ServeHTTP(rr, req)

	if err := validateStatusCode(rr.Code, http.StatusMethodNotAllowed); err != nil {
		t.Fatal(err.Error())
	}

	expected := `{"success":false,"code":405,"message":"Method not allowed"}`
	if err := validateResponseBody(rr.Body.String(), expected); err != nil {
		t.Fatal(err.Error())
	}
}

func TestUnprocessableEntity(t *testing.T) {
	req := newRequest(t, "POST")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		res := resp.NewResponse(w)
		res.UnprocessableEntity("An error occured")
	})
	handler.ServeHTTP(rr, req)

	if err := validateStatusCode(rr.Code, http.StatusUnprocessableEntity); err != nil {
		t.Fatal(err.Error())
	}

	expected := `{"success":false,"code":422,"message":"An error occured"}`
	if err := validateResponseBody(rr.Body.String(), expected); err != nil {
		t.Fatal(err.Error())
	}
}

func TestConflict(t *testing.T) {
	req := newRequest(t, "POST")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		res := resp.NewResponse(w)
		res.Conflict("An error occured")
	})
	handler.ServeHTTP(rr, req)

	if err := validateStatusCode(rr.Code, http.StatusConflict); err != nil {
		t.Fatal(err.Error())
	}

	expected := `{"success":false,"code":409,"message":"An error occured"}`
	if err := validateResponseBody(rr.Body.String(), expected); err != nil {
		t.Fatal(err.Error())
	}
}

func TestInternalServerError(t *testing.T) {
	req := newRequest(t, "POST")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		res := resp.NewResponse(w)
		res.InternalServerError("An unexpected error occured")
	})
	handler.ServeHTTP(rr, req)

	if err := validateStatusCode(rr.Code, http.StatusInternalServerError); err != nil {
		t.Fatal(err.Error())
	}

	expected := `{"success":false,"code":500,"message":"An unexpected error occured"}`
	if err := validateResponseBody(rr.Body.String(), expected); err != nil {
		t.Fatal(err.Error())
	}
}
