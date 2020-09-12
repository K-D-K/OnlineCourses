package interfaces

import (
	"OnlineCourses/models/types/entity"
	"OnlineCourses/models/types/status"
)

// NOTE : Class is under construction :p. Need to handle all cases

/*
	Entity is defined to handled generic cases in models like course, section and lessons.
	Each have parent child relationship but at some point all underlying in same point
*/
type Entity interface {
	/*
		Fetch PK id so that we group values with the help of Interface.
		Delete missing PK'ids on save.
		which can be done by adding an method in EntityGroup Interface
	*/
	GetPKID() *uint64

	/*
		Same like AfterClone need to call internal entities ValidateOnPublish to avoid manual handling
	*/
	ValidateOnPublish() error

	/*
		Get Child Entities
	*/
	GetChildEntities() map[string][]Entity

	/*
		Set Child Entities
	*/
	SetChildEntities(entities map[string][]Entity)

	/*
		Update parent Id
	*/
	UpdateParentID(parentID *uint64)

	/*
		Update Relation ID
	*/
	UpdateRelationID(relID *uint64)

	/*
		Reset Id
	*/
	ResetPKID()

	/*
		Set Status
	*/
	SetStatus(status status.Status)

	/*
		Get Status
	*/
	GetStatus() status.Status

	/*
		Name of an Entity
	*/
	Name() entity.Entity
}

/*
	EntityGroup is defined to handle generic handling for group of entities
*/
type EntityGroup interface {
	GroupAfterClone() EntityGroup
	GroupValidation() error
}

// EntityGroup .
// type EntityGroup []Entity
