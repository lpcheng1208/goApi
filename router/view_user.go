package router

import (
	"github.com/gin-gonic/gin"
	v1 "goApi/api/v1"
	"goApi/middleware"
)

func InitUserRouter(Router *gin.RouterGroup) {
	UserRouter := Router.Group("user")

	NoCheckLoginRouter(UserRouter)

	CheckLoginRouter(UserRouter)

}

func CheckLoginRouter(r *gin.RouterGroup)   {
	r.Use(middleware.JWTAuth())
	{
		r.POST("update", v1.UpdateUserInfo) // 修改密码
		r.GET("info", v1.GetUserInfo)       // 获取个人信息
		r.POST("moneyChange", v1.UpdateUserMoney)
	}

}

func NoCheckLoginRouter(r *gin.RouterGroup)  {
	r.POST("login", v1.UserLogin)
}