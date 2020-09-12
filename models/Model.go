package models

import "time"

// Model : default Modal
type Model struct {
	ID        *uint64    `gorm:"primary_key" json:"id,string"`
	CreatedAt time.Time  `json:"created_at" tazapay:"restrict_manual:true;"`
	UpdatedAt time.Time  `json:"updated_at" tazapay:"restrict_manual:true;"`
	DeletedAt *time.Time `sql:"index" json:"deleted_at" tazapay:"restrict_manual:true;"`
}
