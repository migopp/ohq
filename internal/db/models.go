package db

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"unique;not null"`
}
