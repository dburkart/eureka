package eureka

import "time"

type Audit struct {
	ID             string    `json:"id"`
	AuditAction    string    `json:"audit_action"`
	CreatedAt      time.Time `json:"created_at"`
	Interesting    bool      `json:"interesting"`
	User           User      `json:"user"`
	Contributors   []User    `json:"contributors"`
	AuditableType  string    `json:"auditable_type"`
	AuditableID    string    `json:"auditable_id"`
	AssociatedType string    `json:"associated_type"`
	AssociatedID   string    `json:"associated_id"`
	Description    string    `json:"description"`
	AuditableURL   string    `json:"auditable_url"`
	Changes        []struct {
		FieldName string `json:"field_name"`
		Value     string `json:"value"`
	} `json:"changes"`
}

type AuditEvent struct {
	Event string `json:"event"`
	Audit Audit  `json:"audit"`
}
