package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitDeployProjectRouter(Router *gin.RouterGroup) {
	ApiRouter := Router.Group("deploy/project").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	{
		ApiRouter.POST("projectList", v1.ProjectList)       // 项目列表
		ApiRouter.POST("projectCreate", v1.ProjectCreate)   // 项目创建
		ApiRouter.POST("projectUpdate", v1.ProjectUpdate)   // 项目更新
		ApiRouter.DELETE("projectDelete", v1.ProjectDelete) // 项目删除
		ApiRouter.GET("projectImport", v1.ProjectImport)    // Gitlab 项目导入

	}
}
