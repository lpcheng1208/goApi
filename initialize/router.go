package initialize

import (
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	_ "goApi/docs"
	"goApi/global"
	"goApi/middleware"
	"goApi/router"
)

// 初始化总路由

func Routers() *gin.Engine {
	GinMode := global.SERVER_CONFIG.System.GinMode
	gin.SetMode(GinMode)
	Router := gin.New()

	middleware.InitMiddleware(Router)
	global.LOGGER.Debug("use common middleware success")

	// 生产环境不配置swag， 初始化swag
	if GinMode != gin.ReleaseMode {
		Router.GET("/docs/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
		global.LOGGER.Debug("register swagger handler")
	}
	// 方便统一添加路由组前缀 多服务器上线使用
	ApiGroup := Router.Group("/api")
	router.InitUserRouter(ApiGroup)      // 注册用户路由

	global.LOGGER.Debug("router register success")

	return Router
}
