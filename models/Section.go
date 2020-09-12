package models

import (
	"OnlineCourses/interfaces"
	"OnlineCourses/models/types/entity"
	"OnlineCourses/models/types/status"
)

// Section struct
type Section struct {
	InfoMeta
	CourseID  *uint64  `gorm:"column:course_id" json:"course_id,string"`
	SectionID *uint64  `gorm:"column:parent_id" json:"-" sql:"default:null"` // JSON is not exposed so there is no need to add restrict_manual
	Lesson    []Lesson `json:"lessons" tazapay:"restrict_manual:true;child_entity:true"`
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

// SetPKID .
func (section *Section) SetPKID(pkID *uint64) {
	section.ID = pkID
}

// GetParentID .
func (section *Section) GetParentID() *uint64 {
	return section.SectionID
}

// SetParentID .
func (section *Section) SetParentID(parentID *uint64) {
	section.SectionID = parentID
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

// UpdateRelationID .
func (section *Section) UpdateRelationID(relID *uint64) {
	section.CourseID = relID
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
