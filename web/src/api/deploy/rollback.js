import service from '@/utils/request'
// @Tags Deploy_Rollback
// @Summary 分页获取回滚列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body modelInterface.PageInfo true "分页获取回滚列表"
// @Success 200 {string} json "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /deploy/rollback/rollbackList [post]
// {
//  page     int
//	pageSize int
// }
export const rollbackList = (data) => {
    return service({
        url: "/deploy/rollback/rollbackList",
        method: 'post',
        data
    })
}