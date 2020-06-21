package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitDeployEnvRouter(Router *gin.RouterGroup) {
	ApiRouter := Router.Group("deploy/env").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	{
		ApiRouter.POST("getDeployEnvList", v1.GetDeployEnvList) // 获取所有api
	}
}
