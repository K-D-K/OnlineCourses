package models

// InfoMeta struct contains meta data for models
type InfoMeta struct {
	Model
	Name      string `json:"name"`
	Link      string `json:"source"`
	Status    int    `json:"status"` // Need to create ENUM
	CreatedBy User   `json:"created_by"`
	UpdatedBy User   `json:"updated_by"`
}
