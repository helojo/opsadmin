package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/utils"
)

// @title    ProjectImport
// @description    Gitlab  项目导入
// @auth                     （2020/07/13  18:58）
// @param     id
// @return    err             error

func ProjectImport() (err error) {
	projecs, err := utils.GetGitlabprojects()
	if err == nil {
		go func() {
			for id, project := range projecs {
				gitproject := &model.GitlabProject{
					ID:   id,
					Name: project["name"],
					Url:  project["url"],
				}
				notRegister := global.GVA_DB.Where("id = ?", id).First(&gitproject).RecordNotFound()
				if notRegister {
					err = global.GVA_DB.Create(&gitproject).Error
				}
			}
		}()
	}
	return err
}

// @title    ProjectTags
// @description    获取项目tag
// @auth                     （2020/07/14  10:06）
// @param     id
// @return    err             error
func ProjectTags(p model.DeployProject) (taglist []interface{}, err error) {
	var project model.DeployProject
	var gitlab model.GitlabProject
	err = global.GVA_DB.Where("id = ?", p.ID).First(&project).Error
	if err == nil {
		err = global.GVA_DB.Where("url = ?", project.GitUrl).First(&gitlab).Error
		if err == nil {
			tagList, err := utils.GetProjectTags(gitlab.ID)
			if err == nil {
				return tagList, err
			}
		}
	}
	return taglist, err
}
