package v1

import (
	"fmt"
	"gin-vue-admin/global/response"
	"gin-vue-admin/model/request"
	resp "gin-vue-admin/model/response"
	"gin-vue-admin/service"
	"gin-vue-admin/utils"
	"github.com/gin-gonic/gin"
)

// @Tags Deploy_Testing
// @Summary 分页获取提测列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取提测列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /deploy/test/testingList [post]
func TestingList(c *gin.Context) {
	var pageInfo request.PageInfo
	_ = c.ShouldBindJSON(&pageInfo)
	PageVerifyErr := utils.Verify(pageInfo, utils.CustomizeMap["PageVerify"])
	if PageVerifyErr != nil {
		response.FailWithMessage(PageVerifyErr.Error(), c)
		return
	}
	err, list, total := service.TestingList(pageInfo)
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

// @Tags Deploy_Testing
// @Summary 文件对比
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "文件对比"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"文件对比成功"}"
// @Router /deploy/test/testingContrast [post]
func TestingContrast(c *gin.Context) {
	var testting request.ContrastInfo
	_ = c.ShouldBindJSON(&testting)
	projectVerify := utils.Rules{
		"Tag":             {utils.NotEmpty()},
		"ResourceEnvId":   {utils.NotEmpty()},
		"DeployProjectId": {utils.NotEmpty()},
	}
	projectVerifyErr := utils.Verify(testting, projectVerify)
	if projectVerifyErr != nil {
		response.FailWithMessage(projectVerifyErr.Error(), c)
		return
	}
	err, list, path := service.TestingContrast(testting)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("对比失败，%v", err), c)
	} else {
		response.OkWithData(resp.ContrastResult{
			List: list,
			Path: path,
		}, c)
	}
}

// @Tags Deploy_Testing
// @Summary 提测发布
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "提测发布"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"提测发布成功"}"
// @Router /deploy/test/testingRelease [post]
func TestingRelease(c *gin.Context) {
	var testting request.TestingReleaseInfo
	_ = c.ShouldBindJSON(&testting)
	projectVerify := utils.Rules{
		"Tag":             {utils.NotEmpty()},
		"Path":            {utils.NotEmpty()},
		"ResourceEnvId":   {utils.NotEmpty()},
		"DeployProjectId": {utils.NotEmpty()},
		"Files":           {utils.NotEmpty()},
	}

	projectVerifyErr := utils.Verify(testting, projectVerify)
	if projectVerifyErr != nil {
		response.FailWithMessage(projectVerifyErr.Error(), c)
		return
	}

	err := service.TestingRelease(testting)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("提测失败，%v", err), c)
	} else {
		response.OkWithMessage("提测成功", c)
	}
}
