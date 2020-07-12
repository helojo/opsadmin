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

// @Tags Deploy_Project
// @Summary 创建项目
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body api.projectCreate true "创建项目"
// @Success 200 {string} json "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /deploy/project/projectCreate [post]
export const projectCreate = (data) => {
    return service({
        url: "/deploy/project/projectCreate",
        method: 'post',
        data
    })
}

// @Tags Deploy_Project
// @Summary 更新项目
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body true "更新项目"
// @Success 200 {string} json "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /deploy/project/projectUpdate [post]
export const projectUpdate = (data) => {
    return service({
        url: "/deploy/project/projectUpdate",
        method: 'post',
        data
    })
}

// @Tags Deploy_Project
// @Summary 删除项目
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body true "删除项目"
// @Success 200 {string} json "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /deploy/project/projectDelete [delete]
export const projectDelete = (data) => {
    return service({
        url: "/deploy/project/projectDelete",
        method: 'delete',
        data
    })
}