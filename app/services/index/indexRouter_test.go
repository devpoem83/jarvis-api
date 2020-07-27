package index

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/gin-gonic/gin"	
)

func TestIndexRouter(t *testing.T) {
	r := gin.New()
	IndexRouter(r)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/l4_check", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "SUCCESS", w.Body.String())
}

