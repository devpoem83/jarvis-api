package services

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestDefaultRouter(t *testing.T) {
	router := DefaultRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/l4_check", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "SUCCESS", w.Body.String())
}