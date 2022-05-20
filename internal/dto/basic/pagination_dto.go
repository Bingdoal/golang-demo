package basic

type Pagination struct {
	Page     uint64 `json:"page"`
	PageSize uint64 `json:"pageSize"`
	Total    int64  `json:"total"`
}
