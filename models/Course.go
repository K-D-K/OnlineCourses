package models

import (
	"github.com/jinzhu/copier"
)

// Course Meta data
type Course struct {
	InfoMeta
	CourseID *uint64      `gorm:"column:parent_id" json:"parent_id,string" sql:"default:null"`
	Section  SectionGroup `gorm:"association_autoupdate:false" json:"sections"`
}

// AfterClone for Course
func (course Course) AfterClone() Course {
	course.CourseID = course.ID
	course.ID = nil
	course.Section = course.Section.GroupAfterClone()
	return course
}

// GetPKID for that record
func (course Course) GetPKID() uint64 {
	return *course.CourseID
}

// ValidateOnPublish for course
func (course Course) ValidateOnPublish() error {
	return nil
}

// Clone for Course
func (course Course) Clone() Course {
	dest := Course{}
	copier.Copy(&dest, &course)
	dest = dest.AfterClone()
	return dest
}
