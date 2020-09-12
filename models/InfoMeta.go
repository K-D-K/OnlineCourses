package models

import (
	"OnlineCourses/models/types/status"
)

// InfoMeta struct contains meta data for models
type InfoMeta struct {
	Model
	Name        string        `json:"name"`
	Link        string        `json:"source"`
	Status      status.Status `json:"status" gorm:"type:integer"`
	CreatedByID *uint64       `json:"-" gorm:"column:created_by" sql:"default:null"` // JSON is not exposed so there is no need to add restrict_manual
	UpdatedByID *uint64       `json:"-" gorm:"column:updated_by" sql:"default:null"` // JSON is not exposed so there is no need to add restrict_manual
}

// SetCreatedBy for InfoMeta
func (model *InfoMeta) SetCreatedBy(createdBy User) {
	if model.CreatedByID == nil {
		model.CreatedByID = createdBy.ID
	}
	model.SetUpdatedBy(createdBy)
}

// SetUpdatedBy for Info Meta
func (model *InfoMeta) SetUpdatedBy(updatedBy User) {
	model.UpdatedByID = updatedBy.ID
}
