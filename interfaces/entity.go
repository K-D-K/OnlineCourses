package interfaces

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
	GetPKID() uint64
	/*
		Need to add implementation by iterating fields.
		Ex : Sections inside course need to called programatically with out invoking it explicity
	*/
	AfterClone() Entity
	/*
		Same like AfterClone need to call internal entities ValidateOnPublish to avoid manual handling
	*/
	ValidateOnPublish() error
}

/*
	EntityGroup is defined to handle generic handling for group of entities
*/
type EntityGroup interface {
	BulkAfterClone() EntityGroup
}
