package relation

import (
	"OnlineCourses/controller/relation"
	"OnlineCourses/handler"
	handlerUtils "OnlineCourses/handler/utils"
	"OnlineCourses/models/types/status"
	"OnlineCourses/utils"
	"OnlineCourses/utils/error"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/jinzhu/gorm"
)

// ENROLL a course.
func ENROLL(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	course := handlerUtils.GetCourse(r, db)
	if course.Status != status.STATUS_PUBLISHED {
		error.ThrowAPIError("Published courses only allowed to enroll")
	}

	relationInstance := relation.INSTANCE(db)
	relationInstance.Enroll(*course.ID)

	byteArr, _ := json.Marshal(map[string]string{"message": "Succesfully enrolled"})
	handler.RespondwithJSON(w, http.StatusOK, byteArr)
}

// COMPLETE  a lesson.
func COMPLETE(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	// Getting whole course details is not advisable
	// Need to create a join and fetch course and lesson details and status alone
	// Hence we using ORM so need to find required utils to fetch data
	course := handlerUtils.GetCourse(r, db)
	if course.Status != status.STATUS_PUBLISHED {
		error.ThrowAPIError("Lessons can be completed only for published courses")
	}

	lessonID, err := strconv.ParseUint(chi.URLParam(r, "lesson_id"), 10, 64)
	if err != nil {
		error.ThrowAPIError("Invalid Lesson id")
	}

	relationInstance := relation.INSTANCE(db)
	courseEnrollStatus, err := relationInstance.GetCourseStatus(*course.ID)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		error.ThrowAPIError("User not enrolled to the course")
	}

	if courseEnrollStatus.IsCompleted {
		byteArr, _ := json.Marshal(map[string]string{"message": "Course already completed"})
		handler.RespondwithJSON(w, http.StatusOK, byteArr)
		return
	}

	completedLessons := relationInstance.GetCompletedLessons(*course.ID)
	for _, completedLesson := range completedLessons {
		if lessonID == completedLesson.LessonID {
			byteArr, _ := json.Marshal(map[string]string{"message": "Lesson already completed"})
			handler.RespondwithJSON(w, http.StatusOK, byteArr)
			return
		}
	}

	relationInstance.CompleteLesson(*course.ID, lessonID)
	lessonsIDs := handlerUtils.GetAllLessonIdsFromCourse(course)
	if len(lessonsIDs) == len(completedLessons)+1 {
		relationInstance.UpdateCourseStatus(*course.ID, true)
	}

	byteArr, _ := json.Marshal(map[string]string{"message": "status updated successfully"})
	handler.RespondwithJSON(w, http.StatusOK, byteArr)
}

// PERMISSION .
func PERMISSION(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	courseID, err := strconv.ParseUint(chi.URLParam(r, "course_id"), 10, 64)
	if err != nil {
		error.ThrowAPIError("Invalid Course id")
	}
	var body map[string][]interface{}
	bodyByte, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(bodyByte, &body)

	userIDsToAdd := body["add"]
	if len(userIDsToAdd) > 0 {
		handlerUtils.CreatePermissionForUsers(utils.ConvertToUintArray(userIDsToAdd), courseID, db)
	}

	userIDsToRemove := body["remove"]
	if len(userIDsToRemove) > 0 {
		relationInstance := relation.INSTANCE(db)
		relationInstance.DeletePermission(courseID, utils.ConvertToUintArray(userIDsToRemove))
	}

	byteArr, _ := json.Marshal(map[string]string{"message": "Permissions updated successfully"})
	handler.RespondwithJSON(w, http.StatusOK, byteArr)
}
