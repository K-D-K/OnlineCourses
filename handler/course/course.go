package course

import (
	"OnlineCourses/controller/course"
	"OnlineCourses/handler"
	"OnlineCourses/handler/entity"
	"OnlineCourses/models"
	"OnlineCourses/models/types/status"
	"OnlineCourses/utils/error"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/jinzhu/gorm"
)

// GET requested project
func GET(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	course := get(r, db)
	byteArr, _ := json.Marshal(course)
	handler.RespondwithJSON(w, http.StatusOK, byteArr)
}

func get(r *http.Request, db *gorm.DB) models.Course {
	courseID, err := strconv.ParseUint(chi.URLParam(r, "course_id"), 10, 64)
	if err != nil {
		error.ThrowAPIError("Invalid course id")
	}

	return course.INSTANCE(db).GetCourse(courseID)
}

// GET_ALL courses
func GET_ALL(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	courseInstance := course.INSTANCE(db)
	courses := courseInstance.GetAllCourses()

	byteArr, _ := json.Marshal(courses)
	handler.RespondwithJSON(w, http.StatusOK, byteArr)
}

// POST course
func POST(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	decoder := json.NewDecoder(r.Body)
	decoder.Token()
	courses := []models.Course{}
	for decoder.More() {
		courseModal := models.Course{}
		err := decoder.Decode(&courseModal)
		if err != nil {
			error.ThrowAPIError(err.Error())
		}
		courses = append(courses, courseModal)
	}

	createAll(courses[:], db)
	byteArr, _ := json.Marshal(courses)
	handler.RespondwithJSON(w, http.StatusOK, byteArr)
}

func createAll(courses []models.Course, db *gorm.DB) {
	courseInstance := course.INSTANCE(db)
	for index, course := range courses {
		entity.ValidateEntityOnCreate(&course)
		if course.Status == status.STATUS_PUBLISHED {
			statusComparatorInstance := entity.StatusComparator{
				Status: status.STATUS_PUBLISHED,
			}
			statusComparatorInstance.CompareEntityStatus(&course)
		} else {
			maxStatusValidator := entity.MaxStatusValidation{
				Status: status.STATUS_SAVED,
			}
			maxStatusValidator.CompareEntityStatus(&course)
		}

		// GORM not supported Bulk Insert or Update.
		// Need to handle associations in our end.
		courseInstance.Create(&course)
		courses[index] = course
	}
}

// CLONE course
// Add validation to clone . It should allow only published Course
func CLONE(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	courseInfo := get(r, db)
	entity.CloneEntity(&courseInfo)
	course.INSTANCE(db).Create(&courseInfo)
	byteArr, _ := json.Marshal(courseInfo)
	handler.RespondwithJSON(w, http.StatusOK, byteArr)
}

// PUT course
// Need to handle partial update fields instead of passing whole course data each time.
// For delete of sections we need to handle we need to add a field like _delete which should not persist in DB
func PUT(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	decoder := json.NewDecoder(r.Body)
	decoder.Token()
	courses := []models.Course{}
	courseInstance := course.INSTANCE(db)
	for decoder.More() {
		courseModal := models.Course{}
		err := decoder.Decode(&courseModal)
		if err != nil {
			error.ThrowAPIError(err.Error())
		}
		// Need to add validation case here
		// Validate prev status and current status and should not allow to publish via put request
		// For publish need to expose seperate API
		// Need to add method for update with section removal validation
		courseInstance.Update(&courseModal)
		courses = append(courses, courseModal)
	}

	byteArr, _ := json.Marshal(courses)
	handler.RespondwithJSON(w, http.StatusOK, byteArr)
}

// PUBLISH the Course
func PUBLISH(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	courseInfo := get(r, db)
	var courseID *uint64

	var statusToValidate status.Status
	if courseInfo.GetParentID() == nil {
		statusToValidate = status.STATUS_SAVED
	} else {
		courseID = courseInfo.GetPKID()
		statusToValidate = status.STATUS_MERGED
	}
	statusComparatorInstance := entity.StatusComparator{
		Status: statusToValidate,
	}

	statusComparatorInstance.CompareEntityStatus(&courseInfo)
	entity.PublishEntity(&courseInfo)

	courseInstance := course.INSTANCE(db)
	courseInstance.Update(&courseInfo)
	if courseID != nil {
		courseInstance.Delete(*courseID)
	}

	byteArr, _ := json.Marshal(courseInfo)
	handler.RespondwithJSON(w, http.StatusOK, byteArr)
}
