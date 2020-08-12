package request

type DeployProject struct {
	ID              int     `json:"id"`
	Name            string  `json:"name"`         //项目名称
	GitUrl          string  `json:"git_url"`      //git项目地址
	Directory       string  `json:"directory"`    //项目目录绝对路径
	IgnoreFiles     string  `json:"ignore_files"` //忽略文件
	ReleaseVersion  float64 `json:"release_version"`
	Reservedversion string  `json:"reservedversion"` //保留版本数
	Server          []int   `json:"server"`
	EnvironmentId   int     `json:"environment_id"`
}
