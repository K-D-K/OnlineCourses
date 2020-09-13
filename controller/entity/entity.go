package entity

import (
	"OnlineCourses/models"

	"github.com/jinzhu/gorm"
)

// DeleteEntities .
func DeleteEntities(entityVsIdsMap map[string][]uint64, db *gorm.DB) {
	for entity, ids := range entityVsIdsMap {
		var model interface{}
		switch entity {
		case "course":
			model = models.Course{}
			break
		case "section":
			model = models.Section{}
			break
		case "lesson":
			model = models.Lesson{}
			break
		}
		if model != nil && len(ids) > 0 {
			db.Where("id IN (?)", ids).Delete(model)
		}
	}
}
