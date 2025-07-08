package eureka

type Pagination struct {
	TotalRecords int `json:"total_records"`
	TotalPages   int `json:"total_pages"`
	CurrentPage  int `json:"current_page"`
}

type APIResponse struct {
	Pagination Pagination `json:"pagination"`
}
