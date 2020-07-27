package auth

import (
	"github.com/gin-gonic/gin"
)

// SetupRoutes connects the HTTP API endpoints to the handlers
func AuthRouter(r *gin.Engine) {
	auth := r.Group("auth")
	{
		auth.GET("/login/check", IsLogin)
	}
}
