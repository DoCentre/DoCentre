package main_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/docentre/docentre"
	"github.com/stretchr/testify/assert"
)

func TestCheckHealth(t *testing.T) {
	router := main.SetupRouter()

	req, _ := http.NewRequest(http.MethodGet, "/health", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert := assert.New(t)
	assert.Equal(http.StatusOK, w.Code)
	assert.Equal(`{"message":"health check success"}`, w.Body.String())
}
