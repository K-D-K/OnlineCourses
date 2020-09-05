package models

// Section struct
type Section struct {
	InfoMeta
	CourseID  *uint    `gorm:"column:course_id" json:"course_id,string"`
	SectionID *uint    `gorm:"column:parent_id" json:"parent_id,string" sql:"default:null"`
	Lesson    []Lesson `gorm:"association_autoupdate:false" json:"lessons"`
}

// AfterClone of section
func (section *Section) AfterClone() {
	section.SectionID = section.ID
	section.ID = nil
	section.CourseID = nil
	lessons := []Lesson{}
	for _, lesson := range section.Lesson {
		(&lesson).AfterClone()
		lessons = append(lessons, lesson)
	}
	section.Lesson = lessons
}
