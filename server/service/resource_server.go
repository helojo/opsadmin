package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

// @title    GetServerList
// @description   get server list by pagination, 分页获取数据
// @auth                      （2020/04/05  20:22）
// @param     info             request.PageInfo
// @return    err              error
// @return    list             interface{}
// @return    total            int

func GetServerList(info request.PageInfo) (err error, list interface{}, total int) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&model.ResourceServer{})
	var envList []model.ResourceServer
	err = db.Count(&total).Error
	err = db.Preload("ResourceEnv").Limit(limit).Offset(offset).Find(&envList).Error
	return err, envList, total
}
