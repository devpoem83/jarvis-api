package main // import "github.com/eduwill/jarvis-api"

import (
	"strconv"

	"github.com/eduwill/jarvis-api/app/base/config"
	"github.com/eduwill/jarvis-api/app/common"
	"github.com/eduwill/jarvis-api/app/services"
)

func main() {

	// 프로파일 초기화
	config.ProfileInit()

	// Logger 초기화
	common.LoggerInit(config.GetProfile().LoggerLevel)

	// 캐시 초기화
	config.CacheInit()

	// DB 초기화
	config.DbInit()

	// 라우트 설정
	r := services.DefaultRouter()

	// 정적리소스 설정
	config.StaticInit(r)

	// server 구동
	r.Run(":" + strconv.Itoa(config.GetProfile().ServerPort))
}
