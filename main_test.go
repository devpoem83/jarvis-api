package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/eduwill/jarvis-api/app/services"
)

func TestMain(t *testing.T) {
	router := services.DefaultRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/l4_check", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "SUCCESS", w.Body.String())
}