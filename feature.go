package eureka

import "time"

type Feature struct {
	RefObject
	AssignedToUser               AssignedToUser     `json:"assigned_to_user,omitempty"`
	Attachments                  []any              `json:"attachments,omitempty"`
	BelongsToReleasePhase        ReleasePhase       `json:"belongs_to_release_phase,omitempty"`
	CommentsCount                int                `json:"comments_count,omitempty"`
	Description                  Description        `json:"description,omitempty"`
	DueDate                      string             `json:"due_date,omitempty"`
	Epic                         Epic               `json:"epic,omitempty"`
	EpicReferenceNum             string             `json:"epic_reference_num,omitempty"`
	FeatureLinks                 []FeatureLink      `json:"feature_links,omitempty"`
	FeatureOnlyOriginalEstimate  any                `json:"feature_only_original_estimate,omitempty"`
	FeatureOnlyRemainingEstimate any                `json:"feature_only_remaining_estimate,omitempty"`
	FeatureOnlyWorkDone          any                `json:"feature_only_work_done,omitempty"`
	FullTags                     []FullTag          `json:"full_tags,omitempty"`
	Goals                        []Goal             `json:"goals,omitempty"`
	Initiative                   Initiative         `json:"initiative,omitempty"`
	InitiativeReferenceNum       string             `json:"initiative_reference_num,omitempty"`
	IntegrationFields            []IntegrationField `json:"integration_fields,omitempty"`
	KeyResults                   []KeyResult        `json:"key_results,omitempty"`
	MasterFeature                MasterFeature      `json:"master_feature,omitempty"`
	Position                     int                `json:"position,omitempty"`
	ProductID                    string             `json:"product_id,omitempty"`
	Progress                     any                `json:"progress,omitempty"`
	ProgressSource               string             `json:"progress_source,omitempty"`
	Release                      Release            `json:"release,omitempty"`
	ReleaseReferenceNum          string             `json:"release_reference_num,omitempty"`
	Requirements                 []Requirement      `json:"requirements,omitempty"`
	Resource                     string             `json:"resource,omitempty"`
	Score                        int                `json:"score,omitempty"`
	ScoreFacts                   []ScoreFact        `json:"score_facts,omitempty"`
	StartDate                    string             `json:"start_date,omitempty"`
	StatusChangedOn              any                `json:"status_changed_on,omitempty"`
	Tags                         []string           `json:"tags,omitempty"`
	URL                          string             `json:"url,omitempty"`
}

type MasterFeature struct {
	ID           string    `json:"id"`
	ReferenceNum string    `json:"reference_num,omitempty"`
	Name         string    `json:"name,omitempty"`
	CreatedAt    time.Time `json:"created_at,omitempty"`
	URL          string    `json:"url,omitempty"`
	Resource     string    `json:"resource,omitempty"`
}

type ReleasePhase struct {
	ID             string      `json:"id"`
	Name           string      `json:"name,omitempty"`
	StartOn        string      `json:"start_on,omitempty"`
	EndOn          string      `json:"end_on,omitempty"`
	Type           string      `json:"type,omitempty"`
	ReleaseID      int         `json:"release_id,omitempty"`
	CreatedAt      time.Time   `json:"created_at,omitempty"`
	UpdatedAt      time.Time   `json:"updated_at,omitempty"`
	Progress       any         `json:"progress,omitempty"`
	ProgressSource string      `json:"progress_source,omitempty"`
	DurationSource string      `json:"duration_source,omitempty"`
	Description    Description `json:"description,omitempty"`
}

type AssignedToUser struct {
	ID              string    `json:"id"`
	Name            string    `json:"name,omitempty"`
	Email           string    `json:"email,omitempty"`
	CreatedAt       time.Time `json:"created_at,omitempty"`
	UpdatedAt       time.Time `json:"updated_at,omitempty"`
	DefaultAssignee bool      `json:"default_assignee,omitempty"`
}

type User struct {
	ID        string    `json:"id"`
	Name      string    `json:"name,omitempty"`
	Email     string    `json:"email,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

type Requirement struct {
	ID                string         `json:"id"`
	Name              string         `json:"name,omitempty"`
	ReferenceNum      string         `json:"reference_num,omitempty"`
	Position          int            `json:"position,omitempty"`
	CreatedAt         time.Time      `json:"created_at,omitempty"`
	UpdatedAt         time.Time      `json:"updated_at,omitempty"`
	ReleaseID         int            `json:"release_id,omitempty"`
	CreatedByUser     User           `json:"created_by_user,omitempty"`
	WorkflowStatus    WorkflowStatus `json:"workflow_status,omitempty"`
	URL               string         `json:"url,omitempty"`
	Resource          string         `json:"resource,omitempty"`
	Description       Description    `json:"description,omitempty"`
	Feature           Feature        `json:"feature,omitempty"`
	AssignedToUser    AssignedToUser `json:"assigned_to_user,omitempty"`
	Attachments       []any          `json:"attachments,omitempty"`
	Tags              []any          `json:"tags,omitempty"`
	FullTags          []any          `json:"full_tags,omitempty"`
	CustomFields      []any          `json:"custom_fields,omitempty"`
	IntegrationFields []any          `json:"integration_fields,omitempty"`
	CommentsCount     int            `json:"comments_count,omitempty"`
}

type Initiative struct {
	ID                string             `json:"id"`
	ReferenceNum      string             `json:"reference_num,omitempty"`
	Name              string             `json:"name,omitempty"`
	URL               string             `json:"url,omitempty"`
	Resource          string             `json:"resource,omitempty"`
	CreatedAt         time.Time          `json:"created_at,omitempty"`
	Description       Description        `json:"description,omitempty"`
	IntegrationFields []IntegrationField `json:"integration_fields,omitempty"`
}

type Goal struct {
	ID          string      `json:"id"`
	Name        string      `json:"name,omitempty"`
	URL         string      `json:"url,omitempty"`
	Resource    string      `json:"resource,omitempty"`
	CreatedAt   time.Time   `json:"created_at,omitempty"`
	Description Description `json:"description,omitempty"`
}

type KeyResult struct {
	ID             string    `json:"id"`
	Name           string    `json:"name,omitempty"`
	ReferenceNum   string    `json:"reference_num,omitempty"`
	URL            string    `json:"url,omitempty"`
	Position       int       `json:"position,omitempty"`
	CreatedAt      time.Time `json:"created_at,omitempty"`
	UpdatedAt      time.Time `json:"updated_at,omitempty"`
	Progress       any       `json:"progress,omitempty"`
	TargetMetric   string    `json:"target_metric,omitempty"`
	StartingMetric string    `json:"starting_metric,omitempty"`
	CurrentMetric  string    `json:"current_metric,omitempty"`
}

type FullTag struct {
	ID    int    `json:"id"`
	Name  string `json:"name,omitempty"`
	Color string `json:"color,omitempty"`
}

type ParentRecord struct {
	ID           string    `json:"id"`
	ReferenceNum string    `json:"reference_num,omitempty"`
	Name         string    `json:"name,omitempty"`
	CreatedAt    time.Time `json:"created_at,omitempty"`
	URL          string    `json:"url,omitempty"`
	Resource     string    `json:"resource,omitempty"`
	ProductID    string    `json:"product_id,omitempty"`
}

type ChildRecord struct {
	ID           string    `json:"id"`
	ReferenceNum string    `json:"reference_num,omitempty"`
	Name         string    `json:"name,omitempty"`
	CreatedAt    time.Time `json:"created_at,omitempty"`
	URL          string    `json:"url,omitempty"`
	Resource     string    `json:"resource,omitempty"`
	ProductID    string    `json:"product_id,omitempty"`
}

type FeatureLink struct {
	LinkType     string       `json:"link_type,omitempty"`
	LinkTypeID   int          `json:"link_type_id,omitempty"`
	CreatedAt    time.Time    `json:"created_at,omitempty"`
	ParentRecord ParentRecord `json:"parent_record,omitempty"`
	ChildRecord  ChildRecord  `json:"child_record,omitempty"`
}
