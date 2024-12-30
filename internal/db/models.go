package db

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"unique;not null" yaml:"username"`
	Password string `gorm:"not null" yaml:"password"`
}
