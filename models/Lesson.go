package models

// Lesson modal
type Lesson struct {
	InfoMeta
	SectionID uint `gorm:"column:section_id" json:"section_id"`
}
