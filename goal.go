package eureka

import "time"

type Goal struct {
	RefObject
	Color          string        `json:"color,omitempty"`
	CommentsCount  int           `json:"comments_count,omitempty"`
	CustomFields   []any         `json:"custom_fields,omitempty"`
	Description    Description   `json:"description,omitempty"`
	Effort         int           `json:"effort,omitempty"`
	Features       []any         `json:"features,omitempty"`
	Initiatives    []any         `json:"initiatives,omitempty"`
	KeyResults     []any         `json:"key_results,omitempty"`
	Parent         Parent        `json:"parent,omitempty"`
	Parents        []Parent      `json:"parents,omitempty"`
	Position       int           `json:"position,omitempty"`
	ProductID      string        `json:"product_id,omitempty"`
	Progress       int           `json:"progress,omitempty"`
	ProgressSource string        `json:"progress_source,omitempty"`
	Project        Project       `json:"project,omitempty"`
	Releases       []any         `json:"releases,omitempty"`
	Resource       string        `json:"resource,omitempty"`
	SuccessMetric  SuccessMetric `json:"success_metric,omitempty"`
	URL            string        `json:"url,omitempty"`
	Value          int           `json:"value,omitempty"`
}

type WorkflowStatusTimes struct {
	StatusID   string    `json:"status_id,omitempty"`
	StatusName string    `json:"status_name,omitempty"`
	StartedAt  time.Time `json:"started_at,omitempty"`
	EndedAt    any       `json:"ended_at,omitempty"`
}

type SuccessMetric struct {
	Name                string                `json:"name,omitempty"`
	Description         Description           `json:"description,omitempty"`
	WorkflowStatus      WorkflowStatus        `json:"workflow_status,omitempty"`
	WorkflowStatusTimes []WorkflowStatusTimes `json:"workflow_status_times,omitempty"`
}

type Parent struct {
	ID          string      `json:"id,omitempty"`
	Name        string      `json:"name,omitempty"`
	URL         string      `json:"url,omitempty"`
	Resource    string      `json:"resource,omitempty"`
	CreatedAt   time.Time   `json:"created_at,omitempty"`
	Description Description `json:"description,omitempty"`
}
