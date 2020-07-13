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
