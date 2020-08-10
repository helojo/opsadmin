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
// @return    path             string

func TestingContrast(testting request.ContrastInfo) (err error, list interface{}, path string) {
	var project model.DeployProject
	err = global.GVA_DB.Where("id = ?", testting.DeployProjectId).Preload("ResourceServer").First(&project).Error
	if err != nil {
		return errors.New(fmt.Sprint("查询项目报错, 报错信息: %s", err)), list, path
	}

	path, err = utils.Gitpull(testting.Tag, project.GitUrl)
	if err != nil {
		return errors.New(fmt.Sprint("Git拉取项目报错, 报错信息: %s", err)), list, path
	}

	exclude := strings.Fields(project.IgnoreFiles)
	err, list = utils.FileContrast(path, project.ResourceServer.User, project.ResourceServer.Host, project.Directory, exclude)
	if err != nil {
		return errors.New(fmt.Sprint("对比文件报错, 报错信息: %s", err)), list, path
	}

	return err, list, path
}

// @title    TestingRelease
// @description   提交并同步要发布的文件
// @auth                      （2020/07/17  19:45）
// @param     info             request.TestingReleaseInfo
// @return    err              error

func TestingRelease(testting request.TestingReleaseInfo, username *request.CustomClaims) (err error) {
	var project model.DeployProject
	err = global.GVA_DB.Where("id = ?", testting.DeployProjectId).Preload("ResourceServer").First(&project).Error
	if err != nil {
		return errors.New(fmt.Sprint("查询项目报错, 报错信息: %s", err))
	}

	exclude := strings.Fields(project.IgnoreFiles)
	err, result := utils.FileSync(testting.Path, project.ResourceServer.User, project.ResourceServer.Host, project.Directory, exclude)

	testOrder := &model.DeployTesting{
		Applicant:       username.NickName,
		Tag:             testting.Tag,
		Result:          result,
		DeployProjectId: testting.DeployProjectId,
		Describe:        testting.Describe,
	}

	if err != nil {
		testOrder.Status = 2
	}

	testOrder.Status = 1
	err = global.GVA_DB.Create(testOrder).Error
	if err != nil {
		return errors.New(fmt.Sprint("提测工单创建失败: %s", err))
	}

	return err
}
