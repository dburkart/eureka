package eureka

import "time"

type Product struct {
	ID                              string                      `json:"id,omitempty"`
	ReferencePrefix                 string                      `json:"reference_prefix,omitempty"`
	Name                            string                      `json:"name,omitempty"`
	ProductLine                     bool                        `json:"product_line,omitempty"`
	ProductLineType                 any                         `json:"product_line_type,omitempty"`
	CreatedAt                       time.Time                   `json:"created_at,omitempty"`
	UpdatedAt                       time.Time                   `json:"updated_at,omitempty"`
	Description                     Description                 `json:"description,omitempty"`
	URL                             string                      `json:"url,omitempty"`
	Resource                        string                      `json:"resource,omitempty"`
	Children                        []any                       `json:"children,omitempty"`
	CustomFields                    []CustomField               `json:"custom_fields,omitempty"`
	ScreenDefinitionIds             ScreenDefinitionIds         `json:"screen_definition_ids,omitempty"`
	ScreenDefinitions               []ScreenDefinition          `json:"screen_definitions,omitempty"`
	HasIdeas                        bool                        `json:"has_ideas,omitempty"`
	HasMasterFeatures               bool                        `json:"has_master_features,omitempty"`
	HasEpics                        bool                        `json:"has_epics,omitempty"`
	ReleaseWorkflow                 ReleaseWorkflow             `json:"release_workflow,omitempty"`
	RequirementWorkflow             RequirementWorkflow         `json:"requirement_workflow,omitempty"`
	FeatureWorkflow                 FeatureWorkflow             `json:"feature_workflow,omitempty"`
	InitiativeWorkflow              InitiativeWorkflow          `json:"initiative_workflow,omitempty"`
	EpicWorkflow                    EpicWorkflow                `json:"epic_workflow,omitempty"`
	IdeaWorkflow                    IdeaWorkflow                `json:"idea_workflow,omitempty"`
	StrategicImperativeWorkflow     StrategicImperativeWorkflow `json:"strategic_imperative_workflow,omitempty"`
	KeyResultWorkflow               KeyResultWorkflow           `json:"key_result_workflow,omitempty"`
	CapacityPlanningEnabled         bool                        `json:"capacity_planning_enabled,omitempty"`
	DefaultCapacityUnits            int                         `json:"default_capacity_units,omitempty"`
	EnhancedCapacityPlanningEnabled bool                        `json:"enhanced_capacity_planning_enabled,omitempty"`
	WorkspaceType                   string                      `json:"workspace_type,omitempty"`
}

type ScreenDefinitionIds struct {
	Competitor          int `json:"Competitor,omitempty"`
	Epic                int `json:"Epic,omitempty"`
	Feature             int `json:"Feature,omitempty"`
	IdeasIdea           int `json:"Ideas::Idea,omitempty"`
	IdeasIdeaPortal     int `json:"Ideas::IdeaPortal,omitempty"`
	Initiative          int `json:"Initiative,omitempty"`
	Page                int `json:"Page,omitempty"`
	Project             int `json:"Project,omitempty"`
	Release             int `json:"Release,omitempty"`
	Requirement         int `json:"Requirement,omitempty"`
	StrategicImperative int `json:"StrategicImperative,omitempty"`
}

type Options struct {
	ID    int    `json:"id,omitempty"`
	Label string `json:"label,omitempty"`
}

type CustomFieldDefinitions struct {
	ID                   int       `json:"id,omitempty"`
	Key                  string    `json:"key,omitempty"`
	Position             int       `json:"position,omitempty"`
	Name                 string    `json:"name,omitempty"`
	Type                 string    `json:"type,omitempty"`
	APIType              string    `json:"api_type,omitempty"`
	Required             bool      `json:"required,omitempty"`
	ConfigurationDisplay any       `json:"configuration_display,omitempty"`
	Options              []Options `json:"options,omitempty"`
}

type ScreenDefinition struct {
	ID                     int                      `json:"id,omitempty"`
	ScreenableType         string                   `json:"screenable_type,omitempty"`
	Name                   string                   `json:"name,omitempty"`
	CustomFieldDefinitions []CustomFieldDefinitions `json:"custom_field_definitions,omitempty"`
}

type ReleaseWorkflow struct {
	ID             int    `json:"id,omitempty"`
	Name           string `json:"name,omitempty"`
	StatusableType string `json:"statusable_type,omitempty"`
	WorkflowType   string `json:"workflow_type,omitempty"`
}

type RequirementWorkflow struct {
	ID             int    `json:"id,omitempty"`
	Name           string `json:"name,omitempty"`
	StatusableType string `json:"statusable_type,omitempty"`
	WorkflowType   string `json:"workflow_type,omitempty"`
}

type FeatureWorkflow struct {
	ID             int    `json:"id,omitempty"`
	Name           string `json:"name,omitempty"`
	StatusableType string `json:"statusable_type,omitempty"`
	WorkflowType   string `json:"workflow_type,omitempty"`
}

type InitiativeWorkflow struct {
	ID             int    `json:"id,omitempty"`
	Name           string `json:"name,omitempty"`
	StatusableType string `json:"statusable_type,omitempty"`
	WorkflowType   string `json:"workflow_type,omitempty"`
}

type EpicWorkflow struct {
	ID             int    `json:"id,omitempty"`
	Name           string `json:"name,omitempty"`
	StatusableType string `json:"statusable_type,omitempty"`
	WorkflowType   string `json:"workflow_type,omitempty"`
}

type IdeaWorkflow struct {
	ID             int    `json:"id,omitempty"`
	Name           string `json:"name,omitempty"`
	StatusableType string `json:"statusable_type,omitempty"`
	WorkflowType   string `json:"workflow_type,omitempty"`
}

type StrategicImperativeWorkflow struct {
	ID             int    `json:"id,omitempty"`
	Name           string `json:"name,omitempty"`
	StatusableType string `json:"statusable_type,omitempty"`
	WorkflowType   string `json:"workflow_type,omitempty"`
}

type KeyResultWorkflow struct {
	ID             int    `json:"id,omitempty"`
	Name           string `json:"name,omitempty"`
	StatusableType string `json:"statusable_type,omitempty"`
	WorkflowType   string `json:"workflow_type,omitempty"`
}
