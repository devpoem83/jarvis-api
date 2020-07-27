package tracking

import (
	"github.com/gin-gonic/gin"
)

func TrackingRouter(r *gin.Engine) {
	v1 := r.Group("tracking")
	{
		v1.POST("/visit", Visit)
	}
}
