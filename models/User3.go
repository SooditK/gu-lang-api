package models

import "gorm.io/gorm"

type User3 struct {
	gorm.Model
	Name      string
	Languages []Language `gorm:"many2many:user_languages;"`
}
