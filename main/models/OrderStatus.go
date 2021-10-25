package models

import "github.com/jinzhu/gorm"

type OrderStatus struct {
	gorm.Model
	Name        string
	Description string `gorm:"type:text"`
}
