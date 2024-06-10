package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/handsomexdd1024/sp24-tsdt-go/notes"
	"github.com/stretchr/testify/assert"
)

func TestSetupRouter(t *testing.T) {
	router := notes.App("./templates", "./test.db")

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Contains(t, w.Body.String(), "To-Do List")
}
