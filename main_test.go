package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSetupRouter(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	router.ServeHTTP(w, req)

	if w.Code != 200 {
		t.Fatalf("Failed! Expected: 200, Got: %d", w.Code)
	}
}
