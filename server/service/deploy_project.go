package service

import (
	"errors"
	"fmt"
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
	"os"
	"strconv"
)

// @title    PtList
// @description   get PtList list by pagination, 获取筛选排序条件数据
// @auth                      （2020/07/13  09:48）
// @return    err              error
// @return    list             interface{}
// @return    total            int
func PtList(ordercondition string, wherecondition string, limit, offset int) (err error, list interface{}, total int) {
	db := global.GVA_DB
	var projecList []model.DeployProject
	err = db.Where(wherecondition).Find(&projecList).Count(&total).Error
	err = db.Order(ordercondition).Where(wherecondition).
		Preload("Server").Preload("Environment").
		Limit(limit).Offset(offset).Find(&projecList).Error
	return err, projecList, total
}

// @title    ProjectList
// @description   get project list by pagination, 分页获取数据
// @auth                      （2020/04/05  20:22）
// @param     info             request.PageInfo
// @return    err              error
// @return    list             interface{}
// @return    total            int

func ProjectList(info request.ProjectPageInfo) (err error, list interface{}, total int) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	if info.EnvironmentId != 0 {
		wherecondition := fmt.Sprintf("environment_id = %d", info.EnvironmentId)
		return PtList("", wherecondition, limit, offset)

	}
	return PtList("", "", limit, offset)
}

// @title    ProjectCreate
// @description   create base project, 新增项目
// @auth                     （2020/07/07  17:56）
// @param     api             model.Server
// @return                    error

func ProjectCreate(p request.DeployProject) (err error) {
	findOne := global.GVA_DB.Where("environment_id = ? and name = ?", p.EnvironmentId, p.Name).Find(&model.DeployProject{}).Error
	if findOne == nil {
		return errors.New("存在相同项目")
	} else {
		var gitproject model.GitlabProject
		notRegister := global.GVA_DB.Where("url = ?", p.GitUrl).First(&gitproject).RecordNotFound()
		if notRegister {
			return errors.New("项目地址未找到，请核对Git地址!")
		}

		var servers []model.Server
		err = global.GVA_DB.Where("id in (?)", p.Server).Find(&servers).Error
		fmt.Println(servers)
		if err == nil {
			p.ReleaseVersion = 0.1
			reservedversion, _ := strconv.Atoi(p.Reservedversion)
			project := model.DeployProject{
				Name:            p.Name,
				GitUrl:          p.GitUrl,
				Directory:       p.Directory,
				IgnoreFiles:     p.IgnoreFiles,
				ReleaseVersion:  0.1,
				Reservedversion: reservedversion,
				EnvironmentId:   p.EnvironmentId,
			}
			err = global.GVA_DB.Create(&project).Error
			if err == nil {
				err = global.GVA_DB.Model(&project).Association("Server").Replace(&servers).Error
			}
		}

	}
	return err
}

// @title    ProjectUpdate
// @description   更新项目
// @auth                     （2020/07/10  15:22）
// @param     project         model.DeployProject
// @return                    error

func ProjectUpdate(project request.DeployProject) (err error) {
	var gitproject model.GitlabProject
	notRegister := global.GVA_DB.Where("url = ?", project.GitUrl).First(&gitproject).RecordNotFound()
	if notRegister {
		return errors.New("项目地址未找到，请核对Git地址!")
	}

	err = global.GVA_DB.Where("id = ?", project.ID).First(&model.DeployProject{}).Updates(&project).Error
	if err == nil {
		var servers []model.Server
		err = global.GVA_DB.Where("id in (?)", project.Server).Find(&servers).Error
		if err == nil {
			err = global.GVA_DB.Model(&model.DeployProject{ID: project.ID}).Association("Server").Replace(&servers).Error
		}
	}
	return err
}

// @title    ProjectStatusUpdate
// @description   更新项目状态
// @auth                     （2020/07/10  15:22）
// @param     project         model.DeployProject
// @return                    error

func ProjectStatusUpdate(project model.DeployProject) (err error) {
	err = global.GVA_DB.Where("id = ?", project.ID).First(&model.DeployProject{}).Updates(&project).Error
	return err
}

// @title    ProjectDelete
// @description    删除项目
// @auth                     （2020/07/10  15:11）
// @param     id
// @return    err             error

func ProjectDelete(id float64) (err error) {
	var project model.DeployProject
	err = global.GVA_DB.Where("id = ?", id).Delete(&project).Error
	return err
}

// @title    TestingversionDelete
// @description    删除多余的备份版本
// @auth                     （2020/07/10  15:11）
// @param     id
// @return    err             error

func TestingVersionDelete(id float64) (err error) {
	var project model.DeployProject
	err = global.GVA_DB.Where("id = ?", id).Find(&project).Error
	if err == nil {
		deleteVersion := project.ReleaseVersion - float64(project.Reservedversion)*0.1

		var testting []model.DeployTesting
		err = global.GVA_DB.Where("deploy_project_id = ? and isdelete = 1", project.ID).Find(&testting).Error
		fmt.Println(testting)
		if err == nil {
			for _, testOrder := range testting {
				if testOrder.Version <= deleteVersion {
					err = os.RemoveAll(testOrder.Path)
					if err == nil {
						var testorder model.DeployTesting
						testorder.Isdelete = 2
						err = global.GVA_DB.Where("id = ?", testOrder.ID).First(&model.DeployTesting{}).Updates(&testorder).Error
					}
				}
			}

		}
	}

	return err
}

// @title    OnlineVersionDelete
// @description    删除多余的备份版本
// @auth                     （2020/07/10  15:11）
// @param     id
// @return    err             error

func OnlineVersionDelete(id float64) (err error) {
	var project model.DeployProject
	err = global.GVA_DB.Where("id = ?", id).Find(&project).Error
	if err == nil {
		deleteVersion := project.ReleaseVersion - float64(project.Reservedversion)*0.1

		var onlineting []model.DeployOnline
		err = global.GVA_DB.Where("deploy_project_id = ? and isdelete = 1", project.ID).Find(&onlineting).Error
		if err == nil {
			for _, OnOrder := range onlineting {
				if OnOrder.Version <= deleteVersion {
					err = os.RemoveAll(OnOrder.Path)
					if err == nil {
						var onlineorder model.DeployOnline
						onlineorder.Isdelete = 2
						err = global.GVA_DB.Where("id = ?", OnOrder.ID).First(&model.DeployOnline{}).Updates(&onlineorder).Error
					}
				}
			}

		}
	}

	return err
}
