package utils

import (
	"OnlineCourses/controller/relation"
	"OnlineCourses/models"
	"OnlineCourses/utils/error"
	"strconv"

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

// CheckCoursePermission .
func CheckCoursePermission(courseID uint64, db *gorm.DB) {
	CheckCoursesPermission([]uint64{courseID}, db)
}

// CheckCoursesPermission .
func CheckCoursesPermission(courseIDs []uint64, db *gorm.DB) {
	relationInstance := relation.INSTANCE(db)
	courseRelArr := relationInstance.GetCoursesPermission(courseIDs)

	courseIDsMap := make(map[uint64]bool)
	for _, courseRel := range courseRelArr {
		courseIDsMap[courseRel.CourseID] = true
	}
	for _, courseID := range courseIDs {
		if !courseIDsMap[courseID] {
			error.ThrowAPIError("Permission Denied. Course ID : " + strconv.FormatUint(courseID, 10))
		}
	}
}
