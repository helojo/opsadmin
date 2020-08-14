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

// @Tags Resource_Env
// @Summary 分页获取环境列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取环境列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /resource/env/envList [post]
func EnvList(c *gin.Context) {
	var pageInfo request.EnvironmentPageInfo
	_ = c.ShouldBindJSON(&pageInfo)
	PageVerifyErr := utils.Verify(pageInfo, utils.CustomizeMap["PageVerify"])
	if PageVerifyErr != nil {
		response.FailWithMessage(PageVerifyErr.Error(), c)
		return
	}
	err, list, total := service.EnvList(pageInfo)
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

// @Tags Resource_Env
// @Summary 创建环境
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Environment true "创建环境"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /resource/env/envCreate [post]
func EnvCreate(c *gin.Context) {
	var env model.Environment
	_ = c.ShouldBindJSON(&env)
	EnvVerify := utils.Rules{
		"Name":     {utils.NotEmpty()},
		"EnvLabel": {utils.NotEmpty()},
	}
	EnvVerifyErr := utils.Verify(env, EnvVerify)
	if EnvVerifyErr != nil {
		response.FailWithMessage(EnvVerifyErr.Error(), c)
		return
	}
	err := service.EnvCreate(env)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("创建失败，%v", err), c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags Resource_Env
// @Summary 修改环境
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Environment true "修改环境"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /resource/env/envUpdate [post]
func EnvUpdate(c *gin.Context) {
	var env model.Environment
	_ = c.ShouldBindJSON(&env)
	ApiVerify := utils.Rules{
		"Name":     {utils.NotEmpty()},
		"EnvLabel": {utils.NotEmpty()},
	}
	ApiVerifyErr := utils.Verify(env, ApiVerify)
	if ApiVerifyErr != nil {
		response.FailWithMessage(ApiVerifyErr.Error(), c)
		return
	}
	err := service.EnvUpdate(env)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("修改数据失败，%v", err), c)
	} else {
		response.OkWithMessage("修改数据成功", c)
	}
}

// @Tags Resource_Env
// @Summary 删除环境
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.GetById true "删除环境"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /resource/env/envDelete [delete]
func EnvDelete(c *gin.Context) {
	var reqId request.GetById
	_ = c.ShouldBindJSON(&reqId)
	IdVerifyErr := utils.Verify(reqId, utils.CustomizeMap["IdVerify"])
	if IdVerifyErr != nil {
		response.FailWithMessage(IdVerifyErr.Error(), c)
		return
	}
	err := service.EnvDelete(reqId.Id)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("删除失败，%v", err), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}
