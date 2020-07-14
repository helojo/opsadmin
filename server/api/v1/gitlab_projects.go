package v1

import (
	"fmt"
	"gin-vue-admin/global/response"
	"gin-vue-admin/model"
	resp "gin-vue-admin/model/response"
	"gin-vue-admin/service"
	"gin-vue-admin/utils"
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

// @Tags Gitlab_Project
// @Summary 获取项目分支和tag
// @Security ApiKeyAuth
// @Produce  application/json
// @Param data body request.GetApplicationId true "获取项目tag"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取项目tag 成功！"}"
// @Router /gitlab/project/projectTags [post]
func ProjectTags(c *gin.Context) {
	var project model.DeployProject
	_ = c.ShouldBindJSON(&project)
	projectVerify := utils.Rules{
		"ID": {utils.NotEmpty()},
	}
	projectVerifyErr := utils.Verify(project, projectVerify)
	if projectVerifyErr != nil {
		response.FailWithMessage(projectVerifyErr.Error(), c)
		return
	}
	taglist, err := service.ProjectTags(project)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("获取成功，%v", err), c)
	} else {
		response.OkWithData(resp.PageResult{
			List: taglist,
		}, c)
	}
}
