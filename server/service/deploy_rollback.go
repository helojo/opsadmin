package service

import (
	"errors"
	"fmt"
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
	"gin-vue-admin/utils"
	"strings"
)

// @title    RollbackList
// @description   get rollback order list by pagination, 分页获取数据
// @auth                      （2020/04/05  20:22）
// @param     info             request.PageInfo
// @return    err              error
// @return    list             interface{}
// @return    total            int

func RollbackList(info request.PageInfo) (err error, list interface{}, total int) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&model.DeployRollback{})
	var rollbackList []model.DeployRollback
	err = db.Count(&total).Error
	err = db.Order("id DESC ").Preload("DeployProject.Environment").Preload("DeployProject.Server").Limit(limit).Offset(offset).Find(&rollbackList).Error
	return err, rollbackList, total
}

// @title    RollbackContrast
// @description   回滚对比文件
// @auth                      （2020/07/15  09:50）
// @param     info             request.ContrastInfo
// @return    err              error
// @return    list             interface{}
// @return    path             string

func RollbackContrast(rollback request.RollbackContrast) (err error, list interface{}, path string) {
	var project model.DeployProject
	err = global.GVA_DB.Where("id = ?", rollback.DeployProjectId).Preload("Server").First(&project).Error
	if err != nil {
		return errors.New(fmt.Sprint("查询项目报错, 报错信息: %s", err)), list, path
	}

	var rollbackOrder model.DeployTesting
	err = global.GVA_DB.Where("id = ?", rollback.Version).First(&rollbackOrder).Error

	if err != nil {
		return errors.New(fmt.Sprint("查询回滚工单项目报错, 报错信息: %s", err)), list, path
	}

	exclude := strings.Fields(project.IgnoreFiles)

	var result []map[string]string
	var errMsg string
	for _, value := range project.Server {
		err, ret := utils.FileContrast(rollbackOrder.Path, value.User, value.Host, value.Port, project.Directory, exclude)
		if err != nil {
			errMsg += fmt.Sprint("对比文件报错, 报错信息: %s", err)
		}
		result = append(result, ret...)
	}

	return err, result, path
}

// @title    RollbackRelease
// @description   提交并同步要发布的文件
// @auth                      （2020/07/17  19:45）
// @param     info             request.TestingReleaseInfo
// @return    err              error

func RollbackRelease(rollback request.RollbackContrast, username *request.CustomClaims) (err error) {
	var project model.DeployProject
	err = global.GVA_DB.Where("id = ?", rollback.DeployProjectId).Preload("Server").First(&project).Error
	if err != nil {
		return errors.New(fmt.Sprint("查询项目报错, 报错信息: %s", err))
	}
	var rollbackOrder model.DeployTesting
	err = global.GVA_DB.Where("id = ?", rollback.Version).First(&rollbackOrder).Error

	if err != nil {
		return errors.New(fmt.Sprint("查询回滚工单项目报错, 报错信息: %s", err))
	}

	go func() {
		deployrollback := &model.DeployRollback{
			ReleaseVersion:  project.ReleaseVersion,
			AfterVersion:    rollbackOrder.Version,
			Aperator:        username.NickName,
			Describe:        rollback.Describe,
			DeployProjectId: rollback.DeployProjectId,
			Status:          1,
		}
		err = global.GVA_DB.Create(deployrollback).Error

		result := ""
		exclude := strings.Fields(project.IgnoreFiles)
		for _, value := range project.Server {
			result += fmt.Sprintf("==========================主机开始同步文件: %s=========================\n", value.Host)
			err, ret := utils.FileSync(rollbackOrder.Path, value.User, value.Host, value.Port, project.Directory, exclude)
			if err != nil {
				result += fmt.Sprintf("同步报错: %s", err)
			}
			result += ret

		}

		if strings.HasPrefix(result, "同步报错") {
			err = RollbackUpdate(deployrollback.ID, 2, result)
		} else {
			err = RollbackUpdate(deployrollback.ID, 1, result)
			project.ReleaseVersion = rollbackOrder.Version
			err = ProjectStatusUpdate(project)
		}
	}()

	return err
}

// @title    RollbackUpdate
// @description    更新回滚状态
// @auth                     （2020/04/05  20:22）
// @param     env             model.DeployRollback
// @return                    error

func RollbackUpdate(id uint, status int, result string) (err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&model.DeployRollback{}).Updates(&model.DeployRollback{Status: status, Result: result}).Error
	return err
}
