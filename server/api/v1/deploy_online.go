package v1

import (
	"fmt"
	"gin-vue-admin/global/response"
	"gin-vue-admin/middleware"
	"gin-vue-admin/model/request"
	resp "gin-vue-admin/model/response"
	"gin-vue-admin/service"
	"gin-vue-admin/utils"
	"github.com/gin-gonic/gin"
)

// @Tags Deploy_Online
// @Summary 分页获取上线列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取提测列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /deploy/online/onlineList [post]
func OnlineList(c *gin.Context) {
	var pageInfo request.PageInfo
	_ = c.ShouldBindJSON(&pageInfo)
	PageVerifyErr := utils.Verify(pageInfo, utils.CustomizeMap["PageVerify"])
	if PageVerifyErr != nil {
		response.FailWithMessage(PageVerifyErr.Error(), c)
		return
	}
	err, list, total := service.OnlineList(pageInfo)
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

// @Tags Deploy_Online
// @Summary 文件对比
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "文件对比"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"文件对比成功"}"
// @Router /deploy/online/onlineContrast [post]
func OnlineContrast(c *gin.Context) {
	var testting request.ContrastInfo
	_ = c.ShouldBindJSON(&testting)
	onlineVerify := utils.Rules{
		"Tag":             {utils.NotEmpty()},
		"EnvironmentId":   {utils.NotEmpty()},
		"DeployProjectId": {utils.NotEmpty()},
	}
	onlineVerifyErr := utils.Verify(testting, onlineVerify)
	if onlineVerifyErr != nil {
		response.FailWithMessage(onlineVerifyErr.Error(), c)
		return
	}

	err, list, path := service.OnlineContrast(testting)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("对比失败，%v", err), c)
	} else {
		response.OkWithData(resp.ContrastResult{
			List: list,
			Path: path,
		}, c)
	}
}

// @Tags Deploy_Online
// @Summary 生产发布
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "提测发布"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"提测发布成功"}"
// @Router /deploy/online/onlineRelease [post]
func OnlineRelease(c *gin.Context) {
	var testting request.TestingReleaseInfo
	_ = c.ShouldBindJSON(&testting)
	onlineVerify := utils.Rules{
		"Tag":             {utils.NotEmpty()},
		"Path":            {utils.NotEmpty()},
		"EnvironmentId":   {utils.NotEmpty()},
		"DeployProjectId": {utils.NotEmpty()},
		"Files":           {utils.NotEmpty()},
	}
	onlineVerifyErr := utils.Verify(testting, onlineVerify)
	if onlineVerifyErr != nil {
		response.FailWithMessage(onlineVerifyErr.Error(), c)
		return
	}
	claims, _ := middleware.NewJWT().ParseToken(c.GetHeader("x-token"))
	err := service.OnlineRelease(testting, claims)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("提测失败，%v", err), c)
	} else {
		response.OkWithMessage("提测成功!", c)
	}
}

// @Tags Deploy_Online
// @Summary 可回滚版本
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.GetById true "可回滚版本"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /deploy/online/OnlineRversion [post]
func OnlineRversion(c *gin.Context) {
	var reqId request.GetById
	_ = c.ShouldBindJSON(&reqId)
	IdVerifyErr := utils.Verify(reqId, utils.CustomizeMap["IdVerify"])
	if IdVerifyErr != nil {
		response.FailWithMessage(IdVerifyErr.Error(), c)
		return
	}
	err, list := service.OnlineRversion(reqId.Id)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("获取数据失败，%v", err), c)
	} else {
		response.OkWithData(resp.PageResult{
			List: list,
		}, c)
	}
}
