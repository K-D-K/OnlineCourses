package relation

import (
	"OnlineCourses/datastore"
	"OnlineCourses/models"

	"github.com/jinzhu/gorm"
)

type Controller struct {
	db *gorm.DB
}

// INSTANCE : Get Course Controller instance
func INSTANCE(db *gorm.DB) *Controller {
	return &Controller{db}
}

// Enroll .
func (controller Controller) Enroll(courseID uint64) {
	courseRel := controller.getCourseRelationObj(courseID, false)
	controller.db.Create(&courseRel)
}

// UpdatePermission .
func (controller Controller) UpdatePermission(courseRel interface{}) {
	controller.db.Create(courseRel)
}

// DeletePermission .
func (controller Controller) DeletePermission(courseID uint64, userIDs []uint64) {
	controller.db.Where("user_id IN (?) and course_id = ?", userIDs, courseID).Delete(models.CourseRelation{})
}

func (controller Controller) getCourseRelationObj(courseID uint64, status bool) models.CourseRelation {
	user, _ := datastore.GetUser(controller.db)
	return models.CourseRelation{
		UserID:      *user.ID,
		CourseID:    courseID,
		IsCompleted: false,
	}
}

// CompleteLesson .
func (controller Controller) CompleteLesson(courseID uint64, lessonID uint64) {
	user, _ := datastore.GetUser(controller.db)
	courseTracker := models.CourseTracker{
		UserID:   *user.ID,
		CourseID: courseID,
		LessonID: lessonID,
	}
	controller.db.Create(&courseTracker)
}

// GetCompletedLessons for a course
func (controller Controller) GetCompletedLessons(courseID uint64) []models.CourseTracker {
	user, _ := datastore.GetUser(controller.db)
	courseCompletions := []models.CourseTracker{}
	controller.db.Where("user_id = ? and course_id = ?", *user.ID, courseID).Find(&courseCompletions)
	return courseCompletions
}

// UpdateCourseStatus .
func (controller Controller) UpdateCourseStatus(courseID uint64, status bool) {
	courseRel := controller.getCourseRelationObj(courseID, status)
	controller.db.Save(&courseRel)
}

// GetCourseStatus .
func (controller Controller) GetCourseStatus(courseID uint64) (models.CourseRelation, error) {
	user, _ := datastore.GetUser(controller.db)
	courseRel := models.CourseRelation{}
	err := controller.db.Where("user_id = ? and course_id = ?", *user.ID, courseID).First(&courseRel).Error
	return courseRel, err
}

// GetCoursesPermission .
func (controller Controller) GetCoursesPermission(courseIDs []uint64) []models.CourseRelation {
	user, _ := datastore.GetUser(controller.db)
	courseRel := []models.CourseRelation{}
	controller.db.Where("user_id = ? and course_id in (?)", *user.ID, courseIDs).Find(&courseRel)
	return courseRel
}
