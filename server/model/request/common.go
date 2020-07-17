package request

// Paging common input parameter structure
type PageInfo struct {
	Page     int `json:"page" form:"page"`
	PageSize int `json:"pageSize" form:"pageSize"`
}

// Find by id structure
type GetById struct {
	Id float64 `json:"id" form:"id"`
}

// Paging common input parameter structure
type ServerPageInfo struct {
	Id            float64 `json:"id" form:"id"`
	Page          int     `json:"page" form:"page"`
	PageSize      int     `json:"pageSize" form:"pageSize"`
	ResourceEnvId int     `json:"resource_env_id"  form:"resource_env_id"`
}

type ProjectPageInfo struct {
	Id            float64 `json:"id" form:"id"`
	Page          int     `json:"page" form:"page"`
	PageSize      int     `json:"pageSize" form:"pageSize"`
	ResourceEnvId int     `json:"resource_env_id"  form:"resource_env_id"`
}

type ContrastInfo struct {
	Tag             string `json:"tag" form:"tag"`
	ResourceEnvId   int    `json:"resource_env_id" form:"resource_env_id"`
	DeployProjectId int    `json:"deploy_project_id" form:"deploy_project_id"`
}

type TestingReleaseInfo struct {
	Tag             string   `json:"tag" form:"tag"`
	Files           []string `json:"files" form:"files"`
	ResourceEnvId   int      `json:"resource_env_id" form:"resource_env_id"`
	DeployProjectId int      `json:"deploy_project_id" form:"deploy_project_id"`
}
