package entity

import "time"

type User struct {
	ID        uint64    `gorm:"primary_key:auto_increment" json:"id"`
	Email     string    `gorm:"uniqueIndex;type:varchar(50)" json:"email"`
	Password  string    `gorm:"->;<-;not null" json:"-"`
	Token     string    `gorm:"-" json:"token,omitempty"`
	CreatedAt time.Time `json:"created_At"`
	UpdatedAt time.Time `json:"updated_At"`
}
