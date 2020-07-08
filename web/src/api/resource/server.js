import service from '@/utils/request'
// @Tags Resource_Server
// @Summary 分页获取主机列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body modelInterface.PageInfo true "分页获取主机列表"
// @Success 200 {string} json "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /resource/server/serverList [post]
// {
//  page     int
//	pageSize int
// }
export const serverList = (data) => {
    return service({
        url: "/resource/server/serverList",
        method: 'post',
        data
    })
}

// @Tags Resource_Server
// @Summary 新增主机
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body api.envCreate true "新增主机"
// @Success 200 {string} json "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /resource/server/serverCreate [post]
export const serverCreate = (data) => {
    return service({
        url: "/resource/server/serverCreate",
        method: 'post',
        data
    })
}

// @Tags Resource_Server
// @Summary 更新主机信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body true "更新环境"
// @Success 200 {string} json "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /resource/server/serverUpdate [post]
export const serverUpdate = (data) => {
    return service({
        url: "/resource/server/serverUpdate",
        method: 'post',
        data
    })
}

// @Tags Resource_Server
// @Summary 删除主机信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body true "删除主机信息"
// @Success 200 {string} json "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /resource/server/serverDelete [delete]
export const serverDelete = (data) => {
    return service({
        url: "/resource/server/serverDelete",
        method: 'delete',
        data
    })
}

// @Tags Resource_Server
// @Summary 平台生成密钥对
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} json "{"success":true,"data":{},"msg":"生成成功"}"
// @Router /resource/server/platformCreateKey [delete]
export const platformCreateKey = () => {
    return service({
        url: "/resource/server/platformCreateKey",
        method: 'get'
    })
}

// @Tags Resource_Server
// @Summary 连接测试
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body true "连接测试成功"
// @Success 200 {string} json "{"success":true,"data":{},"msg":"连接测试成功"}"
// @Router /resource/server/serverConnect [post]
export const serverConnect = (data) => {
    return service({
        url: "/resource/server/serverConnect",
        method: 'post',
        data
    })
}