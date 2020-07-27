package alarm

import (
	"github.com/eduwill/jarvis-api/app/base/handler"
	"github.com/gin-gonic/gin"
)

func AlarmRouter(r *gin.Engine) {
	alarm_v1 := r.Group("alarms")
	alarm_v1.Use(handler.LoginCheckHandler())
	{
		alarm_v1.GET("/", List)
		alarm_v1.GET("/count", Count)
		alarm_v1.PUT("/:noticeIdx", Confirm)
		alarm_v1.DELETE("/:noticeIdx", Delete)
	}

	alarm_v0 := r.Group("v0/alarms")
	{
		alarm_v0.GET("/performance", ListTest)
	}
}
