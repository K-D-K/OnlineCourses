package course

import (
	"OnlineCourses/models"

	"github.com/jinzhu/gorm"
)

// Controller for course
type Controller struct {
	db *gorm.DB
}

// INSTANCE : Get Course Controller instance
func INSTANCE(db *gorm.DB) *Controller {
	return &Controller{db}
}

// GetCourses fetch all available courses
func (controller Controller) GetCourses() []models.Course {
	courses := []models.Course{}
	err := controller.db.Find(&courses).Error
	if err != nil {
		panic(err)
	}
	return courses
}

// GetCourse fetch specific course details
func (controller Controller) GetCourse(courceID uint64) models.Course {
	course := models.Course{}
	err := controller.db.Preload("Section.Lesson").Preload("Section").First(&course, courceID).Error
	if err != nil {
		panic(err)
	}
	return course
}

// Create a course
func (controller Controller) Create(course interface{}) {
	controller.db.Create(course)
}
