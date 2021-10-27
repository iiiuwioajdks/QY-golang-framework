package initialize

import (
	"github.com/gin-gonic/gin"
	"jiaoshoujia/web/middleware"
	"jiaoshoujia/web/router"
)

// InitRouter 路由初始化
func InitRouter() *gin.Engine {
	Router := gin.New()
	Router.Use(GinLogger(), GinRecovery(true), middleware.Cors())

	ApiRouter := Router.Group("/uapply/v1")
	router.InitUserRouter(ApiRouter)
	router.InitAdminRouter(ApiRouter)
	return Router
}
