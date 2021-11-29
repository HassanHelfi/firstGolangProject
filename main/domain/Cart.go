package domain

import "github.com/jinzhu/gorm"

type Cart struct {
	gorm.Model
	UserID    uint
	User      User
	ProductID uint
	Product   Product
	Quantity  uint
	Price     uint
}
