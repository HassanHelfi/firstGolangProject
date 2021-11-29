package domain

import (
	"github.com/jinzhu/gorm"
	"time"
)

type User struct {
	//gorm.Model
	FirstName string
	LastName  string
	Email     *string
	Mobile    string
	VerifyAt  time.Time
	Status    uint
	Password  string
	/*Roles       *[]Role       `gorm:"many2many:role_user;"`
	Permissions *[]Permission `gorm:"many2many:permission_user;"`
	Zipcodes    *[]Zipcode    `gorm:"many2many:user_zipcode;"`*/
}

type UserRepository interface {
	FindAll() (*gorm.DB, error)
	FindOne(id string) (*gorm.DB, error)
	Store(name string, name2 string, mobile string, status int64, password string) (*gorm.DB, error)
}
