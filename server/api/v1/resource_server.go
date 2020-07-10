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

// @Tags Resource_Server
// @Summary 分页获取主机列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取主机列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /resource/server/serverList [post]
func ServerList(c *gin.Context) {
	var pageInfo request.ServerPageInfo
	_ = c.ShouldBindJSON(&pageInfo)
	PageVerifyErr := utils.Verify(pageInfo, utils.CustomizeMap["PageVerify"])
	if PageVerifyErr != nil {
		response.FailWithMessage(PageVerifyErr.Error(), c)
		return
	}
	err, list, total := service.ServerList(pageInfo)
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

// @Tags Resource_Server
// @Summary 创建主机
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ResourceEnv true "创建主机"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /resource/server/serverCreate [post]
func ServerCreate(c *gin.Context) {
	var server model.ResourceServer
	_ = c.ShouldBindJSON(&server)
	ServerVerify := utils.Rules{
		"Name":          {utils.NotEmpty()},
		"Host":          {utils.NotEmpty()},
		"Port":          {utils.NotEmpty()},
		"User":          {utils.NotEmpty()},
		"Pwd":           {utils.NotEmpty()},
		"ResourceEnvId": {utils.NotEmpty()},
	}
	ServerVerifyErr := utils.Verify(server, ServerVerify)
	if ServerVerifyErr != nil {
		response.FailWithMessage(ServerVerifyErr.Error(), c)
		return
	}
	err := service.ServerCreate(server)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("创建失败，%v", err), c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags Resource_Server
// @Summary 修改主机信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ResourceServer true "修改主机信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"修改成功"}"
// @Router /resource/server/serverUpdate [post]
func ServerUpdate(c *gin.Context) {
	var server model.ResourceServer
	_ = c.ShouldBindJSON(&server)
	ServerVerify := utils.Rules{
		"Name":          {utils.NotEmpty()},
		"Host":          {utils.NotEmpty()},
		"Port":          {utils.NotEmpty()},
		"User":          {utils.NotEmpty()},
		"Pwd":           {utils.NotEmpty()},
		"ResourceEnvId": {utils.NotEmpty()},
	}
	ApiVerifyErr := utils.Verify(server, ServerVerify)
	if ApiVerifyErr != nil {
		response.FailWithMessage(ApiVerifyErr.Error(), c)
		return
	}
	err := service.ServerUpdate(server)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("编辑失败，%v", err), c)
	} else {
		response.OkWithMessage("编辑成功", c)
	}
}

// @Tags Resource_Server
// @Summary 删除主机信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.GetById true "删除主机信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /resource/server/serverDelete [delete]
func ServerDelete(c *gin.Context) {
	var reqId request.GetById
	_ = c.ShouldBindJSON(&reqId)
	IdVerifyErr := utils.Verify(reqId, utils.CustomizeMap["IdVerify"])
	if IdVerifyErr != nil {
		response.FailWithMessage(IdVerifyErr.Error(), c)
		return
	}
	err := service.ServerDelete(reqId.Id)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("删除失败，%v", err), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags Resource_Server
// @Summary 生成平台密钥对
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.GetById true "生成平台密钥对"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /resource/server/platformCreateKey [get]
func PlatformCreateKey(c *gin.Context) {
	err := service.PlatformCreateKey()
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("生成平台密钥对创建失败，%v", err), c)
	} else {
		response.OkWithMessage("生成平台密钥对创建成功!", c)
	}
}

// @Tags Resource_Server
// @Summary 测试连接
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.GetById true "测试连接"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"连接成功"}"
// @Router /resource/server/serverConnect [delete]
func ServerConnect(c *gin.Context) {
	var reqId request.GetById
	_ = c.ShouldBindJSON(&reqId)
	IdVerifyErr := utils.Verify(reqId, utils.CustomizeMap["IdVerify"])
	if IdVerifyErr != nil {
		response.FailWithMessage(IdVerifyErr.Error(), c)
		return
	}
	err := service.ServerConnect(reqId.Id)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("测试连接失败，%v", err), c)
	} else {
		response.OkWithMessage("测试连接成功", c)
	}
}

// @Tags Resource_Server
// @Summary 推送公钥
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.GetById true "推送公钥"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"推送公钥成功"}"
// @Router /resource/server/serverPushKey [delete]
func ServerPushKey(c *gin.Context) {
	var reqId request.GetById
	_ = c.ShouldBindJSON(&reqId)
	IdVerifyErr := utils.Verify(reqId, utils.CustomizeMap["IdVerify"])
	if IdVerifyErr != nil {
		response.FailWithMessage(IdVerifyErr.Error(), c)
		return
	}
	err := service.ServerPushKey(reqId.Id)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("推送公钥失败，%v", err), c)
	} else {
		response.OkWithMessage("推送公钥成功", c)
	}
}
