package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitDeployOnlineRouter(Router *gin.RouterGroup) {
	ApiRouter := Router.Group("deploy/online").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	{
		ApiRouter.POST("onlineList", v1.OnlineList)         // 上线列表
		ApiRouter.POST("onlineContrast", v1.OnlineContrast) // 上线对比
		ApiRouter.POST("onlineCreate", v1.OnlineCreate)     // 上线提交
		ApiRouter.POST("onlineRversion", v1.OnlineRversion) // 获取可回退版本
	}
}
