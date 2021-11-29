package service

import (
	"crud_echo/domain"
	"github.com/jinzhu/gorm"
)

type UserServices interface {
	GetAllUsers() (*gorm.DB, error)
	GetUser(id string) (*gorm.DB, error)
	StoreUser(first_name string, full_name string, mobile string, status int64, password string) (*gorm.DB, error)
}

type DefaultUserServices struct {
	repo domain.UserRepository
}

func (d DefaultUserServices) GetAllUsers() (*gorm.DB, error) {
	return d.repo.FindAll()
}
func (d DefaultUserServices) GetUser(id string) (*gorm.DB, error) {
	return d.repo.FindOne(id)
}

func (d DefaultUserServices) StoreUser(first_name string, full_name string, mobile string, status int64, password string) (*gorm.DB, error) {
	return d.repo.Store(first_name, full_name, mobile, status, password)
}

func AllUserServices(r domain.UserRepository) DefaultUserServices {
	return DefaultUserServices{r}
}
