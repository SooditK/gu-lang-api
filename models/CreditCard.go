package models

import "gorm.io/gorm"

type CreditCard struct {
	gorm.Model
	Number string
	UserID uint
}
