package model

type DeployTesting struct {
	ID              int           `json:"id" gorm:"not null;primary_key"`
	Applicant       string        `json:"applicant"`                     // 申请人
	DevAuditor      string        `json:"dev_auditor"`                   // 开发审核人
	TestAuditor     string        `json:"test_auditor"`                  // 测试审核人
	OpAuditor       string        `json:"op_auditor"`                    // 运维审核人
	Version         float64       `json:"version"`                       // 提测版本
	Tag             string        `json:"tag"`                           // tag
	Files           string        `gorm:"type:longText" json:"files"`    // 要同步的文件
	Describe        string        `gorm:"type:longText" json:"describe"` // 描述
	Status          int           `json:"status"`                        //1:开发审核 2:测试审核 3:运维审核/已上线 4:审核中 5:已关闭   6;正在关闭
	DeployProjectId int           `json:"deploy_project_id"`
	DeployProject   DeployProject `json:"deployproject" gorm:"ForeignKey:DeployProjectId"`
}
