package model

import "github.com/jinzhu/gorm"

type DeployRollback struct {
	gorm.Model
	ReleaseVersion  float64       `json:"release_version"`               //回退前版本
	AfterVersion    float64       `json:"after_version"`                 //回退后版本
	Aperator        string        `json:"aperator"`                      // 操作人
	Describe        string        `gorm:"type:longText" json:"describe"` // 描述
	Result          string        `gorm:"type:longText" json:"result"`   // 结果
	Status          int           `json:"status"`                        //# 1:回退中 2:回退成功 3:回退失败
	DeployProjectId int           `json:"deploy_project_id"`             //项目
	DeployProject   DeployProject `json:"deployproject" gorm:"ForeignKey:DeployProjectId"`
}
