package domain

import (
	"crud_echo/configs"
	"fmt"
	"github.com/jinzhu/gorm"
)

type ProductRepositoryDB struct {
	client *gorm.DB
}

func AllProductRepositoryDB() ProductRepositoryDB {
	client := configs.CreateCon()
	return ProductRepositoryDB{client}
}

var product []Product

func (p ProductRepositoryDB) FindAllProduct() (*gorm.DB, error) {
	product := p.client.Find(&product)

	return product, nil
}

func (p ProductRepositoryDB) StoreProduct(name string, disc string) (*gorm.DB, error) {
	prod := p.client.Create(&Product{
		Name:        name,
		Description: disc,
	})

	return prod, nil
}

func (p ProductRepositoryDB) OneProduct(id string) (*gorm.DB, error) {
	fmt.Println(id)
	product := p.client.Find(&product, Product{Name: id})

	return product, nil
}
