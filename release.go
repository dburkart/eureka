package eureka

import "time"

type Release struct {
	ID                             string             `json:"id"`
	ProductID                      string             `json:"product_id,omitempty"`
	ReferenceNum                   string             `json:"reference_num,omitempty"`
	Name                           string             `json:"name,omitempty"`
	StartDate                      string             `json:"start_date,omitempty"`
	EndDate                        any                `json:"end_date,omitempty"`
	DevelopmentStartedOn           string             `json:"development_started_on,omitempty"`
	ReleaseDate                    string             `json:"release_date,omitempty"`
	ExternalReleaseDate            string             `json:"external_release_date,omitempty"`
	ExternalReleaseDateDescription string             `json:"external_release_date_description,omitempty"`
	ExternalDateResolution         string             `json:"external_date_resolution,omitempty"`
	Released                       bool               `json:"released,omitempty"`
	ParkingLot                     bool               `json:"parking_lot,omitempty"`
	MasterRelease                  bool               `json:"master_release,omitempty"`
	ReleasedOn                     string             `json:"released_on,omitempty"`
	CreatedAt                      time.Time          `json:"created_at,omitempty"`
	UpdatedAt                      time.Time          `json:"updated_at,omitempty"`
	Position                       any                `json:"position,omitempty"`
	Progress                       any                `json:"progress,omitempty"`
	ProgressSource                 string             `json:"progress_source,omitempty"`
	DurationSource                 string             `json:"duration_source,omitempty"`
	StatusChangedOn                any                `json:"status_changed_on,omitempty"`
	Theme                          Theme              `json:"theme,omitempty"`
	URL                            string             `json:"url,omitempty"`
	Resource                       string             `json:"resource,omitempty"`
	IntegrationFields              []IntegrationField `json:"integration_fields,omitempty"`
	CustomFields                   []CustomField      `json:"custom_fields,omitempty"`
	CommentsCount                  int                `json:"comments_count,omitempty"`
	WorkflowStatus                 WorkflowStatus     `json:"workflow_status,omitempty"`
	Owner                          Owner              `json:"owner,omitempty"`
	Goals                          []Goals            `json:"goals,omitempty"`
	KeyResults                     []any              `json:"key_results,omitempty"`
	Initiatives                    []Initiative       `json:"initiatives,omitempty"`
	Project                        Project            `json:"project,omitempty"`
	CreatedByUser                  User               `json:"created_by_user,omitempty"`
}

type Theme struct {
	ID          string    `json:"id"`
	Body        string    `json:"body,omitempty"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
	Attachments []any     `json:"attachments,omitempty"`
}

type Goals struct {
	ID          string      `json:"id"`
	Name        string      `json:"name,omitempty"`
	URL         string      `json:"url,omitempty"`
	Resource    string      `json:"resource,omitempty"`
	CreatedAt   time.Time   `json:"created_at,omitempty"`
	Description Description `json:"description,omitempty"`
}
