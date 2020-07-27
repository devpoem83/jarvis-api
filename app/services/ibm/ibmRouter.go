package ibm

import (
	"github.com/gin-gonic/gin"
)

func IbmRouter(r *gin.Engine) {
	ibm_v1 := r.Group("ibm/v1")
	{
		ibm_v1.GET("/services", Services)
		ibm_v1.GET("/banner/:bnrNo", Banner)
		ibm_v1.POST("/click", Click)
	}
}
