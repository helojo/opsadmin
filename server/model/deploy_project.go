package model

type DeployProject struct {
	ID             int         `json:"id" gorm:"not null;primary_key"`
	Name           string      `json:"name"`                              //项目名称
	GitUrl         string      `gorm:"type:longText" json:"git_url"`      //git项目地址
	Directory      string      `json:"directory"`                         //项目目录绝对路径
	IgnoreFiles    string      `gorm:"type:longText" json:"ignore_files"` //忽略文件
	ReleaseVersion float64     `json:"release_version"`
	Server         []Server    `gorm:"many2many:deployproject_servers;"`
	EnvironmentId  int         `json:"environment_id"`
	Environment    Environment `json:"environment" gorm:"ForeignKey:EnvironmentId"`
}
