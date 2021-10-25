package models

import "github.com/jinzhu/gorm"

type Vendor struct {
	gorm.Model
	Name     string
	Bio      string     `gorm:"type:text"`
	Products *[]Product `gorm:"many2many:product_vendor;"`
}