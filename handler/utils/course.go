package utils

import (
	"OnlineCourses/controller/course"
	"OnlineCourses/models"
	"OnlineCourses/utils/error"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/jinzhu/gorm"
)

// GetCourse .
func GetCourse(r *http.Request, db *gorm.DB) models.Course {
	courseID, err := strconv.ParseUint(chi.URLParam(r, "course_id"), 10, 64)
	if err != nil {
		error.ThrowAPIError("Invalid course id")
	}
	return course.INSTANCE(db).GetCourse(courseID)
}

// GetAllLessonIdsFromCourse .
func GetAllLessonIdsFromCourse(course models.Course) []uint64 {
	lessonIds := []uint64{}
	for _, section := range course.Section {
		for _, lesson := range section.Lesson {
			lessonIds = append(lessonIds, *lesson.ID)
		}
	}
	return lessonIds
}
