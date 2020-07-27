package ibm

import (
	"net/http"

	"github.com/eduwill/jarvis-api/app/base/config"
	"github.com/eduwill/jarvis-api/app/common/dbTemplate"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/render"
	"github.com/eduwill/jarvis-api/app/common"
	"strconv"
)

var mapperPrefix = "ibm."

func Services(c *gin.Context){
	params := make(map[string]interface{})
	svcCd := c.DefaultQuery("svcCd", "")
	params["svcCd"] = svcCd

	isPreview := c.Query("isPreview");
	previewSvcCd := ""
	previewBnrNo := ""

	common.Logger.Debug("isPreview : " + isPreview)
	if(isPreview == "Y"){
		previewSvcCd = c.Query("previewSvcCd")
		previewBnrNo = c.Query("previewBnrNo")
		common.Logger.Debug("previewSvcCd : " + previewSvcCd)
		common.Logger.Debug("previewBnrNo : " + previewBnrNo)
	}

	db := config.GetGanagosiDB()
	
	services, _ := dbTemplate.SelectList(db, mapperPrefix+"services", params)
	
	var svcCds string = ""
	var bnrNos string = ""

	if services != nil {
		for _, service := range services {		
			var banner map[string]interface{}
			common.Logger.Debug("compare : " + previewSvcCd == service["svcCd"].(string))
	
			common.Logger.Debug("isPreview : " + isPreview)
			common.Logger.Debug("previewSvcCd : " + previewSvcCd)
			common.Logger.Debug("service[svcCd].(string) : " + service["svcCd"].(string))
			if isPreview == "Y" && previewSvcCd == service["svcCd"].(string) {
				service["bnrNo"] = previewBnrNo
				banner, _ = dbTemplate.SelectOne(db, mapperPrefix+"preview-banner", service)	
			}else {
				banner, _ = dbTemplate.SelectOne(db, mapperPrefix+"banner", service)	
			}
			if banner != nil {
				if svcCds == "" {
					svcCds = service["svcCd"].(string)
				}else{
					svcCds = svcCds + "," + service["svcCd"].(string)
				}
		
				if bnrNos == "" {
					bnrNos = strconv.Itoa(int(banner["bnrNo"].(int64)))
				}else{
					bnrNos = bnrNos + "," + strconv.Itoa(int(banner["bnrNo"].(int64)))
				}
		
				contents, _ := dbTemplate.SelectList(db, mapperPrefix+"contents", banner)
				for _, content := range contents {
					links, _ := dbTemplate.SelectList(db, mapperPrefix+"links", content)
					content["links"] = links
				}
				banner["contents"] = contents
				service["banner"] = banner
			}
		}
	}
	
	if svcCds != "" {
		go display(svcCds, bnrNos)
		c.Render(http.StatusOK, render.IndentedJSON{Data: services})
		c.Abort()
	}else{
		c.Render(http.StatusOK, render.IndentedJSON{Data: nil})
	}
}


func Banner(c *gin.Context){
	params := make(map[string]interface{})
	bnrNo := c.Param("bnrNo")
	params["bnrNo"] = bnrNo

	db := config.GetGanagosiDB()
	
	var banner map[string]interface{}
	banner, _ = dbTemplate.SelectOne(db, mapperPrefix+"just-banner", params)
	if banner != nil {
		contents, _ := dbTemplate.SelectList(db, mapperPrefix+"contents", banner)
		for _, content := range contents {
			links, _ := dbTemplate.SelectList(db, mapperPrefix+"links", content)
			content["links"] = links
		}
		banner["contents"] = contents

		go display("", strconv.Itoa(int(banner["bnrNo"].(int64))))
		c.Render(http.StatusOK, render.IndentedJSON{Data: banner})
	}else{
		c.Render(http.StatusOK, render.IndentedJSON{Data: nil})
	}
}

func display(svcCds string, bnrNos string) {
	common.Logger.Debug("svcCds : " + svcCds)
	common.Logger.Debug("bnrNos : " + bnrNos)
	
	db := config.GetGanagosiDB()

	if svcCds != "" && bnrNos != "" {
		params := make(map[string]interface{})
		params["svcCds"] = svcCds
		params["bnrNos"] = bnrNos
	
		dbTemplate.SelectOne(db, mapperPrefix+"banner-display", params)
	}
}

func Click(c *gin.Context) {
	result := false
	
	svcCd := c.PostForm("svcCd")
	bnrNo := c.PostForm("bnrNo")

	common.Logger.Debug("svcCd : " + svcCd)
	common.Logger.Debug("bnrNo : " + bnrNo)
	
	db := config.GetGanagosiDB()


	if svcCd != "" && bnrNo != "" {
		params := make(map[string]interface{})
		params["svcCd"] = svcCd
		params["bnrNo"] = bnrNo
	
		obj, _ := dbTemplate.SelectOne(db, mapperPrefix+"banner-click", params)
		if obj["result"] != "0" {
			result = true
		}
	}
	c.Render(http.StatusOK, render.IndentedJSON{Data: result})
}