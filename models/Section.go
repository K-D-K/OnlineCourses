package models

import (
	"OnlineCourses/interfaces"
	"OnlineCourses/models/types/entity"
	"OnlineCourses/models/types/status"
	"errors"
	"strconv"
)

// Section struct
type Section struct {
	InfoMeta
	CourseID  *uint64  `gorm:"column:course_id" json:"course_id,string"`
	SectionID *uint64  `gorm:"column:parent_id" json:"-" sql:"default:null"` // JSON is not exposed so there is no need to add restrict_manual
	Lesson    []Lesson `json:"lessons" gorm:"association_autoupdate:false;" tazapay:"restrict_manual:true;child_entity:true"`
}

// SectionGroup groups section
type SectionGroup []Section

// Name of the modal
func (section *Section) Name() entity.Entity {
	return entity.SECTION
}

// GetPKID for section
func (section *Section) GetPKID() *uint64 {
	return section.ID
}

// ValidateOnPublish section for publish
func (section *Section) ValidateOnPublish() error {
	if section.Status == status.STATUS_MERGED || section.Status == status.STATUS_PUBLISHED {
		return LessonGroup(section.Lesson).GroupValidation()
	}
	if section.SectionID != nil {
		return errors.New("Kindly merge the Section " + strconv.FormatUint(*section.ID, 10) + " with " + strconv.FormatUint(*section.SectionID, 10))
	}
	// TODO : Fix me
	return errors.New("Kindly publish/Save the Section")
}

// BeforePublish .
func (section *Section) BeforePublish() {
	if section.SectionID != nil {
		section.ID = section.SectionID
		section.SectionID = nil
	}
	section.Status = status.STATUS_PUBLISHED
	section.Lesson = LessonGroup(section.Lesson).GroupBeforePublish()
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
		section.BeforePublish()
		sections = append(sections, section)
	}
	return sections
}

// GetChildEntities .
func (section *Section) GetChildEntities() map[string][]interfaces.Entity {
	entitiesMap := make(map[string][]interfaces.Entity)
	entitiesMap[entity.LESSON] = convertLessonToEntityArr(section.Lesson)
	return entitiesMap
}

// SetChildEntities .
func (section *Section) SetChildEntities(entitiesMap map[string][]interfaces.Entity) {
	section.Lesson = convertEntityToLessonArr(entitiesMap[entity.LESSON])
}

// UpdateParentID .
func (section *Section) UpdateParentID(parentID *uint64) {
	section.SectionID = parentID
}

// UpdateRelationID .
func (section *Section) UpdateRelationID(relID *uint64) {
	section.CourseID = relID
}

// ResetPKID .
func (section *Section) ResetPKID() {
	section.ID = nil
}

// SetStatus .
func (section *Section) SetStatus(status status.Status) {
	section.Status = status
}

// GetStatus .
func (section *Section) GetStatus() status.Status {
	return section.Status
}

func convertSectionToEntityArr(sections []Section) []interfaces.Entity {
	entities := make([]interfaces.Entity, len(sections))
	for index, section := range sections {
		sectionClone := section
		entities[index] = &sectionClone
	}
	return entities
}

func convertEntityToSectionArr(entities []interfaces.Entity) []Section {
	sections := make([]Section, len(entities))
	for index, entity := range entities {
		section := entity.(*Section)
		sections[index] = *section
	}
	return sections
}
