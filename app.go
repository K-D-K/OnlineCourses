package main

import (
	"OnlineCourses/handler"
	"OnlineCourses/handler/course"
	"OnlineCourses/handler/relation"
	"OnlineCourses/handler/user"
	"OnlineCourses/middleware"
	"OnlineCourses/migration"
	"net/http"

	"github.com/go-chi/chi"
)

func main() {
	// Declare a new router
	router := chi.NewRouter()
	migration.RunMigration()
	router.Post("/users", handler.ExecutorWithDB(user.POST))
	router.Group(func(router chi.Router) {
		router.Use(middleware.UserValidation)
		router.Post("/course/{course_id}/enroll", handler.ExecutorWithDB(relation.ENROLL))
		router.Get("/course/{course_id}", handler.ExecutorWithDB(course.GET))
		router.Get("/courses", handler.ExecutorWithDB(course.GET_ALL))

		router.Post("/lesson/{course_id}/{section_id}/{lesson_id}", handler.ExecutorWithDB(relation.COMPLETE))
		router.Group(func(router chi.Router) {
			router.Use(middleware.ValidateAuthorPermissions)

			router.Get("/users", handler.ExecutorWithDB(user.GET_ALL))
			router.Post("/courses", handler.ExecutorWithDB(course.POST))
			router.Put("/courses", handler.ExecutorWithDB(course.PUT))
			router.Post("/course/{course_id}/clone", handler.ExecutorWithDB(course.CLONE))
			router.Put("/course/{course_id}/publish", handler.ExecutorWithDB(course.PUBLISH))

			router.Put("/course/{course_id}/permission", handler.ExecutorWithDB(relation.PERMISSION))
		})
	})
	http.ListenAndServe(":8001", router)
}

/*
create database onlinecourse ;
grant ALL PRIVILEGES on database onlinecourse to kdk ;
*/
