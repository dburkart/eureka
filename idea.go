package eureka

import "time"

type Idea struct {
	ID                           string             `json:"id"`
	Name                         string             `json:"name,omitempty"`
	ReferenceNum                 string             `json:"reference_num,omitempty"`
	Score                        int                `json:"score,omitempty"`
	CreatedAt                    time.Time          `json:"created_at,omitempty"`
	UpdatedAt                    time.Time          `json:"updated_at,omitempty"`
	ProductID                    string             `json:"product_id,omitempty"`
	Votes                        int                `json:"votes,omitempty"`
	InitialVotes                 int                `json:"initial_votes,omitempty"`
	StatusChangedAt              any                `json:"status_changed_at,omitempty"`
	WorkflowStatus               WorkflowStatus     `json:"workflow_status,omitempty"`
	Description                  Description        `json:"description,omitempty"`
	Visibility                   string             `json:"visibility,omitempty"`
	URL                          string             `json:"url,omitempty"`
	Resource                     string             `json:"resource,omitempty"`
	Product                      Product            `json:"product,omitempty"`
	CreatedByUser                CreatedByUser      `json:"created_by_user,omitempty"`
	AssignedToUser               any                `json:"assigned_to_user,omitempty"`
	EndorsementsCount            int                `json:"endorsements_count,omitempty"`
	CommentsCount                int                `json:"comments_count,omitempty"`
	ScoreFacts                   []ScoreFact        `json:"score_facts,omitempty"`
	Tags                         []string           `json:"tags,omitempty"`
	FullTags                     []FullTag          `json:"full_tags,omitempty"`
	Categories                   []Category         `json:"categories,omitempty"`
	CustomFields                 []CustomField      `json:"custom_fields,omitempty"`
	CustomObjectLinks            []CustomObjectLink `json:"custom_object_links,omitempty"`
	SubmittedIdeaPortalRecordURL string             `json:"submitted_idea_portal_record_url,omitempty"`
	IntegrationFields            []any              `json:"integration_fields,omitempty"`
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
