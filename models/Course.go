package models

import (
	"OnlineCourses/models/types"
	"errors"
	"strconv"

	"github.com/jinzhu/copier"
)

// Course Meta data
type Course struct {
	InfoMeta
	CourseID *uint64      `gorm:"column:parent_id" json:"parent_id,string" sql:"default:null"`
	Section  SectionGroup `json:"sections"` // gorm:"association_autoupdate:false" Commented temporarily
}

// AfterClone for Course
func (course Course) AfterClone() Course {
	course.CourseID = course.ID
	course.ID = nil
	course.Section = course.Section.GroupAfterClone()
	return course
}

// GetPKID for that record
func (course Course) GetPKID() uint64 {
	return *course.ID
}

// ValidateOnPublish for course
func (course Course) ValidateOnPublish() error {
	if course.Status == types.STATUS_MERGED || course.Status == types.STATUS_PUBLISHED {
		return course.Section.GroupValidation()
	}
	if course.CourseID != nil {
		return errors.New("Kindly merge the Course " + strconv.FormatUint(*course.ID, 10) + " with " + strconv.FormatUint(*course.CourseID, 10))
	}
	// TODO : Fix me
	return errors.New("Kindly publish/Save the Course")
}

// Clone for Course
func (course Course) Clone() Course {
	dest := Course{}
	copier.Copy(&dest, &course)
	dest = dest.AfterClone()
	return dest
}

func (course Course) BeforePublish() Course {
	if course.CourseID != nil {
		course.ID = course.CourseID
		course.CourseID = nil
	}
	course.Status = types.STATUS_PUBLISHED
	course.Section = course.Section.GroupBeforePublish()
	return course
}
