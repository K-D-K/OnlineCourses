package entity

// Entity enum
type Entity string

const (
	COURSE  Entity = "course"
	SECTION        = "section"
	LESSON         = "lesson"
)

func (entity Entity) String() string {
	return string(entity)
}
