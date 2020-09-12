package models

import (
	"OnlineCourses/interfaces"
	"OnlineCourses/models/types/entity"
	"OnlineCourses/models/types/status"
	"errors"
	"strconv"
)

// Course Meta data
type Course struct {
	InfoMeta
	CourseID *uint64   `gorm:"column:parent_id" json:"-" sql:"default:null"` // JSON is not exposed so there is no need to add restrict_manual
	Section  []Section `json:"sections" gorm:"association_autoupdate:false;" tazapay:"restrict_manual:true;child_entity:true"`
}

// CourseGroup .
type CourseGroup []Course

// Name of the modal
func (course *Course) Name() entity.Entity {
	return entity.COURSE
}

// GetPKID for that record
func (course *Course) GetPKID() *uint64 {
	return course.ID
}

// ValidateOnPublish for course
func (course *Course) ValidateOnPublish() error {
	if course.Status == status.STATUS_MERGED || course.Status == status.STATUS_PUBLISHED {
		return SectionGroup(course.Section).GroupValidation()
	}
	if course.CourseID != nil {
		return errors.New("Kindly merge the Course " + strconv.FormatUint(*course.ID, 10) + " with " + strconv.FormatUint(*course.CourseID, 10))
	}
	// TODO : Fix me
	return errors.New("Kindly publish/Save the Course")
}

// BeforePublish .
func (course Course) BeforePublish() Course {
	if course.CourseID != nil {
		course.ID = course.CourseID
		course.CourseID = nil
	}
	course.Status = status.STATUS_PUBLISHED
	course.Section = SectionGroup(course.Section).GroupBeforePublish()
	return course
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

// UpdateParentID .
func (course *Course) UpdateParentID(parentID *uint64) {
	course.CourseID = parentID
}

// UpdateRelationID .
func (course *Course) UpdateRelationID(relID *uint64) {
	return
}

// ResetPKID .
func (course *Course) ResetPKID() {
	course.ID = nil
}

// SetStatus .
func (course *Course) SetStatus(status status.Status) {
	course.Status = status
}

// GetStatus .
func (course *Course) GetStatus() status.Status {
	return course.Status
}

func convertCourseIntoEntityArr(courses []Course) []interfaces.Entity {
	var entities []interfaces.Entity
	for _, course := range courses {
		entities = append(entities, &course)
	}
	return entities
}
