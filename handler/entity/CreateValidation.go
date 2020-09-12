package entity

import (
	"OnlineCourses/interfaces"
	"OnlineCourses/utils/error"
)

// ValidateEntityOnCreate .
func ValidateEntityOnCreate(entity interfaces.Entity) {
	ValidateEntity(entity, true)
}

// ValidateEntitiesOnCreate .
func ValidateEntitiesOnCreate(entities []interfaces.Entity) {
	ValidateEntities(entities, true)
}

// ValidateEntityOnUpdate .
func ValidateEntityOnUpdate(entity interfaces.Entity) {
	ValidateEntity(entity, false)
}

// ValidateEntitiesOnUpdate .
func ValidateEntitiesOnUpdate(entities []interfaces.Entity) {
	ValidateEntities(entities, false)
}

// ValidateEntity .
func ValidateEntity(entity interfaces.Entity, isCreate bool) {
	if isCreate {
		if entity.GetPKID() != nil {
			error.ThrowAPIError("PKId found on record create validation")
		}
	} else {
		isCreate = entity.GetPKID() != nil
	}

	entities := entity.GetChildEntities()
	for _, entityGroup := range entities {
		ValidateEntities(entityGroup, isCreate)
	}
}

// ValidateEntities .
func ValidateEntities(entities []interfaces.Entity, isCreate bool) {
	for _, entity := range entities {
		ValidateEntity(entity, isCreate)
	}
}
