package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/render"
)

var mapperPrefix = "auth."

func IsLogin(c *gin.Context) {
	result := make(map[string]interface{})
	edwUser, isExist := c.Get("edwUser")
	if !isExist || edwUser == nil {
		result["IS_LOGIN"] = false

	} else {
		userId, _ := c.Get("edwUserId")
		userNm, _ := c.Get("edwUserNm")
		result["IS_LOGIN"] = true
		result["USER_ID"] = userId
		result["USER_NM"] = userNm
	}

	c.Render(http.StatusOK, render.IndentedJSON{Data: result})
}
