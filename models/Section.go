package models

// Section struct
type Section struct {
	InfoMeta
	CourseID uint     `gorm:"column:course_id" json:"course_id"`
	Lesson   []Lesson `gorm:"association_autoupdate:false" json:"lessons"`
}
