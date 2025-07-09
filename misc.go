package eureka

import "time"

type RefObject struct {
	ID           string    `json:"id"`
	ReferenceNum string    `json:"reference_num"`
	CreatedAt    time.Time `json:"created_at,omitempty"`
	UpdatedAt    time.Time `json:"updated_at,omitempty"`
	Name         string    `json:"name,omitempty"`
}

type Attachments struct {
	ID               string    `json:"id"`
	DownloadURL      string    `json:"download_url,omitempty"`
	CreatedAt        time.Time `json:"created_at,omitempty"`
	UpdatedAt        time.Time `json:"updated_at,omitempty"`
	OriginalFileSize int       `json:"original_file_size,omitempty"`
	ContentType      string    `json:"content_type,omitempty"`
	FileName         string    `json:"file_name,omitempty"`
	FileSize         int       `json:"file_size,omitempty"`
}

type Description struct {
	ID          string        `json:"id"`
	Body        string        `json:"body,omitempty"`
	CreatedAt   time.Time     `json:"created_at,omitempty"`
	UpdatedAt   time.Time     `json:"updated_at,omitempty"`
	Attachments []Attachments `json:"attachments,omitempty"`
}

type Owner struct {
	ID        string    `json:"id"`
	Name      string    `json:"name,omitempty"`
	Email     string    `json:"email,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

type Project struct {
	ID              string    `json:"id"`
	ReferencePrefix string    `json:"reference_prefix,omitempty"`
	Name            string    `json:"name,omitempty"`
	ProductLine     bool      `json:"product_line,omitempty"`
	CreatedAt       time.Time `json:"created_at,omitempty"`
	WorkspaceType   string    `json:"workspace_type,omitempty"`
	URL             string    `json:"url,omitempty"`
}

type ScoreFact struct {
	ID    string `json:"id"`
	Value int    `json:"value,omitempty"`
	Name  string `json:"name,omitempty"`
}

type WorkflowStatus struct {
	ID       string `json:"id"`
	Name     string `json:"name,omitempty"`
	Position int    `json:"position,omitempty"`
	Complete bool   `json:"complete,omitempty"`
	Color    string `json:"color,omitempty"`
}

type IntegrationField struct {
	ID            string    `json:"id"`
	Name          string    `json:"name,omitempty"`
	Value         string    `json:"value,omitempty"`
	IntegrationID int       `json:"integration_id,omitempty"`
	ServiceName   string    `json:"service_name,omitempty"`
	CreatedAt     time.Time `json:"created_at,omitempty"`
}

type CustomField struct {
	ID          string    `json:"id,omitempty"`
	Key         string    `json:"key,omitempty"`
	Name        string    `json:"name,omitempty"`
	Body        string    `json:"body,omitempty"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
	Attachments []any     `json:"attachments,omitempty"`
	Value       any       `json:"value,omitempty"`
	Type        string    `json:"type,omitempty"`
}

type Comment struct {
	ID          string        `json:"id,omitempty"`
	Body        string        `json:"body,omitempty"`
	CreatedAt   time.Time     `json:"created_at,omitempty"`
	UpdatedAt   time.Time     `json:"updated_at,omitempty"`
	User        *User         `json:"user,omitempty"`
	Attachments []Attachments `json:"attachments,omitempty"`
	URL         *string       `json:"url,omitempty"`
	Resource    *string       `json:"resource,omitempty"`
	Commentable *Commentable  `json:"commentable,omitempty"`
}

type Attachment struct {
	ID               string    `json:"id"`
	DownloadURL      string    `json:"download_url,omitempty"`
	CreatedAt        time.Time `json:"created_at,omitempty"`
	UpdatedAt        time.Time `json:"updated_at,omitempty"`
	OriginalFileSize int       `json:"original_file_size,omitempty"`
	ContentType      string    `json:"content_type,omitempty"`
	FileName         string    `json:"file_name,omitempty"`
	FileSize         int       `json:"file_size,omitempty"`
}

type Commentable struct {
	ID        string `json:"id"`
	Type      string `json:"type,omitempty"`
	ProductID string `json:"product_id,omitempty"`
	URL       string `json:"url,omitempty"`
	Resource  string `json:"resource,omitempty"`
}
