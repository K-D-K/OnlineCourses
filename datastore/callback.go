package datastore

import (
	"OnlineCourses/models"
	"OnlineCourses/utils/constants"

	"github.com/jinzhu/gorm"
)

func getUser(scope *gorm.Scope) (models.User, bool) {
	user, hasUser := scope.DB().Get(constants.GORMInstanceUserKey)
	if hasUser {
		user, isUser := user.(models.User)
		return user, isUser
	}
	return models.User{}, false
}

func assignCreatedBy(scope *gorm.Scope) {
	user, ok := getUser(scope)
	if ok {
		scope.SetColumn("CreatedBy", user)
		scope.SetColumn("UpdatedBy", user)
	}
}

func assignUpdatedBy(scope *gorm.Scope) {
	user, ok := getUser(scope)
	if ok {
		scope.SetColumn("UpdatedBy", user)
	}
}

// RegisterCallback to insert/update created-by and updated-by columns
func RegisterCallback(db *gorm.DB) {
	callback := db.Callback()
	callback.Create().After("gorm:before_create").Register("audited:assign_created_by", assignCreatedBy)
	callback.Update().After("gorm:before_update").Register("audited:assign_updated_by", assignUpdatedBy)
}
