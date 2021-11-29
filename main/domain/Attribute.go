package domain

import "github.com/jinzhu/gorm"

type Attribute struct {
	gorm.Model
	Name     string
	Products *[]Product `gorm:"many2many:attribute_product"`
}
