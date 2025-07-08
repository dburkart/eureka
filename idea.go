package eureka

import "time"

type Idea struct {
	ID                string         `json:"id"`
	CreatedAt         time.Time      `json:"created_at"`
	UpdatedAt         time.Time      `json:"updated_at"`
	ReferenceNumber   string         `json:"reference_num"`
	Name              string         `json:"name"`
	WorkflowStatus    WorkflowStatus `json:"workflow_status"`
	URL               string         `json:"url"`
	Resource          string         `json:"resource"`
	Score             int            `json:"score"`
	ProductID         string         `json:"product_id"`
	Votes             int            `json:"votes"`
	InitialVotes      int            `json:"initial_votes"`
	Visibility        string         `json:"visibility"`
	Product           Product        `json:"product"`
	EndorsementsCount int            `json:"endorsements_count"`
	CommentsCount     int            `json:"comments_count"`
}

type IdeaListResponse struct {
	PaginatedResponse
	Ideas []Idea `json:"ideas"`
}
