import service from '@/utils/request'
// @Tags Resource_Env
// @Summary 分页获取环境列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body modelInterface.PageInfo true "分页获取环境列表"
// @Success 200 {string} json "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /resource/env/getEnvList [post]
// {
//  page     int
//	pageSize int
// }
export const envList = (data) => {
    return service({
        url: "/resource/env/envList",
        method: 'post',
        data
    })
}

// @Tags Resource_Env
// @Summary 创建环境
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body api.CreateApiParams true "创建环境"
// @Success 200 {string} json "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /resource/env/envCreate [post]
export const envCreate = (data) => {
    return service({
        url: "/resource/env/envCreate",
        method: 'post',
        data
    })
}

// @Tags Resource_Env
// @Summary 更新环境
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body true "更新环境"
// @Success 200 {string} json "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /resource/env/envUpdate [post]
export const envUpdate = (data) => {
    return service({
        url: "/resource/env/envUpdate",
        method: 'post',
        data
    })
}

// @Tags Resource_Env
// @Summary 删除环境
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body true "删除环境"
// @Success 200 {string} json "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /resource/env/envDelete [delete]
export const envDelete = (data) => {
    return service({
        url: "/resource/env/envDelete",
        method: 'delete',
        data
    })
}