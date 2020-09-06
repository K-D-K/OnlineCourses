package models

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
func (lesson Lesson) GetPKID() uint64 {
	return *lesson.LessonID
}

// ValidateOnPublish .
func (lesson Lesson) ValidateOnPublish() error {
	return nil
}

// GroupAfterClone .
func (lessonGroup LessonGroup) GroupAfterClone() LessonGroup {
	lessons := []Lesson{}
	for _, lesson := range lessonGroup {
		(&lesson).AfterClone()
		lessons = append(lessons, lesson)
	}
	return lessons
}
