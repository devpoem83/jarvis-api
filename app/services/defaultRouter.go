package services

import (
	"github.com/eduwill/jarvis-api/app/base/handler"
	"github.com/eduwill/jarvis-api/app/services/auth"
	"github.com/eduwill/jarvis-api/app/services/common/alarm"
	"github.com/eduwill/jarvis-api/app/services/ibm"
	"github.com/eduwill/jarvis-api/app/services/tracking"
	"github.com/eduwill/jarvis-api/app/services/index"
	"github.com/eduwill/jarvis-api/app/services/eco"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// SetupRoutes connects the HTTP API endpoints to the handlers
func DefaultRouter() *gin.Engine {

	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(handler.DefaultHandler()) // 핸들러

	// same as
	// config := cors.DefaultConfig()
	// config.AllowAllOrigins = true
	// r.Use(cors.New(config))
	r.Use(cors.Default())
	//r.Run()

	index.IndexRouter(r)
	auth.AuthRouter(r)
	alarm.AlarmRouter(r)
	//banner.BannerRouter(r)
	ibm.IbmRouter(r)
	tracking.TrackingRouter(r)
	eco.EcoRouter(r)

	return r
}
