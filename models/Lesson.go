package models

import (
	"OnlineCourses/interfaces"
	"OnlineCourses/models/types/entity"
	"OnlineCourses/models/types/status"
)

// Lesson modal
type Lesson struct {
	InfoMeta
	LessonID  *uint64 `gorm:"column:parent_id" json:"-" sql:"default:null"` // JSON is not exposed so there is no need to add restrict_manual
	SectionID *uint64 `gorm:"column:section_id" json:"section_id,string"`
}

// LessonGroup for lessons
type LessonGroup []Lesson

// Name of the modal
func (lesson *Lesson) Name() entity.Entity {
	return entity.LESSON
}

// GetPKID .
func (lesson *Lesson) GetPKID() *uint64 {
	return lesson.ID
}

// SetPKID .
func (lesson *Lesson) SetPKID(pkID *uint64) {
	lesson.ID = pkID
}

// GetParentID .
func (lesson *Lesson) GetParentID() *uint64 {
	return lesson.LessonID
}

// SetParentID .
func (lesson *Lesson) SetParentID(parentID *uint64) {
	lesson.LessonID = parentID
}

// GetChildEntities .
func (lesson *Lesson) GetChildEntities() map[string][]interfaces.Entity {
	return map[string][]interfaces.Entity{}
}

// SetChildEntities .
func (lesson *Lesson) SetChildEntities(entitiesMap map[string][]interfaces.Entity) {
	return
}

// UpdateRelationID .
func (lesson *Lesson) UpdateRelationID(relID *uint64) {
	lesson.SectionID = relID
}

// SetStatus .
func (lesson *Lesson) SetStatus(status status.Status) {
	lesson.Status = status
}

// GetStatus .
func (lesson *Lesson) GetStatus() status.Status {
	return lesson.Status
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
