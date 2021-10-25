package models

import (
	"github.com/jinzhu/gorm"
	"github.com/lib/pq"
)

type Product struct {
	gorm.Model
	Name        string
	Description string         `gorm:"type:text"`
	Images      pq.StringArray `gorm:"type:text[]"`
	Attributes  *[]Attribute   `gorm:"many2many:attribute_product"`
	Categories  *[]Category    `gorm:"many2many:category_product;"`
	Vendors     *[]Vendor      `gorm:"many2many:product_vendor;"`
}
