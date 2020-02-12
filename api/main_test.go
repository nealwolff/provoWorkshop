package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/nealwolff/provoWorkshop/handlers"
)

func TestPass(t *testing.T) {
	req, err := http.NewRequest(http.MethodGet, "/route", nil)

	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handlers.BasicHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Fatalf("Handler returned the wrong status code, got %v wanted %v", status, http.StatusOK)
	}

	expected := `{"key":"Hello World"}`

	if rr.Body.String() != expected {
		t.Fatalf("handler returned unexpected body: Got %v wanted %v", rr.Body.String(), expected)
	}
}
