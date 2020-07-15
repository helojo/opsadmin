import service from '@/utils/request'
// @Tags Deploy_Testing
// @Summary 分页获取提测列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body modelInterface.PageInfo true "分页获取提测列表"
// @Success 200 {string} json "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /deploy/test/testingList [post]
// {
//  page     int
//	pageSize int
// }
export const testingList = (data) => {
    return service({
        url: "/deploy/test/testingList",
        method: 'post',
        data
    })
}

// @Tags Deploy_Testing
// @Summary 文件对比
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body modelInterface.PageInfo true "文件对比"
// @Success 200 {string} json "{"success":true,"data":{},"msg":"文件对比成功"}"
// @Router /deploy/test/testingContrast [post]
// {
//  page     int
//	pageSize int
// }
export const testingContrast = (data) => {
    return service({
        url: "/deploy/test/testingContrast",
        method: 'post',
        data
    })
}