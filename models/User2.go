package models

import "gorm.io/gorm"

type User2 struct {
	gorm.Model
	Name         string
	CreditCards2 []CreditCard2 `gorm:"foreignKey:UserID"`
}
