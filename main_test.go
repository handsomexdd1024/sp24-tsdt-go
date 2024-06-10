package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/handsomexdd1024/sp24-tsdt-go/notes"
	"github.com/stretchr/testify/assert"
)

func TestSetupRouter(t *testing.T) {
	dbName := time.Now().Format("test20060102150405.db")
	router := notes.App(dbName)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Contains(t, w.Body.String(), "Start a new to-do list")
}
