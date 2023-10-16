package common

// PageParam 分页参数
type PageParam struct {
	Current int `json:"current" dc:"分页-当前页码（从1开始）"`
	Size    int `json:"size" dc:"分页-每页记录数（默认为10）"`
}

// PageResult 分页结果
type PageResult struct {
	Current int `json:"current" dc:"分页-当前页码（从1开始）"`
	Size    int `json:"size" dc:"分页-每页记录数（默认为10）"`
	Pages   int `json:"pages" dc:"分页-页数"`
	Total   int `json:"total" dc:"分页-总记录"`
}
