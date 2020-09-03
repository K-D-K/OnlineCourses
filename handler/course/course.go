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
	params := mux.Vars(r)
	courseID, err := strconv.ParseUint(params["course_id"], 10, 64)
	if err != nil {
		error.ThrowAPIError("Invalid course id")
	}

	courseInstance := course.INSTANCE(db)
	course := courseInstance.GetCourse(courseID)

	byteArr, _ := json.Marshal(course)

	handler.RespondwithJSON(w, http.StatusOK, byteArr)
}

// GET_ALL courses
func GET_ALL(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	courseInstance := course.INSTANCE(db)
	courses := courseInstance.GetCourses()

	byteArr, _ := json.Marshal(courses)
	handler.RespondwithJSON(w, http.StatusOK, byteArr)
}

// POST course
func POST(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	decoder := json.NewDecoder(r.Body)
	decoder.Token()
	courses := []models.Course{}
	courseModal := models.Course{}
	courseInstance := course.INSTANCE(db)
	for decoder.More() {
		decoder.Decode(&courseModal)
		courseInstance.Create(&courseModal)
		courses = append(courses, courseModal)
	}

	byteArr, _ := json.Marshal(courses)
	handler.RespondwithJSON(w, http.StatusOK, byteArr)
}
