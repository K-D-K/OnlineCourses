package models

// Course Meta data
type Course struct {
	InfoMeta
	Section []Section `gorm:"association_autoupdate:false" json:"sections"`
}
