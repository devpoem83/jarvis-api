package alarm

import (
	"net/http"

	"github.com/eduwill/jarvis-api/app/base/config"
	"github.com/eduwill/jarvis-api/app/common/dbTemplate"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/render"
)

var mapperPrefix = "alarm."

/*
	알림 카운트
*/
func Count(c *gin.Context) {
	params := make(map[string]interface{})

	userId, _ := c.Get("edwUserId")
	gubun := c.DefaultQuery("type", "")

	params["userId"] = userId
	params["type"] = gubun

	db := config.GetGanagosiDB()
	count, _ := dbTemplate.SelectOne(db, mapperPrefix+"count", params)
	c.Render(http.StatusOK, render.IndentedJSON{Data: count})
}

/*
	알림 목록조회
*/
func List(c *gin.Context) {
	params := make(map[string]interface{})

	userId, _ := c.Get("edwUserId")
	page := c.DefaultQuery("page", "1")
	pageSize := c.DefaultQuery("pageSize", "10")
	searchTitle := c.DefaultQuery("searchTitle", "")
	searchCont := c.DefaultQuery("searchTitle", "")

	params["userId"] = userId
	params["page"] = page
	params["pageSize"] = pageSize
	params["searchTitle"] = searchTitle
	params["searchCont"] = searchCont

	db := config.GetGanagosiDB()
	list, _ := dbTemplate.SelectList(db, mapperPrefix+"list", params)
	c.Render(http.StatusOK, render.IndentedJSON{Data: list})
}

func ListTest(c *gin.Context) {
	params := make(map[string]interface{})

	userId := "edutest06"
	page := c.DefaultQuery("page", "1")
	pageSize := c.DefaultQuery("pageSize", "10")
	searchTitle := c.DefaultQuery("searchTitle", "")
	searchCont := c.DefaultQuery("searchTitle", "")

	params["userId"] = userId
	params["page"] = page
	params["pageSize"] = pageSize
	params["searchTitle"] = searchTitle
	params["searchCont"] = searchCont

	db := config.GetGanagosiDB()
	list, _ := dbTemplate.SelectList(db, mapperPrefix+"list", params)
	c.Render(http.StatusOK, render.IndentedJSON{Data: list})
}

/*
	알림 수신처리
*/
func Confirm(c *gin.Context) {
	params := make(map[string]interface{})

	userId, _ := c.Get("edwUserId")
	noticeIdx := c.Param("noticeIdx")
	recvTypeCd := c.DefaultQuery("recvTypeCd", "")

	params["userId"] = userId
	params["noticeIdx"] = noticeIdx
	params["recvTypeCd"] = recvTypeCd

	result := false
	db := config.GetGanagosiDB()
	_, rowsAffected := dbTemplate.Exec(db, mapperPrefix+"confirm", params)
	if rowsAffected > 0 {
		result = true
	}
	c.Render(http.StatusOK, render.IndentedJSON{Data: result})
}

/*
	알림 삭제처리
*/
func Delete(c *gin.Context) {
	params := make(map[string]interface{})

	userId, _ := c.Get("edwUserId")
	noticeIdx := c.Param("noticeIdx")
	delType := c.DefaultQuery("delType", "")

	params["userId"] = userId
	params["noticeIdx"] = noticeIdx
	params["delType"] = delType

	result := false
	db := config.GetGanagosiDB()
	_, rowsAffected := dbTemplate.Exec(db, mapperPrefix+"delete", params)
	if rowsAffected > 0 {
		result = true
	}
	c.Render(http.StatusOK, render.IndentedJSON{Data: result})
}
