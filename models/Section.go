package models

import (
	"OnlineCourses/interfaces"
)

// Section struct
type Section struct {
	InfoMeta
	CourseID  *uint64     `gorm:"column:course_id" json:"course_id,string"`
	SectionID *uint64     `gorm:"column:parent_id" json:"parent_id,string" sql:"default:null"`
	Lesson    LessonGroup `gorm:"association_autoupdate:false" json:"lessons"`
}

// SectionGroup groups section
type SectionGroup []Section

// AfterClone of section
func (section Section) AfterClone() interfaces.Entity {
	section.SectionID = section.ID
	section.ID = nil
	section.CourseID = nil
	section.Lesson = section.Lesson.GroupAfterClone()
	return section
}

// GetPKID for section
func (section Section) GetPKID() uint64 {
	return *section.SectionID
}

// ValidateOnPublish section for publish
func (section Section) ValidateOnPublish() error {
	return nil
}

// GroupAfterClone invoke afterClone for each entity
func (sectionGroup SectionGroup) GroupAfterClone() SectionGroup {
	sections := []Section{}
	for _, section := range sectionGroup {
		(&section).AfterClone()
		sections = append(sections, section)
	}
	return sections
}
