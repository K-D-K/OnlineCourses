package user

import (
	"OnlineCourses/models"

	"github.com/jinzhu/gorm"
)

// Controller .
type Controller struct {
	db *gorm.DB
}

// INSTANCE : to create new instance of controller
func INSTANCE(db *gorm.DB) *Controller {
	return &Controller{db}
}

// GetUsers fetch all users
func (controller Controller) GetUsers() []models.User {
	users := []models.User{}
	err := controller.db.Find(&users).Error
	if err != nil {
		panic(err)
	}
	return users
}

// Create a user
func (controller Controller) Create(user interface{}) {
	controller.db.Create(user)
}
