package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID        uint   `gorm:"primaryKey" json:"id"`
	Username  string `gorm:"unique;not null" json:"username"`
	Password  string `gorm:"not null" json:"password"`
	Email     string `gorm:"unique;not null" json:"email"`
	FirstName string `gorm:"not null" json:"first_name"`
	LastName  string `gorm:"not null" json:"last_name"`
}
