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

// @Tags Deploy_Rollback
// @Summary 分页获取回滚列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取回滚列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /deploy/rollback/rollbackList [post]
func RollbackList(c *gin.Context) {
	var pageInfo request.PageInfo
	_ = c.ShouldBindJSON(&pageInfo)
	PageVerifyErr := utils.Verify(pageInfo, utils.CustomizeMap["PageVerify"])
	if PageVerifyErr != nil {
		response.FailWithMessage(PageVerifyErr.Error(), c)
		return
	}
	err, list, total := service.RollbackList(pageInfo)
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

// @Tags Deploy_Rollback
// @Summary 文件对比
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "文件对比"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"文件对比成功"}"
// @Router /deploy/rollback/rollbackContrast [post]
func RollbackContrast(c *gin.Context) {
	var rollback request.RollbackContrast
	_ = c.ShouldBindJSON(&rollback)
	RollbackVerify := utils.Rules{
		"Version":         {utils.NotEmpty()},
		"Describe":        {utils.NotEmpty()},
		"EnvironmentId":   {utils.NotEmpty()},
		"DeployProjectId": {utils.NotEmpty()},
	}
	fmt.Println(rollback)
	rollbackVerifyErr := utils.Verify(rollback, RollbackVerify)
	if rollbackVerifyErr != nil {
		response.FailWithMessage(rollbackVerifyErr.Error(), c)
		return
	}
	response.OkWithMessage("到这里了", c)
	//err, list, path := service.TestingContrast(testting)
	//if err != nil {
	//	response.FailWithMessage(fmt.Sprintf("对比失败，%v", err), c)
	//} else {
	//	response.OkWithData(resp.ContrastResult{
	//		List: list,
	//		Path: path,
	//	}, c)
	//}
}
