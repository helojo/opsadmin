package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitResourceServerRouter(Router *gin.RouterGroup) {
	ApiRouter := Router.Group("resource/server").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	{
		ApiRouter.POST("serverList", v1.ServerList)              // 主机列表
		ApiRouter.POST("serverCreate", v1.ServerCreate)          // 主机创建
		ApiRouter.POST("serverUpdate", v1.ServerUpdate)          // 主机信息更新
		ApiRouter.DELETE("serverDelete", v1.ServerDelete)        // 主机信息删除
		ApiRouter.GET("platformCreateKey", v1.PlatformCreateKey) // 平台密钥对创建
	}
}
