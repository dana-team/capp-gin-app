package server

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

const envValue = "test"

func TestHandleEchoWithoutEnv(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/echo?var=test", nil)
	if err != nil {
		t.Fatal(err)
	}

	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, fmt.Sprintf(echoMessage, "test", ""), w.Body.String())
}

func TestHandleEchoWithEnv(t *testing.T) {
	router := setupRouter()
	err := os.Setenv(envKey, envValue)
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/echo?var=test", nil)
	if err != nil {
		t.Fatal(err)
	}

	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, fmt.Sprintf(echoMessage, "test", fmt.Sprintf(envMessage, envValue)), w.Body.String())
	err = os.Unsetenv(envKey)
	if err != nil {
		t.Fatal(err)
	}
}

func TestHandleEchoWithoutQuery(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/echo", nil)
	if err != nil {
		t.Fatal(err)
	}

	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, errNoQueryValue, w.Body.String())
}

func TestHandleHealthWithoutEnv(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/health", nil)
	if err != nil {
		t.Fatal(err)
	}

	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, fmt.Sprintf(healthMessage, ""), w.Body.String())
}

func TestHandleHealthWithEnv(t *testing.T) {
	router := setupRouter()
	err := os.Setenv(envKey, envValue)
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/health", nil)
	if err != nil {
		t.Fatal(err)
	}

	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, fmt.Sprintf(healthMessage, fmt.Sprintf(envMessage, envValue)), w.Body.String())
	err = os.Unsetenv(envKey)
	if err != nil {
		t.Fatal(err)
	}
}
