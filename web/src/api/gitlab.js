import service from '@/utils/request'
// @Tags Gitlab_Project
// @Summary gitlab 项目导入
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body modelInterface.PageInfo true "项目导入"
// @Success 200 {string} json "{"success":true,"data":{},"msg":"项目导入成功"}"
// @Router /gitlab/project/projectImport [get]
// {
//  page     int
//	pageSize int
// }
export const projectImport = () => {
    return service({
        url: "/gitlab/project/projectImport",
        method: 'get',
    })
}
