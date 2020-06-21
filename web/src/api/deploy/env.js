import service from '@/utils/request'
// @Tags Deploy_Env
// @Summary 分页获取环境列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body modelInterface.PageInfo true "分页获取环境列表"
// @Success 200 {string} json "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /deploy/env/getDeployEnvList [post]
// {
//  page     int
//	pageSize int
// }
export const getDeployEnvList = (data) => {
    return service({
        url: "/deploy/env/getDeployEnvList",
        method: 'post',
        data
    })
}