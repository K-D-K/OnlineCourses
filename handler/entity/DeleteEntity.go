package entity

import "OnlineCourses/interfaces"

// DeleteEntity for entities
type DeleteEntity struct {
	entityVsIDMap map[string][]uint64
}

// DeleteEntity .
func (instance DeleteEntity) DeleteEntity(entity interfaces.Entity) {
	if entity.IsDeleted() {
		instance.entityVsIDMap[entity.Name().String()] = append(instance.entityVsIDMap[entity.Name().String()], *entity.GetPKID())
	} else {
		entities := entity.GetChildEntities()
		for _, entityGroup := range entities {
			instance.DeleteEntities(entityGroup)
		}
	}

}

// DeleteEntities .
func (instance DeleteEntity) DeleteEntities(entities []interfaces.Entity) {
	for _, entity := range entities {
		instance.DeleteEntity(entity)
	}
}

// CollectDeletedData .
func CollectDeletedData(entity interfaces.Entity) map[string][]uint64 {
	return CollectDeletedDataForEntities([]interfaces.Entity{entity})
}

// CollectDeletedDataForEntities .
// TODO : Collected data need to removed from entities.
func CollectDeletedDataForEntities(entities []interfaces.Entity) map[string][]uint64 {
	instance := DeleteEntity{
		entityVsIDMap: make(map[string][]uint64),
	}
	instance.DeleteEntities(entities)
	return instance.entityVsIDMap
}
