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

