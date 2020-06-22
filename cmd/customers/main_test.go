package main

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/MarianoArias/ApiGo/internal/app/customers/api"
	"github.com/stretchr/testify/assert"
)

var responseInterface map[string]interface{}

func TestHealth(t *testing.T) {
	recorder := setupRecorder("GET", "/health/", nil)

	responseInterfaceError := json.Unmarshal([]byte(recorder.Body.String()), &responseInterface)
	statusResponseInterface, statusResponseInterfaceSuccess := responseInterface["status"].(map[string]interface{})

	assert.Equal(t, http.StatusOK, recorder.Code)
	assert.Nil(t, responseInterfaceError)
	assert.True(t, statusResponseInterfaceSuccess)
	assert.Equal(t, "UP", statusResponseInterface["code"])
}

func TestCget(t *testing.T) {
	recorder := setupRecorder("GET", "/customers/", nil)

	assert.Equal(t, http.StatusOK, recorder.Code)
	assert.NotNil(t, recorder.Body)
}

func setupRecorder(method, path string, reader io.Reader) *httptest.ResponseRecorder {
	router := api.SetupRouter()

	recorder := httptest.NewRecorder()
	request, _ := http.NewRequest(method, path, reader)

	router.ServeHTTP(recorder, request)

	return recorder
}
