package models

import (
	"OnlineCourses/interfaces"
	"OnlineCourses/models/types/entity"
	"OnlineCourses/models/types/status"
)

// Course Meta data
type Course struct {
	InfoMeta
	CourseID *uint64   `gorm:"column:parent_id" json:"-" sql:"default:null"` // JSON is not exposed so there is no need to add restrict_manual
	Section  []Section `json:"sections" tazapay:"restrict_manual:true;child_entity:true"`
}

// Name of the modal
func (course *Course) Name() entity.Entity {
	return entity.COURSE
}

// GetPKID for that record
func (course *Course) GetPKID() *uint64 {
	return course.ID
}

// SetPKID .
func (course *Course) SetPKID(pkID *uint64) {
	course.ID = pkID
}

// GetParentID .
func (course *Course) GetParentID() *uint64 {
	return course.CourseID
}

// SetParentID .
func (course *Course) SetParentID(parentID *uint64) {
	course.CourseID = parentID
}

// GetChildEntities .
func (course *Course) GetChildEntities() map[string][]interfaces.Entity {
	entitiesMap := make(map[string][]interfaces.Entity)
	entitiesMap[entity.SECTION] = convertSectionToEntityArr(course.Section)
	return entitiesMap
}

// SetChildEntities .
func (course *Course) SetChildEntities(entitiesMap map[string][]interfaces.Entity) {
	course.Section = convertEntityToSectionArr(entitiesMap[entity.SECTION])
}

// SetStatus .
func (course *Course) SetStatus(status status.Status) {
	course.Status = status
}

// GetStatus .
func (course *Course) GetStatus() status.Status {
	return course.Status
}

// SetRelationID .
func (course *Course) SetRelationID(relID *uint64) {
	return
}

// IsDeleted .
func (course *Course) IsDeleted() bool {
	return course.Delete
}

// ConvertCourseIntoEntityArr .
func ConvertCourseIntoEntityArr(courses []Course) []interfaces.Entity {
	var entities []interfaces.Entity
	for _, course := range courses {
		entities = append(entities, &course)
	}
	return entities
}

// ConvertEntityToCourseArr .
func ConvertEntityToCourseArr(entities []interfaces.Entity) []Course {
	courses := make([]Course, len(entities))
	for index, entity := range entities {
		course := entity.(*Course)
		courses[index] = *course
	}
	return courses
}
