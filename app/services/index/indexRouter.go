package index

import (
	"github.com/gin-gonic/gin"
)

// SetupRoutes connects the HTTP API endpoints to the handlers
func IndexRouter(r *gin.Engine) {
	r.GET("/l4_check", func(c *gin.Context) {
		c.String(200, "SUCCESS")
	})
}
