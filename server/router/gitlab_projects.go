package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitGitlabProjectRouter(Router *gin.RouterGroup) {
	ApiRouter := Router.Group("gitlab/project").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	{
		ApiRouter.GET("projectImport", v1.ProjectImport) // Gitlab 项目导入

	}
}
