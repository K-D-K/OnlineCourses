package models

import (
	"github.com/jinzhu/copier"
)

// Course Meta data
type Course struct {
	InfoMeta
	CourseID *uint     `gorm:"column:parent_id" json:"parent_id,string" sql:"default:null"`
	Section  []Section `gorm:"association_autoupdate:false" json:"sections"`
}

// AfterClone for Course
func (course *Course) AfterClone() {
	course.CourseID = course.ID
	course.ID = nil
	sections := []Section{}
	for _, section := range course.Section {
		(&section).AfterClone()
		sections = append(sections, section)
	}
	course.Section = sections
}

// Clone for Course
func (course Course) Clone() Course {
	dest := &Course{}
	copier.Copy(dest, &course)
	dest.AfterClone()
	return *dest
}
