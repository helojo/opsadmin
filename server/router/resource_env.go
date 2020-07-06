package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitResourceEnvRouter(Router *gin.RouterGroup) {
	ApiRouter := Router.Group("resource/env").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	{
		ApiRouter.POST("getEnvList", v1.GetEnvList) // 获取所有环境
		ApiRouter.POST("envCreate", v1.EnvCreate)   // 创建环境
		ApiRouter.POST("envUpdate", v1.EnvUpdate)   // 更新环境
		ApiRouter.DELETE("envDelete", v1.EnvDelete) // 删除环境
	}
}
