package server

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHandleEcho(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/echo?var=test", nil)
	if err != nil {
		t.Fatal(err)
	}

	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "echoing: test", w.Body.String())
}

func TestHandleHealth(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/health", nil)
	if err != nil {
		t.Fatal(err)
	}

	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "OK", w.Body.String())
}
