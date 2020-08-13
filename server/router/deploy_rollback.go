package router

import (
	"gin-vue-admin/api/v1"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitDeployRollbackRouter(Router *gin.RouterGroup) {
	ApiRouter := Router.Group("deploy/rollback").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	{
		ApiRouter.POST("rollbackList", v1.RollbackList) // 回滚列表
	}
}
