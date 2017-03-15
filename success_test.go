package respond_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	resp "github.com/nicklaw5/go-respond"
)

func TestOk(t *testing.T) {
	req := newRequest(t, "GET")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp.NewResponse(w).
			Ok(nil)
	})
	handler.ServeHTTP(rr, req)

	if err := validateStatusCode(rr.Code, http.StatusOK); err != nil {
		t.Fatal(err.Error())
	}

	if err := validateResponseBody(rr.Body.String(), ""); err != nil {
		t.Fatal(err.Error())
	}
}

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func TestCreated(t *testing.T) {
	req := newRequest(t, "POST")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		users := []User{
			{1, "Billy", "billy@example.com"},
			{2, "Joan", "joan@example.com"},
		}

		resp.NewResponse(w).
			Created(users)
	})
	handler.ServeHTTP(rr, req)

	if err := validateStatusCode(rr.Code, http.StatusCreated); err != nil {
		t.Fatal(err.Error())
	}

	expected := `[{"id":1,"name":"Billy","email":"billy@example.com"},` +
		`{"id":2,"name":"Joan","email":"joan@example.com"}]`
	if err := validateResponseBody(rr.Body.String(), expected); err != nil {
		t.Fatal(err.Error())
	}
}

func TestAccepted(t *testing.T) {
	req := newRequest(t, "POST")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		res := resp.NewResponse(w)
		res.Accepted(nil)
	})
	handler.ServeHTTP(rr, req)

	if err := validateStatusCode(rr.Code, http.StatusAccepted); err != nil {
		t.Fatal(err.Error())
	}

	if err := validateResponseBody(rr.Body.String(), ""); err != nil {
		t.Fatal(err.Error())
	}
}

func TestNoContent(t *testing.T) {
	req := newRequest(t, "POST")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		res := resp.NewResponse(w)
		res.NoContent()
	})
	handler.ServeHTTP(rr, req)

	if err := validateStatusCode(rr.Code, http.StatusNoContent); err != nil {
		t.Fatal(err.Error())
	}

	if err := validateResponseBody(rr.Body.String(), ""); err != nil {
		t.Fatal(err.Error())
	}

}
