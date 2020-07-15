package service

import (
	"fmt"
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

// @title    TestingList
// @description   get test order list by pagination, 分页获取数据
// @auth                      （2020/04/05  20:22）
// @param     info             request.PageInfo
// @return    err              error
// @return    list             interface{}
// @return    total            int

func TestingList(info request.PageInfo) (err error, list interface{}, total int) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&model.DeployTesting{})
	var testingList []model.DeployTesting
	err = db.Count(&total).Error
	err = db.Preload("DeployProject").Limit(limit).Offset(offset).Find(&testingList).Error
	return err, testingList, total
}

// @title    TestingContrast
// @description   对比文件
// @auth                      （2020/07/15  09:50）
// @param     info             request.ContrastInfo
// @return    err              error
// @return    list             interface{}
// @return    total            int

func TestingContrast(info request.ContrastInfo) (err error, list interface{}) {
	fmt.Println(info)
	return err, list
}
