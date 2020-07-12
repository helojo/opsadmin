package model

import "github.com/jinzhu/gorm"

type DeployTesting struct {
	gorm.Model
	Applicant       string        `json:"applicant"`                     // 申请人
	Version         float64       `json:"version"`                       // 提测版本
	Tag             string        `json:"tag"`                           // tag
	Files           string        `gorm:"type:longText" json:"files"`    // 同步的文件
	Describe        string        `gorm:"type:longText" json:"describe"` // 描述
	Status          int           `json:"status"`                        //0, 提测中, 1:提测成功 2:提测失败
	DeployProjectId int           `json:"deploy_project_id"`
	DeployProject   DeployProject `json:"deployproject" gorm:"ForeignKey:DeployProjectId"`
}
