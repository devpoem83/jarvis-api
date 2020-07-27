package eco

import (
	"net/http"

	"github.com/eduwill/jarvis-api/app/base/config"
	"github.com/eduwill/jarvis-api/app/common/dbTemplate"
	"github.com/gin-gonic/gin"
	"fmt"
	"github.com/yosssi/gohtml"
)

var mapperPrefix = "eco."

func Download(c *gin.Context){

	params := make(map[string]interface{})
	itemId := c.Param("itemId")
	params["itemId"] = itemId

	db := config.GetGanagosiDB()
	
	var object map[string]interface{}
	object, _ = dbTemplate.SelectOne(db, mapperPrefix+"template", params)
	
	title := gohtml.Format(object["title"].(string)) 
	html := gohtml.Format(object["template"].(string))

	c.Writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=%s", title + ".html"))
    c.Writer.Header().Add("Content-Type", "application/octet-stream")
	c.String(http.StatusOK, html)
}

func Template(c *gin.Context){
	params := make(map[string]interface{})
	itemId := c.Param("itemId")
	params["itemId"] = itemId

	db := config.GetGanagosiDB()
	
	var object map[string]interface{}
	object, _ = dbTemplate.SelectOne(db, mapperPrefix+"template", params)
	
	html := gohtml.Format(object["template"].(string))
	c.String(http.StatusOK, html)
}