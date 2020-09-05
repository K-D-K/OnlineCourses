package main

import (
	"OnlineCourses/handler"
	"OnlineCourses/handler/course"
	"OnlineCourses/handler/user"
	"OnlineCourses/middleware"
	"OnlineCourses/migration"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// Declare a new router
	router := mux.NewRouter()
	migration.RunMigration()
	router.HandleFunc("/users", handler.ExecutorWithDB(user.POST)).Methods("POST")

	router.Use(middleware.UserValidation)

	router.HandleFunc("/course/{course_id}", handler.ExecutorWithDB(course.GET)).Methods("GET")
	router.HandleFunc("/courses", handler.ExecutorWithDB(course.GET_ALL)).Methods("GET")

	router.Use(middleware.ValidateAuthorPermissions)

	router.HandleFunc("/users", handler.ExecutorWithDB(user.GET_ALL)).Methods("GET")
	router.HandleFunc("/courses", handler.ExecutorWithDB(course.POST)).Methods("POST")
	http.ListenAndServe(":8001", router)
}
