package service

import (
	"errors"
	"fmt"
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
	"gin-vue-admin/utils"
	"strconv"
	"strings"
)

// @title    OnlineList
// @description   get test order list by pagination, 分页获取数据
// @auth                      （2020/04/05  20:22）
// @param     info             request.PageInfo
// @return    err              error
// @return    list             interface{}
// @return    total            int

func OnlineList(info request.PageInfo) (err error, list interface{}, total int) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&model.DeployOnline{})
	var onlineList []model.DeployOnline
	err = db.Count(&total).Error
	err = db.Order("id DESC ").Preload("DeployProject.Environment").Preload("DeployProject.Server").Limit(limit).Offset(offset).Find(&onlineList).Error
	return err, onlineList, total
}

// @title    OnlineContrast
// @description   对比文件
// @auth                      （2020/07/15  09:50）
// @param     info             request.ContrastInfo
// @return    err              error
// @return    list             interface{}
// @return    path             string

func OnlineContrast(online request.ContrastInfo) (err error, list interface{}, path string) {
	var project model.DeployProject
	err = global.GVA_DB.Where("id = ?", online.DeployProjectId).Preload("Server").First(&project).Error
	if err != nil {
		return errors.New(fmt.Sprint("查询项目报错, 报错信息: %s", err)), list, path
	}

	var onlineOrder model.DeployOnline
	nofound := global.GVA_DB.Where("deploy_project_id = ? and status < 3", online.DeployProjectId).First(&onlineOrder).RecordNotFound()
	if !nofound {
		return errors.New(fmt.Sprint("该项目有未完成工单！")), list, path
	}

	filepath := fmt.Sprintf("%d-%s-%.1f", project.ID, project.Name, project.ReleaseVersion+0.1)
	path, err = utils.Gitpull(online.Tag, project.GitUrl, filepath)
	if err != nil {
		return errors.New(fmt.Sprint("Git拉取项目报错, 报错信息: %s", err)), list, path
	}

	exclude := strings.Fields(project.IgnoreFiles)

	var result []map[string]string
	var errMsg string
	for _, value := range project.Server {
		err, ret := utils.FileContrast(path, value.User, value.Host, value.Port, project.Directory, exclude)
		if err != nil {
			errMsg += fmt.Sprint("对比文件报错, 报错信息: %s", err)
		}
		result = append(result, ret...)
	}

	return err, result, path
}

// @title    OnlineCreate
// @description   创建上线工单
// @auth                      （2020/07/17  19:45）
// @param     info             request.TestingReleaseInfo
// @return    err              error

func OnlineCreate(online request.OnlineInfo, username *request.CustomClaims) (err error) {

	var onlineOrder model.DeployOnline
	nofound := global.GVA_DB.Where("deploy_project_id = ? and status < 3", online.DeployProjectId).First(&onlineOrder).RecordNotFound()
	if !nofound {
		return errors.New(fmt.Sprint("该项目有未完成工单！"))
	}

	var project model.DeployProject
	err = global.GVA_DB.Where("id = ? ", online.DeployProjectId).First(&project).Error
	if err != nil {
		return errors.New(fmt.Sprint("查询项目报错, 报错信息: %s", err))
	}

	err = global.GVA_DB.Where("id = ?", online.DeployProjectId).Preload("Server").First(&project).Error
	if err != nil {
		return errors.New(fmt.Sprint("查询项目报错, 报错信息: %s", err))
	}
	version, _ := strconv.ParseFloat(fmt.Sprintf("%.1f", project.ReleaseVersion+0.1), 64)
	testOrder := &model.DeployOnline{
		Applicant:       username.NickName,
		Tag:             online.Tag,
		DeployProjectId: online.DeployProjectId,
		Describe:        online.Describe,
		Path:            online.Path,
		Version:         version,
		Isdelete:        1,
	}

	err = global.GVA_DB.Create(testOrder).Error
	return err
}

// @title    OnlineUpdate
// @description    更新上线工单
// @auth                     （2020/04/05  20:22）
// @param     env             model.DeployOnline
// @return                    error

func OnlineUpdate(id uint, status int, username string, result string) (err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&model.DeployOnline{}).Updates(&model.DeployOnline{Status: status, OpAuditor: username, Result: result}).Error
	return err
}

// @title   OnlineRversion
// @description    可回滚版本
// @auth                     （2020/07/10  15:11）
// @param     id
// @return    err             error

func OnlineRversion(id float64) (err error, list interface{}) {
	var online []model.DeployOnline
	err = global.GVA_DB.Where("deploy_project_id = ? and isdelete = 1", id).Group("version").Order("created_at desc").Find(&online).Error
	return err, online
}

// @title   DevAudit
// @description    开发审核
// @auth                     （2020/07/10  15:11）
// @param     id
// @return    err             error

func DevAudit(id float64, username string) (err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&model.DeployOnline{}).Updates(&model.DeployOnline{Status: 1, DevAuditor: username}).Error
	return err
}

// @title   TestAudit
// @description    测试审核
// @auth                     （2020/07/10  15:11）
// @param     id
// @return    err             error

func TestAudit(id float64, username string) (err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&model.DeployOnline{}).Updates(&model.DeployOnline{Status: 2, TestAuditor: username}).Error
	return err
}

// @title  OpsAudit
// @description    运维审核
// @auth                     （2020/07/10  15:11）
// @param     id
// @return    err             error

func OpsAudit(id float64, username string) (err error) {
	//err = global.GVA_DB.Where("id = ?", id).First(&model.DeployOnline{}).Updates(&model.DeployOnline{Status: 3, OpAuditor: username}).Error

	var onlineOrder model.DeployOnline
	err = global.GVA_DB.Where("id = ?", id).First(&onlineOrder).Error
	if err != nil {
		return errors.New(fmt.Sprint("查询上线工单报错, 报错信息: %s", err))
	}

	var project model.DeployProject
	err = global.GVA_DB.Where("id = ?", onlineOrder.DeployProjectId).Preload("Server").First(&project).Error
	if err != nil {
		return errors.New(fmt.Sprint("查询项目报错, 报错信息: %s", err))
	}

	err = OnlineUpdate(onlineOrder.ID, 4, username, "文件同步中")
	go func() {
		result := ""
		exclude := strings.Fields(project.IgnoreFiles)
		for _, value := range project.Server {
			result += fmt.Sprintf("==========================主机开始同步文件: %s=========================\n", value.Host)
			err, ret := utils.FileSync(onlineOrder.Path, value.User, value.Host, value.Port, project.Directory, exclude)
			if err != nil {
				result += fmt.Sprintf("同步报错: %s", err)
			}
			result += ret
		}

		if strings.HasPrefix(result, "同步报错") {
			err = OnlineUpdate(onlineOrder.ID, 2, username, result)
		} else {
			err = OnlineUpdate(onlineOrder.ID, 3, username, result)
			project.ReleaseVersion = onlineOrder.Version
			err = ProjectStatusUpdate(project)
		}

		// 删除过期备份
		_ = OnlineVersionDelete(float64(project.ID))
	}()

	return err
}
