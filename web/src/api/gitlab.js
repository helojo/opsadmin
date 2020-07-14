import service from '@/utils/request'
// @Tags Gitlab_Project
// @Summary gitlab 项目导入
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body modelInterface.PageInfo true "项目导入"
// @Success 200 {string} json "{"success":true,"data":{},"msg":"项目导入成功"}"
// @Router /gitlab/project/projectImport [get]
export const projectImport = () => {
    return service({
        url: "/gitlab/project/projectImport",
        method: 'get',
    })
}

// @Tags Gitlab_Project
// @Summary 获取项目tag
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body modelInterface.PageInfo true "获取项目tag"
// @Success 200 {string} json "{"success":true,"data":{},"msg":"获取项目tag成功"}"
// @Router /gitlab/project/projectTags [post]
export const projectTags = (data) => {
    return service({
        url: "/gitlab/project/projectTags",
        method: 'post',
        data
    })
}
