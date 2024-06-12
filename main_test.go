package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	. "github.com/handsomexdd1024/sp24-tsdt-go/router"
	"github.com/stretchr/testify/assert"
)

func TestApp(t *testing.T) {
	dbName := time.Now().Format("test20060102150405.db")
	router := App(dbName)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Contains(t, w.Body.String(), "Start a new to-do list")

	// send form data 'description=233' at /new to create a new list
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("POST", "/new", nil)
	req.PostForm = make(map[string][]string)
	req.PostForm.Add("description", "233")
	router.ServeHTTP(w, req)
	// expect return 302
	assert.Equal(t, 302, w.Code)
	// expect redirect to /int
	redirUrl, _ := w.Result().Location()
	assert.Equal(t, "/1/", redirUrl.Path)

	// retrieve the newly created list
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", redirUrl.Path, nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
	assert.Contains(t, w.Body.String(), "Your to-do list")
	assert.Contains(t, w.Body.String(), "233")

	// add a new item to the list
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("POST", "/1/new", nil)
	req.PostForm = make(map[string][]string)
	req.PostForm.Add("description", "114514")
	router.ServeHTTP(w, req)
	assert.Equal(t, 302, w.Code)
	redirUrl, _ = w.Result().Location()
	assert.Equal(t, "/1/", redirUrl.Path)

	// retrieve the list again
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/1/", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
	assert.Contains(t, w.Body.String(), "Your to-do list")
	assert.Contains(t, w.Body.String(), "233")
	assert.Contains(t, w.Body.String(), "114514")

	// start another list
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("POST", "/new", nil)
	req.PostForm = make(map[string][]string)
	req.PostForm.Add("description", "1919810")
	router.ServeHTTP(w, req)
	assert.Equal(t, 302, w.Code)
	redirUrl, _ = w.Result().Location()
	assert.Equal(t, "/2/", redirUrl.Path)

	// retrieve the new list
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/2/", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
	assert.Contains(t, w.Body.String(), "Your to-do list")
	assert.Contains(t, w.Body.String(), "1919810")

	// retrieve the previous list
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/1/", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
	assert.Contains(t, w.Body.String(), "Your to-do list")
	assert.Contains(t, w.Body.String(), "233")
	assert.Contains(t, w.Body.String(), "114514")

}
