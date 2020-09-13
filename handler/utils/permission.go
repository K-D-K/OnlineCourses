package utils

import (
	"OnlineCourses/controller/relation"
	"OnlineCourses/models"

	"github.com/jinzhu/gorm"
)

// CreatePermission for
func CreatePermission(userIDs []uint64, courseIDs []uint64, db *gorm.DB) {
	relationInstance := relation.INSTANCE(db)
	for _, userID := range userIDs {
		for _, courseID := range courseIDs {
			courseRel := models.CourseRelation{
				UserID:      userID,
				CourseID:    courseID,
				IsCompleted: false,
			}
			relationInstance.UpdatePermission(&courseRel)
		}
	}
}

// CreatePermissionForUsers for Users
func CreatePermissionForUsers(userIDs []uint64, courseID uint64, db *gorm.DB) {
	CreatePermission(userIDs, []uint64{courseID}, db)
}

// CreatePermissionForCourses for courses
func CreatePermissionForCourses(userID uint64, courseIDs []uint64, db *gorm.DB) {
	CreatePermission([]uint64{userID}, courseIDs, db)
}
