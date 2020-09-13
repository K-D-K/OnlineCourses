package models

// CourseRelation .
type CourseRelation struct {
	UserID      uint64 `gorm:"primaryKey;autoIncrement:false"`
	CourseID    uint64 `gorm:"primaryKey;autoIncrement:false"`
	IsCompleted bool
}

// CourseTracker .
type CourseTracker struct {
	UserID   uint64 `gorm:"primaryKey;autoIncrement:false"`
	CourseID uint64 `gorm:"primaryKey;autoIncrement:false"`
	LessonID uint64 `gorm:"primaryKey;autoIncrement:false"`
}
