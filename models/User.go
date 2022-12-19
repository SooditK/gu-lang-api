package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name       string
	CreditCard CreditCard `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
