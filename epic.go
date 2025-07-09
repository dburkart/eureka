package eureka

import "time"

type Epic struct {
	RefObject
	AssignedToUser                     AssignedToUser     `json:"assigned_to_user,omitempty"`
	Attachments                        []any              `json:"attachments,omitempty"`
	CommentsCount                      int                `json:"comments_count,omitempty"`
	CreatedByUser                      User               `json:"created_by_user,omitempty"`
	CustomFields                       []CustomField      `json:"custom_fields,omitempty"`
	Description                        Description        `json:"description,omitempty"`
	DueDate                            string             `json:"due_date,omitempty"`
	DurationSource                     string             `json:"duration_source,omitempty"`
	EpicLinks                          []EpicLink         `json:"epic_links,omitempty"`
	EpicOnlyOriginalEstimate           any                `json:"epic_only_original_estimate,omitempty"`
	EpicOnlyRemainingEstimate          any                `json:"epic_only_remaining_estimate,omitempty"`
	EpicOnlyWorkDone                   any                `json:"epic_only_work_done,omitempty"`
	Features                           []Feature          `json:"features,omitempty"`
	FullTags                           []FullTag          `json:"full_tags,omitempty"`
	Goals                              []Goals            `json:"goals,omitempty"`
	Initiative                         Initiative         `json:"initiative,omitempty"`
	InitiativeReferenceNum             string             `json:"initiative_reference_num,omitempty"`
	IntegrationFields                  []IntegrationField `json:"integration_fields,omitempty"`
	KeyResults                         []any              `json:"key_results,omitempty"`
	MasterFeatureOnlyOriginalEstimate  any                `json:"master_feature_only_original_estimate,omitempty"`
	MasterFeatureOnlyRemainingEstimate any                `json:"master_feature_only_remaining_estimate,omitempty"`
	MasterFeatureOnlyWorkDone          any                `json:"master_feature_only_work_done,omitempty"`
	Position                           int                `json:"position,omitempty"`
	ProductID                          string             `json:"product_id,omitempty"`
	Progress                           any                `json:"progress,omitempty"`
	ProgressSource                     string             `json:"progress_source,omitempty"`
	Project                            Project            `json:"project,omitempty"`
	Release                            Release            `json:"release,omitempty"`
	Resource                           string             `json:"resource,omitempty"`
	Score                              int                `json:"score,omitempty"`
	ScoreFacts                         []any              `json:"score_facts,omitempty"`
	StartDate                          string             `json:"start_date,omitempty"`
	StatusChangedOn                    any                `json:"status_changed_on,omitempty"`
	Tags                               []string           `json:"tags,omitempty"`
	URL                                string             `json:"url,omitempty"`
	WorkflowStatus                     WorkflowStatus     `json:"workflow_status,omitempty"`
}

type EpicLink struct {
	LinkType     string       `json:"link_type,omitempty"`
	LinkTypeID   int          `json:"link_type_id,omitempty"`
	CreatedAt    time.Time    `json:"created_at,omitempty"`
	ParentRecord ParentRecord `json:"parent_record,omitempty"`
	ChildRecord  ChildRecord  `json:"child_record,omitempty"`
}
