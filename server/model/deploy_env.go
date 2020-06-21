package model

type DepolyEnv struct {
	ID   int    `json:"id" gorm:"not null;primary_key"`
	Name string `json:"name"` //环境名称
}
