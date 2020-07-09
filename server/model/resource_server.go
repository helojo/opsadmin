package model

type ResourceServer struct {
	ID            int         `json:"id" gorm:"not null;primary_key"`
	Name          string      `json:"name"`   //主机名称
	Host          string      `json:"host"`   //主机地址
	Port          int64       `json:"port"`   //SSH端口
	User          string      `json:"user"`   //用户名
	Pwd           string      `json:"pwd"`    //密码
	Status        int         `json:"status"` //测试连接，1.未测试，2.连接中, 3.连接成功, 4. 连接异常, 5.推送公钥成功，6。推送公钥失败
	ResourceEnvId int         `json:"resource_env_id"`
	ResourceEnv   ResourceEnv `json:"resourceenv" gorm:"ForeignKey:ResourceEnvId"`
}
