package entity

import (
	"OnlineCourses/interfaces"
	"OnlineCourses/models/types/status"
)

// CloneEntity : to clone a entity
func CloneEntity(entity interfaces.Entity) {
	pkID := entity.GetPKID()
	entity.SetParentID(pkID)
	entity.SetPKID(nil)
	entity.UpdateRelationID(nil)
	entity.SetStatus(status.STATUS_MERGED)

	entityMap := entity.GetChildEntities()
	for entityName, entityGroup := range entityMap {
		CloneEntities(entityGroup[:])
		entityMap[entityName] = entityGroup
	}
	entity.SetChildEntities(entityMap)
}

// CloneEntities : To clone list of entities
func CloneEntities(entities []interfaces.Entity) {
	for _, entity := range entities {
		CloneEntity(entity)
	}
}
