package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitDeployProjectRouter(Router *gin.RouterGroup) {
	ApiRouter := Router.Group("deploy/project").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	{
		ApiRouter.POST("projectList", v1.ProjectList) //   项目列表
	}
}
