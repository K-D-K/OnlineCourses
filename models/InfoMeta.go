package models

import "OnlineCourses/models/types"

// InfoMeta struct contains meta data for models
type InfoMeta struct {
	Model
	Name        string       `json:"name"`
	Link        string       `json:"source"`
	Status      types.Status `json:"status" gorm:"type:integer"`
	CreatedByID *uint        `json:"-" gorm:"column:created_by" sql:"default:null"`
	UpdatedByID *uint        `json:"-" gorm:"column:updated_by" sql:"default:null"`
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
