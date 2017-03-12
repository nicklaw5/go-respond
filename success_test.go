package respond_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	resp "github.com/nicklaw5/go-respond"
)

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func TestCreated(t *testing.T) {
	req, err := http.NewRequest("POST", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		res := resp.NewResponse(w)
		res.Created(&User{1, "Billy", "billy@example.com"})
	})
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("Handler returned wrong status code: got %v wanted %v",
			status, http.StatusOK)
	}

	expected := `{"success":true,"data":{"id":1,"name":"Billy","email":"billy@example.com"}}`
	if rr.Body.String() != expected {
		t.Errorf("Handler returned unexpected body: got %v wanted %v",
			rr.Body.String(), expected)
	}
}
