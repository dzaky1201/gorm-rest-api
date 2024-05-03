package domain

import (
	"gorm.io/gorm"
)

type Address struct {
	gorm.Model
	AddressID  int    `gorm:"column:address_id;primaryKey;autoIncrement"`
	UserIDFK   int    `gorm:"column:user_id_fk"`
	City       string `gorm:"column:city"`
	Province   string `gorm:"column:province"`
	PostalCode int    `gorm:"column:postal_code"`
}
