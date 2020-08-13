package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
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
