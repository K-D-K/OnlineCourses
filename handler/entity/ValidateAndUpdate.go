package entity

import (
	"OnlineCourses/interfaces"
	"OnlineCourses/models/types/status"
	"OnlineCourses/utils"
	"OnlineCourses/utils/error"
	"strconv"
)

// CompareAndUpdateValueForEntity .
func CompareAndUpdateValueForEntity(updatedEntity interfaces.Entity, oldEntity interfaces.Entity) {
	if oldEntity.GetStatus() == status.STATUS_PUBLISHED {
		error.ThrowAPIError("we can't update Published " + updatedEntity.Name().String())
	}

	// we won't expose parent id to JSON.
	updatedEntity.SetParentID(oldEntity.GetParentID())

	if (status.GetIndexForStatus(oldEntity.GetStatus()) - status.GetIndexForStatus(updatedEntity.GetStatus())) > 1 {
		error.ThrowAPIError("we can't update status from " + oldEntity.GetStatus().String() + " to " + updatedEntity.GetStatus().String() + " for entity " + updatedEntity.Name().String())
	}

	updatedEntityMap := updatedEntity.GetChildEntities()
	oldEntitiesMap := oldEntity.GetChildEntities()
	for entityName, entityGroup := range updatedEntityMap {
		CompareAndUpdateValueForEntities(entityGroup[:], oldEntitiesMap[entityName])
		updatedEntityMap[entityName] = entityGroup
	}
	updatedEntity.SetChildEntities(updatedEntityMap)
}

// CompareAndUpdateValueForEntities .
func CompareAndUpdateValueForEntities(updatedEntities []interfaces.Entity, oldEntities []interfaces.Entity) {
	pkIDVsEntityMap := utils.GetPKIDVsEntityMap(oldEntities)
	for index, updatedEntity := range updatedEntities {
		if updatedEntity.GetPKID() != nil {
			oldEntity, ok := pkIDVsEntityMap[*updatedEntity.GetPKID()]
			if !ok {
				error.ThrowAPIError("Record Not found for " + updatedEntity.Name().String() + " . Record Id : " + strconv.FormatUint(*updatedEntity.GetPKID(), 10))
			}
			CompareAndUpdateValueForEntity(updatedEntity, oldEntity)
			updatedEntities[index] = updatedEntity
		}
	}
}
