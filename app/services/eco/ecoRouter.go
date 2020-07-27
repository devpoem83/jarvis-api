package eco

import (
	"github.com/gin-gonic/gin"
)

func EcoRouter(r *gin.Engine) {
	eco_v2 := r.Group("eco/v2")
	{
		eco_v2.GET("/templates/:itemId", Template)
		eco_v2.GET("/templates/:itemId/download", Download)
	}
}
