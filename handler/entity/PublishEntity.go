package entity

import (
	"OnlineCourses/interfaces"
	"OnlineCourses/models/types/status"
)

// PublishEntity : to publish a entity
func PublishEntity(entity interfaces.Entity) {
	parentID := entity.GetParentID()
	if parentID != nil {
		entity.SetPKID(parentID)
	}
	entity.SetParentID(nil)
	entity.UpdateRelationID(nil)
	entity.SetStatus(status.STATUS_PUBLISHED)

	entityMap := entity.GetChildEntities()
	for entityName, entityGroup := range entityMap {
		PublishEntities(entityGroup[:])
		entityMap[entityName] = entityGroup
	}
	entity.SetChildEntities(entityMap)
}

// PublishEntities : To clone list of entities
func PublishEntities(entities []interfaces.Entity) {
	for _, entity := range entities {
		PublishEntity(entity)
	}
}
