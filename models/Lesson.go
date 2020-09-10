package models

import (
	"OnlineCourses/models/types"
	"errors"
	"strconv"
)

// Lesson modal
type Lesson struct {
	InfoMeta
	LessonID  *uint64 `gorm:"column:parent_id" json:"parent_id,string" sql:"default:null"`
	SectionID *uint64 `gorm:"column:section_id" json:"section_id,string"`
}

// LessonGroup for lessons
type LessonGroup []Lesson

// AfterClone of Lesson
func (lesson Lesson) AfterClone() Lesson {
	lesson.LessonID = lesson.ID
	lesson.SectionID = nil
	lesson.ID = nil
	return lesson
}

// GetPKID .
func (lesson Lesson) GetPKID() *uint64 {
	return lesson.ID
}

// ValidateOnPublish .
func (lesson Lesson) ValidateOnPublish() error {
	if lesson.Status == types.STATUS_MERGED || lesson.Status == types.STATUS_PUBLISHED {
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
	lesson.Status = types.STATUS_PUBLISHED
	return lesson
}

// GroupAfterClone .
func (lessonGroup LessonGroup) GroupAfterClone() LessonGroup {
	lessons := []Lesson{}
	for _, lesson := range lessonGroup {
		lessons = append(lessons, lesson.AfterClone())
	}
	return lessons
}

// GroupValidation need to be invoked via reflection instead of calling it directly
// Deprecated
func (lessonGroup LessonGroup) GroupValidation() error {
	for _, lesson := range lessonGroup {
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
