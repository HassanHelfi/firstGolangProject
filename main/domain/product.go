package domain

import "github.com/jinzhu/gorm"

type Product struct {
	//gorm.Model
	Name        string
	Description string `gorm:"type:text"`
	//Images      pq.StringArray `gorm:"type:text[]"`
	//Attributes  *[]Attribute   `gorm:"many2many:attribute_product"`
	//Categories  *[]Category    `gorm:"many2many:category_product;"`
	//Vendors     *[]Vendor      `gorm:"many2many:product_vendor;"`
}

type ProductRepository interface {
	FindAllProduct() (*gorm.DB, error)
	OneProduct(id string) (*gorm.DB, error)
	StoreProduct(name string, disc string) (*gorm.DB, error)
}
