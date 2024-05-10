package domain

import "time"

type Address struct {
	AddressID  int    `gorm:"column:address_id;primaryKey;autoIncrement"`
	UserIDFK   int    `gorm:"column:user_id_fk"`
	City       string `gorm:"column:city"`
	Province   string `gorm:"column:province"`
	PostalCode int    `gorm:"column:postal_code"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
