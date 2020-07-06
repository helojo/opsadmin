package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitResourceServerRouter(Router *gin.RouterGroup) {
	ApiRouter := Router.Group("resource/server").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	{
		ApiRouter.POST("serverList", v1.ServerList) // 获取所有环境
	}
}
