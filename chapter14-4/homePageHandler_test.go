package home

import (
	"net/http"
	"net/http/httptest"
	"testing"
)


func TestHomePageHandler(t *testing.T) {
	res := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)

	HomePageHandler(res, req)

	if res.Code != 200 {
		t.Fatalf("Expected status to == 200, but got %d", res.Code)
	}
}

func TestHomePageHandler2(t *testing.T) {
	res := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)

	HomePageHandler(res, req)

	if res.Code != 300 {
		t.Fatalf("Expected status to == 200, but got %d", res.Code)
	}
}