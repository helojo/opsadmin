package service

import (
	"errors"
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

// @title    ProjectList
// @description   get project list by pagination, 分页获取数据
// @auth                      （2020/04/05  20:22）
// @param     info             request.PageInfo
// @return    err              error
// @return    list             interface{}
// @return    total            int

func ProjectList(info request.PageInfo) (err error, list interface{}, total int) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&model.DeployProject{})
	var projectList []model.DeployProject
	err = db.Count(&total).Error
	err = db.Preload("ResourceServer.ResourceEnv").Limit(limit).Offset(offset).Find(&projectList).Error
	return err, projectList, total
}

// @title    ProjectCreate
// @description   create base project, 新增项目
// @auth                     （2020/07/07  17:56）
// @param     api             model.ResourceServer
// @return                    error

func ProjectCreate(project model.DeployProject) (err error) {
	findOne := global.GVA_DB.Where("name = ?", project.Name).Find(&model.DeployProject{}).Error
	if findOne == nil {
		return errors.New("存在相同项目")
	} else {
		project.ReleaseVersion = 0.1
		err = global.GVA_DB.Create(&project).Error
	}
	return err
}
