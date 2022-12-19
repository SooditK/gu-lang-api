package models

import "gorm.io/gorm"

type CreditCard2 struct {
	gorm.Model
	Number string
	UserID uint
}
