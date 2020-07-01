package model

type DepolyServer struct {
	ID   int    `json:"id" gorm:"not null;primary_key"`
	Name string `json:"name"` //环境名称
	Host string `json:"host"` //主机地址
	Port int64  `json:"port"` //端口
	User string `json:"user"` //用户名
	Pwd  string `json:"pwd"`  //密码
}
