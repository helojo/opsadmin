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
type EnvironmentPageInfo struct {
	Page     int `json:"page" form:"page"`
	Status   int `json:"status" form:"status"`
	PageSize int `json:"pageSize" form:"pageSize"`
}

// Paging common input parameter structure
type ServerPageInfo struct {
	Id            float64 `json:"id" form:"id"`
	Page          int     `json:"page" form:"page"`
	PageSize      int     `json:"pageSize" form:"pageSize"`
	EnvironmentId int     `json:"environment_id"  form:"environment_id"`
}

type ProjectPageInfo struct {
	Id            float64 `json:"id" form:"id"`
	Page          int     `json:"page" form:"page"`
	PageSize      int     `json:"pageSize" form:"pageSize"`
	EnvironmentId int     `json:"environment_id"  form:"environment_id"`
}

type ContrastInfo struct {
	Tag             string `json:"tag" form:"tag"`
	EnvironmentId   int    `json:"environment_id" form:"environment_id"`
	DeployProjectId int    `json:"deploy_project_id" form:"deploy_project_id"`
}

type RollbackContrast struct {
	Version         int    `json:"version" form:"version"`
	Describe        string `json:"describe" form:"describe"`
	EnvironmentId   int    `json:"environment_id" form:"environment_id"`
	DeployProjectId int    `json:"deploy_project_id" form:"deploy_project_id"`
}

type TestingReleaseInfo struct {
	Tag             string   `json:"tag" form:"tag"`
	Path            string   `json:"path" form:"path"`
	Files           []string `json:"files" form:"files"`
	Describe        string   `json:"describe" form:"describe"`
	EnvironmentId   int      `json:"environment_id" form:"environment_id"`
	DeployProjectId int      `json:"deploy_project_id" form:"deploy_project_id"`
}

type OnlineInfo struct {
	Tag             string   `json:"tag" form:"tag"`
	Path            string   `json:"path" form:"path"`
	Files           []string `json:"files" form:"files"`
	Describe        string   `json:"describe" form:"describe"`
	EnvironmentId   int      `json:"environment_id" form:"environment_id"`
	DeployProjectId int      `json:"deploy_project_id" form:"deploy_project_id"`
}
