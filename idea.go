package eureka

import "time"

type Idea struct {
	RefObject
	AssignedToUser               any                `json:"assigned_to_user,omitempty"`
	Categories                   []Category         `json:"categories,omitempty"`
	CommentsCount                int                `json:"comments_count,omitempty"`
	CreatedByUser                CreatedByUser      `json:"created_by_user,omitempty"`
	CustomFields                 []CustomField      `json:"custom_fields,omitempty"`
	CustomObjectLinks            []CustomObjectLink `json:"custom_object_links,omitempty"`
	Description                  Description        `json:"description,omitempty"`
	EndorsementsCount            int                `json:"endorsements_count,omitempty"`
	FullTags                     []FullTag          `json:"full_tags,omitempty"`
	InitialVotes                 int                `json:"initial_votes,omitempty"`
	IntegrationFields            []any              `json:"integration_fields,omitempty"`
	Product                      Product            `json:"product,omitempty"`
	ProductID                    string             `json:"product_id,omitempty"`
	Resource                     string             `json:"resource,omitempty"`
	Score                        int                `json:"score,omitempty"`
	ScoreFacts                   []ScoreFact        `json:"score_facts,omitempty"`
	StatusChangedAt              any                `json:"status_changed_at,omitempty"`
	SubmittedIdeaPortalRecordURL string             `json:"submitted_idea_portal_record_url,omitempty"`
	Tags                         []string           `json:"tags,omitempty"`
	URL                          string             `json:"url,omitempty"`
	Visibility                   string             `json:"visibility,omitempty"`
	Votes                        int                `json:"votes,omitempty"`
	WorkflowStatus               WorkflowStatus     `json:"workflow_status,omitempty"`
}

type CreatedByUser struct {
	ID        string    `json:"id"`
	Name      string    `json:"name,omitempty"`
	Email     string    `json:"email,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

type Category struct {
	ID        string    `json:"id"`
	Name      string    `json:"name,omitempty"`
	ParentID  int       `json:"parent_id,omitempty"`
	ProjectID int       `json:"project_id,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}

type CustomObjectLink struct {
	Key        string `json:"key,omitempty"`
	Name       string `json:"name,omitempty"`
	RecordType string `json:"record_type,omitempty"`
	RecordIds  []int  `json:"record_ids,omitempty"`
}
