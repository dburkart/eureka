package eureka

type WorkflowStatus struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Position int    `json:"position"`
	Complete bool   `json:"complete"`
	Color    string `json:"color"`
}
