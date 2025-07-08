package eureka

import "time"

// Aha! Product
type Product struct {
	ID              string    `json:"id"`
	CreatedAt       time.Time `json:"created_at"`
	Name            string    `json:"name"`
	ReferencePrefix string    `json:"reference_prefix"`
	ProductLine     bool      `json:"product_line"`
	WorkspaceType   string    `json:"workspace_type"`
	URL             string    `json:"url"`
}

type ProductListResponse struct {
	PaginatedResponse
	Products []Product `json:"products"`
}
