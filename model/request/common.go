package request

// Paging common input parameter structure
type PageInfo struct {
	Page     int `json:"page" form:"page"`
	PageSize int `json:"page_size" form:"page_size"`
}

// Find by id structure
type GetById struct {
	Id float64 `json:"id" form:"id"`
}
