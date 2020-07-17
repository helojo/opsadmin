package response

type PageResult struct {
	List     interface{} `json:"list"`
	Total    int         `json:"total"`
	Page     int         `json:"page"`
	PageSize int         `json:"pageSize"`
}

type ContrastResult struct {
	List interface{} `json:"list"`
	Path string      `json:"path"`
}
