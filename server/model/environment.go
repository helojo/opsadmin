package model

type Environment struct {
	ID       int    `json:"id" gorm:"not null;primary_key"`
	Name     string `json:"name"`      //环境名称
	EnvLabel int    `json:"env_label"` // 0: 无标签 1: 开发 2:测试, 3:灰度 4: 生产

}
