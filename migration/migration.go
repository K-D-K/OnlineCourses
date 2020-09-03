package migration

import (
	"OnlineCourses/datastore"
	"OnlineCourses/models"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// RunMigration on server start
func RunMigration() {
	db := datastore.GetDBConnection()
	defer db.Close()
	db.AutoMigrate(&models.Course{}, &models.Lesson{}, &models.Section{}, &models.User{})
}
