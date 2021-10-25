package models

import "github.com/jinzhu/gorm"

type Transaction struct {
	gorm.Model
	Code    string
	Mode    uint
	Status  uint
	Type    uint
	OrderID int
	Order   Order
}
