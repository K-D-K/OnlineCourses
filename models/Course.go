package models

// Course Meta data
type Course struct {
	InfoMeta
	CourseID *uint     `gorm:"column:parent_id" json:"parent_id,string" sql:"default:null"`
	Section  []Section `gorm:"association_autoupdate:false" json:"sections"`
}
