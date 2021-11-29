package domain

import (
	"crud_echo/configs"
	"github.com/jinzhu/gorm"
	"time"
)

type UserRepositoryDB struct {
	client *gorm.DB
}

var user []User

func AllUserRepositoryDB() UserRepositoryDB {
	client := configs.CreateCon()
	return UserRepositoryDB{client}
}

func (u UserRepositoryDB) FindAll() (*gorm.DB, error) {

	users2 := u.client.Find(&user)
	return users2, nil
}

func (u UserRepositoryDB) FindOne(id string) (*gorm.DB, error) {
	users2 := u.client.Find(&user, id)
	return users2, nil
}

func (u UserRepositoryDB) Store(first_name string, full_name string, mobile string, status int64, password string) (*gorm.DB, error) {
	user := u.client.Create(&User{
		FirstName: first_name,
		LastName:  full_name,
		//Email:     &email,
		Mobile:   mobile,
		VerifyAt: time.Now(),
		Status:   uint(status),
		Password: password,
	})

	return user, nil
}
