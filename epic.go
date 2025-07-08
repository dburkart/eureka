package eureka

import "time"

type Epic struct {
	ID                                 string             `json:"id"`
	Name                               string             `json:"name,omitempty"`
	ReferenceNum                       string             `json:"reference_num,omitempty"`
	InitiativeReferenceNum             string             `json:"initiative_reference_num,omitempty"`
	Position                           int                `json:"position,omitempty"`
	Score                              int                `json:"score,omitempty"`
	CreatedAt                          time.Time          `json:"created_at,omitempty"`
	UpdatedAt                          time.Time          `json:"updated_at,omitempty"`
	StartDate                          string             `json:"start_date,omitempty"`
	DueDate                            string             `json:"due_date,omitempty"`
	ProductID                          string             `json:"product_id,omitempty"`
	Progress                           any                `json:"progress,omitempty"`
	ProgressSource                     string             `json:"progress_source,omitempty"`
	DurationSource                     string             `json:"duration_source,omitempty"`
	StatusChangedOn                    any                `json:"status_changed_on,omitempty"`
	CreatedByUser                      User               `json:"created_by_user,omitempty"`
	WorkflowStatus                     WorkflowStatus     `json:"workflow_status,omitempty"`
	Project                            Project            `json:"project,omitempty"`
	Description                        Description        `json:"description,omitempty"`
	Attachments                        []any              `json:"attachments,omitempty"`
	IntegrationFields                  []IntegrationField `json:"integration_fields,omitempty"`
	URL                                string             `json:"url,omitempty"`
	Resource                           string             `json:"resource,omitempty"`
	Release                            Release            `json:"release,omitempty"`
	AssignedToUser                     AssignedToUser     `json:"assigned_to_user,omitempty"`
	Features                           []Feature          `json:"features,omitempty"`
	Initiative                         Initiative         `json:"initiative,omitempty"`
	Goals                              []Goals            `json:"goals,omitempty"`
	KeyResults                         []any              `json:"key_results,omitempty"`
	CommentsCount                      int                `json:"comments_count,omitempty"`
	ScoreFacts                         []any              `json:"score_facts,omitempty"`
	Tags                               []string           `json:"tags,omitempty"`
	FullTags                           []FullTag          `json:"full_tags,omitempty"`
	CustomFields                       []CustomField      `json:"custom_fields,omitempty"`
	EpicLinks                          []EpicLink         `json:"epic_links,omitempty"`
	MasterFeatureOnlyOriginalEstimate  any                `json:"master_feature_only_original_estimate,omitempty"`
	MasterFeatureOnlyRemainingEstimate any                `json:"master_feature_only_remaining_estimate,omitempty"`
	MasterFeatureOnlyWorkDone          any                `json:"master_feature_only_work_done,omitempty"`
	EpicOnlyOriginalEstimate           any                `json:"epic_only_original_estimate,omitempty"`
	EpicOnlyRemainingEstimate          any                `json:"epic_only_remaining_estimate,omitempty"`
	EpicOnlyWorkDone                   any                `json:"epic_only_work_done,omitempty"`
}

type EpicLink struct {
	LinkType     string       `json:"link_type,omitempty"`
	LinkTypeID   int          `json:"link_type_id,omitempty"`
	CreatedAt    time.Time    `json:"created_at,omitempty"`
	ParentRecord ParentRecord `json:"parent_record,omitempty"`
	ChildRecord  ChildRecord  `json:"child_record,omitempty"`
}
