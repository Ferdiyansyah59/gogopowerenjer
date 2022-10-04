package entity

import "time"

type Article struct {
	ID        uint64    `gorm:"primary_key:auto_increment" json:"id"`
	Title     string    `gorm:"type:varchar(255)" json:"title"`
	Slug      string    `gorm:"uniqueIndex;type:varchar(255)" json:"slug"`
	MetaDesc  string    `gorm:"type:varchar(255)" json:"meta_desc"`
	Content   string    `gorm:"type:text" json:"content"`
	Cover     string    `gorm:"type:varchar(255)" json:"cover"`
	CreatedAt time.Time `json:"created_At"`
	UpdatedAt time.Time `json:"updated_At"`
}
