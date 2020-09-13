package course

import (
	"OnlineCourses/models"
	"OnlineCourses/utils/error"
	"errors"
	"strconv"

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

// GetAllCourses fetch all available courses
func (controller Controller) GetAllCourses() []models.Course {
	courses := []models.Course{}
	err := controller.db.Find(&courses).Error
	if err != nil {
		panic(err)
	}
	return courses
}

// GetCourse fetch specific course details
func (controller Controller) GetCourse(courseID uint64) models.Course {
	course := models.Course{}
	err := controller.db.Preload("Section.Lesson").Preload("Section").First(&course, courseID).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			error.ThrowAPIError("Course not found. Course ID : " + strconv.FormatUint(courseID, 10))
		}
		panic(err)
	}
	return course
}

// GetCourses fetch specific course details
func (controller Controller) GetCourses(courseIDs []uint64) []models.Course {
	courses := []models.Course{}
	err := controller.db.Preload("Section.Lesson").Preload("Section").Where("id in (?)", courseIDs).Find(&courses).Error
	if err != nil {
		panic(err)
	}
	return courses
}

// Create a course
func (controller Controller) Create(course *models.Course) {
	controller.db.Create(course)
}

// Update a course
func (controller Controller) Update(course interface{}) {
	controller.db.Save(course)
}

// Delete the course
func (controller Controller) Delete(courseID uint64) {
	controller.db.Where("id = ?", courseID).Delete(models.Course{})
}
