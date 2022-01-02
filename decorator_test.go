package gopatterns

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestDecorator(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := durationLogger(helloEndpoint)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Error("handler returned wrong status code")
	}
}
