package models

import (
	"gorm.io/gorm"
)

type user struct {
	gorm.Model
	Username string `gorm:"unique"`
	Password string
}
