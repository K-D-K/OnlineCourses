package entity

import (
	"OnlineCourses/interfaces"
	"errors"
)

// ValidateEntityOnCreate .
func ValidateEntityOnCreate(entity interfaces.Entity) error {
	return ValidateEntity(entity, true)
}

// ValidateEntitiesOnCreate .
func ValidateEntitiesOnCreate(entities []interfaces.Entity) error {
	return ValidateEntities(entities, true)
}

// ValidateEntityOnUpdate .
func ValidateEntityOnUpdate(entity interfaces.Entity) error {
	return ValidateEntity(entity, false)
}

// ValidateEntitiesOnUpdate .
func ValidateEntitiesOnUpdate(entities []interfaces.Entity) error {
	return ValidateEntities(entities, false)
}

// ValidateEntity .
func ValidateEntity(entity interfaces.Entity, isCreate bool) error {
	if isCreate {
		if entity.GetPKID() != nil {
			return errors.New("PKId found on record create validation")
		}
	} else {
		isCreate = entity.GetPKID() == nil
	}

	entities := entity.GetChildEntities()
	for _, entityGroup := range entities {
		ValidateEntities(entityGroup, isCreate)
	}
	return nil
}

// ValidateEntities .
func ValidateEntities(entities []interfaces.Entity, isCreate bool) error {
	var err error
	for _, entity := range entities {
		err = ValidateEntity(entity, isCreate)
		if err != nil {
			break
		}
	}
	return err
}
