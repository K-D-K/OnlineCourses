package models

import (
	"OnlineCourses/interfaces"
	"OnlineCourses/models/types/status"
	"errors"
	"strconv"
)

// Lesson modal
type Lesson struct {
	InfoMeta
	LessonID  *uint64 `gorm:"column:parent_id" json:"-" sql:"default:null"`
	SectionID *uint64 `gorm:"column:section_id" json:"section_id,string"`
}

// LessonGroup for lessons
type LessonGroup []Lesson

// GetPKID .
func (lesson *Lesson) GetPKID() *uint64 {
	return lesson.ID
}

// ValidateOnPublish .
func (lesson *Lesson) ValidateOnPublish() error {
	if lesson.Status == status.STATUS_MERGED || lesson.Status == status.STATUS_PUBLISHED {
		return nil
	}
	if lesson.LessonID != nil {
		return errors.New("Kindly merge the Lesson " + strconv.FormatUint(*lesson.ID, 10) + " with " + strconv.FormatUint(*lesson.LessonID, 10))
	}
	// TODO : Fix me
	return errors.New("Kindly publish/Save the lesson")
}

func (lesson Lesson) BeforePublish() Lesson {
	if lesson.LessonID != nil {
		lesson.ID = lesson.LessonID
		lesson.LessonID = nil
	}
	lesson.Status = status.STATUS_PUBLISHED
	return lesson
}

// GroupValidation need to be invoked via reflection instead of calling it directly
// Deprecated
func (lessons LessonGroup) GroupValidation() error {
	for _, lesson := range lessons {
		err := lesson.ValidateOnPublish()
		if err != nil {
			return err
		}
	}
	return nil
}

// GroupBeforePublish invoke afterClone for each entity
func (lessonGroup LessonGroup) GroupBeforePublish() LessonGroup {
	lessons := []Lesson{}
	for _, lesson := range lessonGroup {
		lessons = append(lessons, lesson.BeforePublish())
	}
	return lessons
}

// GetChildEntities .
func (lesson *Lesson) GetChildEntities() map[string][]interfaces.Entity {
	return map[string][]interfaces.Entity{}
}

// SetChildEntities .
func (lesson *Lesson) SetChildEntities(entitiesMap map[string][]interfaces.Entity) {
	return
}

// UpdateParentID .
func (lesson *Lesson) UpdateParentID(parentID *uint64) {
	lesson.LessonID = parentID
}

// UpdateRelationID .
func (lesson *Lesson) UpdateRelationID(relID *uint64) {
	lesson.SectionID = relID
}

// ResetPKID .
func (lesson *Lesson) ResetPKID() {
	lesson.ID = nil
}

// UpdateStatus .
func (lesson *Lesson) UpdateStatus(status status.Status) {
	lesson.Status = status
}

func convertLessonToEntityArr(lessons []Lesson) []interfaces.Entity {
	entities := make([]interfaces.Entity, len(lessons))
	for index, lesson := range lessons {
		lessonClone := lesson
		entities[index] = &lessonClone
	}
	return entities
}

func convertEntityToLessonArr(entities []interfaces.Entity) []Lesson {
	lessons := make([]Lesson, len(entities))
	for index, entity := range entities {
		lesson := entity.(*Lesson)
		lessons[index] = *lesson
	}
	return lessons
}
