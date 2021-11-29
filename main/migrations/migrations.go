package migrations

import (
	"crud_echo/configs"
	"crud_echo/domain"
)

func Migrate() {
	var DBCon = configs.CreateCon()
	DBCon.AutoMigrate(
		&domain.User{},
		&domain.Role{},
		&domain.Permission{},
		&domain.Category{},
		&domain.Product{},
		&domain.Attribute{},
		&domain.Vendor{},
		&domain.AttributeProduct{},
		&domain.ProductVendor{},
		&domain.Comment{},
		&domain.Order{},
		&domain.Transaction{},
		&domain.OrderStatus{},
		&domain.Shipper{},
		&domain.Zipcode{},
		&domain.Cart{},
	)
}
