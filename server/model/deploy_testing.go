package model

import "github.com/jinzhu/gorm"

type DeployTesting struct {
	gorm.Model
	Applicant       string        `json:"applicant"`                     // 申请人
	Version         float64       `json:"version"`                       // 提测版本
	Tag             string        `json:"tag"`                           // tag
	Describe        string        `gorm:"type:longText" json:"describe"` // 描述
	Path            string        `gorm:"type:longText" json:"path"`     // 存储目录
	Status          int           `json:"status"`                        //1:提测成功 2:提测失败
	Result          string        `gorm:"type:longText" json:"result"`   // 发布结果
	Isdelete        int           `json:"isdelete"`                      //1:备份未删除 2:备份已删除
	DeployProjectId int           `json:"deploy_project_id"`
	DeployProject   DeployProject `json:"deployproject" gorm:"ForeignKey:DeployProjectId"`
}
