package entity

import (
	"OnlineCourses/interfaces"
	"OnlineCourses/models/types/status"
	"OnlineCourses/utils/error"
)

// MaxStatusValidation .
type MaxStatusValidation struct {
	Status status.Status
}

// CompareEntityStatus .
func (instance MaxStatusValidation) CompareEntityStatus(entity interfaces.Entity) {
	if status.GetIndexForStatus(entity.GetStatus()) < status.GetIndexForStatus(instance.Status) {
		// Need to modify it as a JSON
		error.ThrowAPIError("status is higher than expected value. Entity Name : " + entity.Name().String() + " . Expected Status : " + instance.Status.String() + " . Status Found : " + entity.GetStatus().String())
	}

	entities := entity.GetChildEntities()
	for _, entityGroup := range entities {
		instance.CompareEntitiesStatus(entityGroup)
	}
}

// CompareEntitiesStatus .
func (instance MaxStatusValidation) CompareEntitiesStatus(entities []interfaces.Entity) {
	for _, entity := range entities {
		instance.CompareEntityStatus(entity)
	}
}
