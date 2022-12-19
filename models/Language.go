package models

import "gorm.io/gorm"

type Language struct {
	gorm.Model
	Name  string
	Users []User3 `gorm:"many2many:user_languages;"`
}
