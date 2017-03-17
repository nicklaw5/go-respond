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
			BadRequest(&Error{400, "An error occurred"})
	})
	handler.ServeHTTP(rr, req)

	if err := validateStatusCode(rr.Code, http.StatusBadRequest); err != nil {
		t.Fatal(err.Error())
	}

	expected := `{"code":400,"message":"An error occurred"}`
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
			Forbidden(&Error{403, "Forbidden"})
	})
	handler.ServeHTTP(rr, req)

	if err := validateStatusCode(rr.Code, http.StatusForbidden); err != nil {
		t.Fatal(err.Error())
	}

	expected := `{"code":403,"message":"Forbidden"}`
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

func TestLengthRequired(t *testing.T) {
	req := newRequest(t, "POST")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		res := resp.NewResponse(w)
		res.LengthRequired(&Error{411, "Content-Type header not long enough"})
	})
	handler.ServeHTTP(rr, req)

	if err := validateStatusCode(rr.Code, http.StatusLengthRequired); err != nil {
		t.Fatal(err.Error())
	}

	expected := `{"code":411,"message":"Content-Type header not long enough"}`
	if err := validateResponseBody(rr.Body.String(), expected); err != nil {
		t.Fatal(err.Error())
	}
}

func TestPreconditionFailed(t *testing.T) {
	req := newRequest(t, "POST")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		res := resp.NewResponse(w)
		res.PreconditionFailed(&Error{412, "X-Auth-Key header is not present"})
	})
	handler.ServeHTTP(rr, req)

	if err := validateStatusCode(rr.Code, http.StatusPreconditionFailed); err != nil {
		t.Fatal(err.Error())
	}

	expected := `{"code":412,"message":"X-Auth-Key header is not present"}`
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
			InternalServerError(&Error{500, "An unexpected error occurred"})
	})
	handler.ServeHTTP(rr, req)

	if err := validateStatusCode(rr.Code, http.StatusInternalServerError); err != nil {
		t.Fatal(err.Error())
	}

	expected := `{"code":500,"message":"An unexpected error occurred"}`
	if err := validateResponseBody(rr.Body.String(), expected); err != nil {
		t.Fatal(err.Error())
	}
}

func TestNotImplemented(t *testing.T) {
	req := newRequest(t, "POST")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp.NewResponse(w).
			NotImplemented(&Error{501, "Unsupported request"})
	})
	handler.ServeHTTP(rr, req)

	if err := validateStatusCode(rr.Code, http.StatusNotImplemented); err != nil {
		t.Fatal(err.Error())
	}

	expected := `{"code":501,"message":"Unsupported request"}`
	if err := validateResponseBody(rr.Body.String(), expected); err != nil {
		t.Fatal(err.Error())
	}
}
