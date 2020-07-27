package handler

import (
	"github.com/eduwill/jarvis-api/app/common"
	"github.com/eduwill/jarvis-api/app/common/utils/parser"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
	"time"
)

var ConstEdwUser = "edwUser"
var ConstEdwUserId = "edwUserId"
var ConstEdwUserNm = "edwUserNm"

func DefaultHandler() gin.HandlerFunc {
	return func(c *gin.Context) {

		// before Process
		t := time.Now()

		uri := c.Request.RequestURI
		if uri == "/l4_check" || strings.Index(uri, ".js") > -1 || strings.Index(uri, ".css") > -1 || strings.Index(uri, ".ico") > -1 {
			return
		}
		loggerHandler(c)
		baseHandler(c)

		c.Next()

		status := c.Writer.Status()
		common.Logger.Info("STATUS   : " + strconv.Itoa(status))

		latency := time.Since(t)
		common.Logger.Info("LATENCY  : ", latency)
	}
}

func loggerHandler(c *gin.Context) {
	uri := c.Request.RequestURI
	ip := c.ClientIP()
	agent := c.Request.UserAgent()
	referer := c.Request.Referer()
	method := c.Request.Method

	common.Logger.Info("IP       : ", ip)
	common.Logger.Info("AGENT    : ", agent)
	common.Logger.Info("URI      : ", uri)
	common.Logger.Info("METHOD   : ", method)
	common.Logger.Info("REFERER  : ", referer)
}

func baseHandler(c *gin.Context) {
	edwUser, err := parser.GetEdwUser(c)
	if err != nil {
		c.Set(ConstEdwUser, nil)
		c.Set("LOGIN", false)
	} else {
		c.Set(ConstEdwUser, edwUser)
		c.Set(ConstEdwUserId, edwUser.UserId)
		c.Set(ConstEdwUserNm, edwUser.UserNm)
		c.Set("LOGIN", true)
	}
}

func LoginCheckHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		edwUser, isExist := c.Get(ConstEdwUser)
		if !isExist || edwUser == nil {
			c.JSON(http.StatusUnauthorized, gin.H{})
			c.Abort()
		}
	}
}
