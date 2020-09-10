package models

import (
	"OnlineCourses/models/types"
	"errors"
	"strconv"
)

// Section struct
type Section struct {
	InfoMeta
	CourseID  *uint64     `gorm:"column:course_id" json:"course_id,string"`
	SectionID *uint64     `gorm:"column:parent_id" json:"parent_id,string" sql:"default:null"`
	Lesson    LessonGroup `json:"lessons"` // gorm:"association_autoupdate:false" Commented temporarily
}

// SectionGroup groups section
type SectionGroup []Section

// AfterClone of section
func (section Section) AfterClone() Section {
	section.SectionID = section.ID
	section.ID = nil
	section.CourseID = nil
	section.Lesson = section.Lesson.GroupAfterClone()
	return section
}

// GetPKID for section
func (section Section) GetPKID() *uint64 {
	return section.ID
}

// ValidateOnPublish section for publish
func (section Section) ValidateOnPublish() error {
	if section.Status == types.STATUS_MERGED || section.Status == types.STATUS_PUBLISHED {
		return section.Lesson.GroupValidation()
	}
	if section.SectionID != nil {
		return errors.New("Kindly merge the Section " + strconv.FormatUint(*section.ID, 10) + " with " + strconv.FormatUint(*section.SectionID, 10))
	}
	// TODO : Fix me
	return errors.New("Kindly publish/Save the Section")
}

func (section Section) BeforePublish() Section {
	if section.SectionID != nil {
		section.ID = section.SectionID
		section.SectionID = nil
	}
	section.Status = types.STATUS_PUBLISHED
	section.Lesson = section.Lesson.GroupBeforePublish()
	return section
}

// GroupAfterClone invoke afterClone for each entity
func (sectionGroup SectionGroup) GroupAfterClone() SectionGroup {
	sections := []Section{}
	for _, section := range sectionGroup {
		sections = append(sections, section.AfterClone())
	}
	return sections
}

// GroupValidation need to be invoked via reflection instead of calling it directly
// Deprecated
func (sectionGroup SectionGroup) GroupValidation() error {
	for _, section := range sectionGroup {
		err := section.ValidateOnPublish()
		if err != nil {
			return err
		}
	}
	return nil
}

// GroupBeforePublish invoke afterClone for each entity
func (sectionGroup SectionGroup) GroupBeforePublish() SectionGroup {
	sections := []Section{}
	for _, section := range sectionGroup {
		sections = append(sections, section.BeforePublish())
	}
	return sections
}
