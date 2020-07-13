package v1

import (
	"fmt"
	"gin-vue-admin/global/response"
	"gin-vue-admin/service"
	"github.com/gin-gonic/gin"
)

// @Tags Gitlab_Project
// @Summary Gitlab 项目导入
// @Security ApiKeyAuth
// @Produce  application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"Gitlab 项目导入 成功！"}"
// @Router /gitlab/project/projectImport [post]
func ProjectImport(c *gin.Context) {
	err := service.ProjectImport()
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("Gitlab 项目导入失败，%v", err), c)
	} else {
		response.OkWithMessage("Gitlab 项目导入成功!", c)
	}
}
