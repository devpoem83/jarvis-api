package tracking

import (
	"net/http"
	"github.com/eduwill/jarvis-api/app/base/config"
	"github.com/eduwill/jarvis-api/app/common/dbTemplate"
	"github.com/gin-gonic/gin"
	"github.com/eduwill/jarvis-api/app/common"
)

var mapperPrefix = "traking."

type TrackingObj struct {
	UserId		string `form:"userId" json:"userId"`
	PageUrl		string `form:"pageUrl" json:"pageUrl"`
	RefererUrl	string `form:"refererUrl" json:"refererUrl"`
	AuthCookie	string `form:"authCookie" json:"authCookie"`
	UserIp		string `form:"userIp" json:"userIp"`
	ServerIp	string `form:"serverIp" json:"serverIp"`
	Agent		string `form:"agent" json:"agent"`
}

/*
	배너 목록조회
*/
func Visit(c *gin.Context) {
	var obj TrackingObj
	if c.ShouldBind(&obj) == nil {
		go setTracking(&obj)	// 병렬로 처리
	}else{
		common.Logger.Info("전송된 데이터가 없어 처리하지 않음!")
	}
	c.JSON(http.StatusOK, gin.H{})
}

func setTracking(obj *TrackingObj) {
	params := make(map[string]interface{})
	db := config.GetLogDB()

	params["userId"] = obj.UserId
	params["pageUrl"] = obj.PageUrl
	params["refererUrl"] = obj.RefererUrl
	params["authCookie"] = obj.AuthCookie
	params["userIp"] = obj.UserIp
	params["serverIp"] = obj.ServerIp
	params["agent"] = obj.Agent

	if obj.UserId != "" && obj.PageUrl != "" {
		dbTemplate.SelectOne(db, mapperPrefix+"visit", params)
		common.Logger.Info("트래킹 정보입력 성공!")
	}else{
		common.Logger.Info("필수정보가 부족하여 처리하지 않음!")
	}
}