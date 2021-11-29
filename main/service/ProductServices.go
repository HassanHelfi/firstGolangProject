package service

import (
	"crud_echo/domain"
	"github.com/jinzhu/gorm"
)

type ProductServices interface {
	GetProduct() (*gorm.DB, error)
	StoreProduct(name string, disc string) (*gorm.DB, error)
	GetOneProduct(id string) (*gorm.DB, error)
}

type DefaultProductServices struct {
	repo domain.ProductRepository
}

func (d DefaultProductServices) GetProduct() (*gorm.DB, error) {
	return d.repo.FindAllProduct()
}

func (d DefaultProductServices) StoreProduct(name string, disc string) (*gorm.DB, error) {
	return d.repo.StoreProduct(name, disc)
}

func (d DefaultProductServices) GetOneProduct(id string) (*gorm.DB, error) {
	return d.repo.OneProduct(id)
}
func AllProductServices(d domain.ProductRepository) DefaultProductServices {
	return DefaultProductServices{d}
}
