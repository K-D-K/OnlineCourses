package course

import (
	"OnlineCourses/controller/course"
	entityController "OnlineCourses/controller/entity"
	"OnlineCourses/datastore"
	"OnlineCourses/handler"
	"OnlineCourses/handler/entity"
	handlerUtils "OnlineCourses/handler/utils"
	"OnlineCourses/models"
	"OnlineCourses/models/types/status"
	"OnlineCourses/utils"
	"OnlineCourses/utils/error"
	"encoding/json"
	"net/http"

	"github.com/jinzhu/gorm"
)

// GET requested project
func GET(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	course := handlerUtils.GetCourse(r, db)
	handlerUtils.CheckCoursePermission(*course.ID, db)
	byteArr, _ := json.Marshal(course)
	handler.RespondwithJSON(w, http.StatusOK, byteArr)
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
	courseIDs := make([]uint64, len(courses))
	for index, course := range courses {
		entity.ValidateEntityOnCreate(&course)
		var statusToValidate status.Status
		if course.Status == status.STATUS_PUBLISHED {
			statusToValidate = status.STATUS_PUBLISHED
		} else {
			statusToValidate = status.STATUS_SAVED
		}
		maxStatusValidator := entity.MaxStatusValidation{
			Status: statusToValidate,
		}
		maxStatusValidator.CompareEntityStatus(&course)

		// GORM not supported Bulk Insert or Update.
		// Need to handle associations in our end.
		courseInstance.Create(&course)
		courses[index] = course
		courseIDs[index] = *course.ID
	}
	user, _ := datastore.GetUser(db)
	handlerUtils.CreatePermissionForCourses(*user.ID, courseIDs, db)
}

// CLONE course
// Add validation to clone . It should allow only published Course
func CLONE(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	courseInfo := handlerUtils.GetCourse(r, db)
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
	coursesToCreate := []models.Course{}
	coursesToUpdate := []models.Course{}
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

		if courseModal.GetPKID() == nil {
			coursesToCreate = append(coursesToCreate, courseModal)
		} else {
			coursesToUpdate = append(coursesToUpdate, courseModal)
		}
	}

	if len(coursesToCreate) > 0 {
		createAll(coursesToCreate[:], db)
	}
	if len(coursesToUpdate) > 0 {
		updateAll(coursesToUpdate[:], db)
	}
	courses := append(coursesToCreate, coursesToUpdate...)
	byteArr, _ := json.Marshal(courses)
	handler.RespondwithJSON(w, http.StatusOK, byteArr)
}

func updateAll(courses []models.Course, db *gorm.DB) {
	updatedCourseEntityArr := models.ConvertCourseIntoEntityArr(courses)
	pkIDs := utils.GetPKIDs(updatedCourseEntityArr)
	if len(pkIDs) == 0 {
		return
	}
	handlerUtils.CheckCoursesPermission(pkIDs, db)
	deleteEntitiesMap := entity.CollectDeletedDataForEntities(updatedCourseEntityArr)

	courseInstance := course.INSTANCE(db)
	oldCoursesEntityArr := models.ConvertCourseIntoEntityArr(courseInstance.GetCourses(pkIDs))
	entity.CompareAndUpdateValueForEntities(updatedCourseEntityArr[:], oldCoursesEntityArr)
	entity.CompareEntityStatus(updatedCourseEntityArr)
	courses = models.ConvertEntityToCourseArr(updatedCourseEntityArr)
	for index, course := range courses {
		// GORM not supported Bulk Insert or Update.
		// Need to handle associations in our end.
		courseInstance.Update(&course)
		courses[index] = course
	}

	entityController.DeleteEntities(deleteEntitiesMap, db)
}

// PUBLISH the Course
func PUBLISH(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	courseInfo := handlerUtils.GetCourse(r, db)
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
	// Need to delete missing entries from clonned course
	entity.PublishEntity(&courseInfo)

	courseInstance := course.INSTANCE(db)
	courseInstance.Update(&courseInfo)
	if courseID != nil {
		courseInstance.Delete(*courseID)

		// TODO : We need to check existing lessons if new one is added.
		// we need to update relation table also as incompleted lessons for all user who already enrolled.
		// This one need to be done in scheduler as huge data may present.
	}

	byteArr, _ := json.Marshal(courseInfo)
	handler.RespondwithJSON(w, http.StatusOK, byteArr)
}
