package v1

import (
	"fmt"
	"gin-vue-admin/global/response"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
	resp "gin-vue-admin/model/response"
	"gin-vue-admin/service"
	"gin-vue-admin/utils"
	"github.com/gin-gonic/gin"
)

// @Tags Deploy_Project
// @Summary 分页获取项目列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取项目列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /deploy/project/projectList [post]
func ProjectList(c *gin.Context) {
	var pageInfo request.ProjectPageInfo
	_ = c.ShouldBindJSON(&pageInfo)
	PageVerifyErr := utils.Verify(pageInfo, utils.CustomizeMap["PageVerify"])
	if PageVerifyErr != nil {
		response.FailWithMessage(PageVerifyErr.Error(), c)
		return
	}
	err, list, total := service.ProjectList(pageInfo)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("获取数据失败，%v", err), c)
	} else {
		response.OkWithData(resp.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, c)
	}
}

// @Tags Deploy_Project
// @Summary 创建项目
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DeployProject true "创建主机"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /deploy/project/projectCreate [post]
func ProjectCreate(c *gin.Context) {
	var project model.DeployProject
	_ = c.ShouldBindJSON(&project)
	projectVerify := utils.Rules{
		"Name":             {utils.NotEmpty()},
		"GitUrl":           {utils.NotEmpty()},
		"Directory":        {utils.NotEmpty()},
		"IgnoreFiles":      {utils.NotEmpty()},
		"ResourceServerId": {utils.NotEmpty()},
		"ResourceEnvId":    {utils.NotEmpty()},
	}
	ServerVerifyErr := utils.Verify(project, projectVerify)
	if ServerVerifyErr != nil {
		response.FailWithMessage(ServerVerifyErr.Error(), c)
		return
	}
	err := service.ProjectCreate(project)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("创建失败，%v", err), c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags Deploy_Project
// @Summary 更新项目
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DeployProject true "更新项目"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新项目成功"}"
// @Router /deploy/project/projectUpdate [post]
func ProjectUpdate(c *gin.Context) {
	var project model.DeployProject
	_ = c.ShouldBindJSON(&project)
	projectVerify := utils.Rules{
		"Name":             {utils.NotEmpty()},
		"GitUrl":           {utils.NotEmpty()},
		"Directory":        {utils.NotEmpty()},
		"IgnoreFiles":      {utils.NotEmpty()},
		"ResourceServerId": {utils.NotEmpty()},
		"ResourceEnvId":    {utils.NotEmpty()},
	}
	ServerVerifyErr := utils.Verify(project, projectVerify)
	if ServerVerifyErr != nil {
		response.FailWithMessage(ServerVerifyErr.Error(), c)
		return
	}
	err := service.ProjectUpdate(project)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("更新失败，%v", err), c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags Deploy_Project
// @Summary 删除项目
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.GetById true "删除项目"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除项目成功"}"
// @Router /deploy/project/projectDelete [delete]
func ProjectDelete(c *gin.Context) {
	var reqId request.GetById
	_ = c.ShouldBindJSON(&reqId)
	IdVerifyErr := utils.Verify(reqId, utils.CustomizeMap["IdVerify"])
	if IdVerifyErr != nil {
		response.FailWithMessage(IdVerifyErr.Error(), c)
		return
	}
	err := service.ProjectDelete(reqId.Id)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("删除失败，%v", err), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags Deploy_Project
// @Summary Gitlab 项目导入
// @Security ApiKeyAuth
// @Produce  application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"Gitlab 项目导入 成功！"}"
// @Router /deploy/project/projectImport [post]
func ProjectImport(c *gin.Context) {
	err := service.ProjectImport()
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("Gitlab 项目导入失败，%v", err), c)
	} else {
		response.OkWithMessage("Gitlab 项目导入成功!", c)
	}
}
