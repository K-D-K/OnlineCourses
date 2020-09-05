package models

// Lesson modal
type Lesson struct {
	InfoMeta
	LessonID  *uint `gorm:"column:parent_id" json:"parent_id,string" sql:"default:null"`
	SectionID uint  `gorm:"column:section_id" json:"section_id,string"`
}
