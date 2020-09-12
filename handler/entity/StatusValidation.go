package entity

import (
	"OnlineCourses/interfaces"
	"OnlineCourses/models/types/status"
	"OnlineCourses/utils/error"
)

// StatusComparator .
type StatusComparator struct {
	Status status.Status
}

// CompareEntityStatus .
func (instance StatusComparator) CompareEntityStatus(entity interfaces.Entity) {
	if entity.GetStatus() != instance.Status {
		// Need to modify it as a JSON
		error.ThrowAPIError("status mismatch found between. Entity Name : " + entity.Name().String() + " . Expected Status : " + instance.Status.String() + " . Status Found : " + entity.GetStatus().String())
	}

	entities := entity.GetChildEntities()
	for _, entityGroup := range entities {
		instance.CompareEntitiesStatus(entityGroup)
	}
}

// CompareEntitiesStatus .
func (instance StatusComparator) CompareEntitiesStatus(entities []interfaces.Entity) {
	for _, entity := range entities {
		instance.CompareEntityStatus(entity)
	}
}
