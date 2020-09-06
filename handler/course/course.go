package course

import (
	"OnlineCourses/controller/course"
	"OnlineCourses/handler"
	"OnlineCourses/models"
	"OnlineCourses/utils/error"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

// GET requested project
func GET(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	course := get(r, db)
	byteArr, _ := json.Marshal(course)
	handler.RespondwithJSON(w, http.StatusOK, byteArr)
}

func get(r *http.Request, db *gorm.DB) models.Course {
	params := mux.Vars(r)
	courseID, err := strconv.ParseUint(params["course_id"], 10, 64)
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
	courseInstance := course.INSTANCE(db)
	for decoder.More() {
		courseModal := models.Course{}
		err := decoder.Decode(&courseModal)
		if err != nil {
			error.ThrowAPIError(err.Error())
		}
		courseInstance.Create(&courseModal)
		courses = append(courses, courseModal)
	}

	byteArr, _ := json.Marshal(courses)
	handler.RespondwithJSON(w, http.StatusOK, byteArr)
}

// CLONE course
// Add validation to clone . It should allow only published Course
func CLONE(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	courseInfo := get(r, db)
	clonedCourse := courseInfo.Clone()
	course.INSTANCE(db).Create(&clonedCourse)
	byteArr, _ := json.Marshal(clonedCourse)
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
		courseInstance.Update(&courseModal)
		courses = append(courses, courseModal)
	}

	byteArr, _ := json.Marshal(courses)
	handler.RespondwithJSON(w, http.StatusOK, byteArr)
}
