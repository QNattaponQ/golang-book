package home

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"encoding/json"
	"strings"
)

func TestJsonHandler(t *testing.T) {
	res := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/json",strings.NewReader(`{"first_name":"espresso", "last_name": "longshot","email": "espresso@speedy.coffee"}`))
	home := HomePageHandler{}
	home.ServeHTTP(res, req)
	if res.Code != 201 {
		t.Fatalf("Expected status to == 201, but got %d", res.Code)
	}
	user := new(User)
	json.NewDecoder(res.Body).Decode(user)
	if user.FirstName != "espresso" {
		t.Fatalf("Expected user first to == espresso, but got %s",
			user.FirstName)
	}
}
