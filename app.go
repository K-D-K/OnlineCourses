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

	// Need to add route group to avoid middleware validation
	router.Use(middleware.UserValidation)

	router.HandleFunc("/course/{course_id}", handler.ExecutorWithDB(course.GET)).Methods("GET")
	router.HandleFunc("/courses", handler.ExecutorWithDB(course.GET_ALL)).Methods("GET")

	// Need to add route group to avoid middleware validation
	router.Use(middleware.ValidateAuthorPermissions)

	router.HandleFunc("/users", handler.ExecutorWithDB(user.GET_ALL)).Methods("GET")
	router.HandleFunc("/courses", handler.ExecutorWithDB(course.POST)).Methods("POST")
	router.HandleFunc("/courses", handler.ExecutorWithDB(course.PUT)).Methods("PUT")
	router.HandleFunc("/course/{course_id}/clone", handler.ExecutorWithDB(course.CLONE)).Methods("POST")
	router.HandleFunc("/course/{course_id}/publish", handler.ExecutorWithDB(course.PUBLISH)).Methods("PUT")
	http.ListenAndServe(":8001", router)
}

/*
create database onlinecourse ;
grant ALL PRIVILEGES on database onlinecourse to kdk ;
*/
