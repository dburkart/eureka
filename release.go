package eureka

import "time"

type Release struct {
	RefObject
	CommentsCount                  int                `json:"comments_count,omitempty"`
	CreatedByUser                  User               `json:"created_by_user,omitempty"`
	CustomFields                   []CustomField      `json:"custom_fields,omitempty"`
	DevelopmentStartedOn           string             `json:"development_started_on,omitempty"`
	DurationSource                 string             `json:"duration_source,omitempty"`
	EndDate                        any                `json:"end_date,omitempty"`
	ExternalDateResolution         string             `json:"external_date_resolution,omitempty"`
	ExternalReleaseDate            string             `json:"external_release_date,omitempty"`
	ExternalReleaseDateDescription string             `json:"external_release_date_description,omitempty"`
	Goals                          []Goals            `json:"goals,omitempty"`
	Initiatives                    []Initiative       `json:"initiatives,omitempty"`
	IntegrationFields              []IntegrationField `json:"integration_fields,omitempty"`
	KeyResults                     []any              `json:"key_results,omitempty"`
	MasterRelease                  bool               `json:"master_release,omitempty"`
	Owner                          Owner              `json:"owner,omitempty"`
	ParkingLot                     bool               `json:"parking_lot,omitempty"`
	Position                       any                `json:"position,omitempty"`
	Progress                       any                `json:"progress,omitempty"`
	ProgressSource                 string             `json:"progress_source,omitempty"`
	Project                        Project            `json:"project,omitempty"`
	ReleaseDate                    string             `json:"release_date,omitempty"`
	Released                       bool               `json:"released,omitempty"`
	ReleasedOn                     string             `json:"released_on,omitempty"`
	Resource                       string             `json:"resource,omitempty"`
	StartDate                      string             `json:"start_date,omitempty"`
	StatusChangedOn                any                `json:"status_changed_on,omitempty"`
	Theme                          Theme              `json:"theme,omitempty"`
	URL                            string             `json:"url,omitempty"`
	WorkflowStatus                 WorkflowStatus     `json:"workflow_status,omitempty"`
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
