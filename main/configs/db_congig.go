package configs

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func CreateCon() *gorm.DB {
	db, err := gorm.Open("postgres", "user=postgres password=root dbname=postgres sslmode=disable")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("db is connected")
	}

	err = db.Error
	fmt.Println(err)
	if err != nil {
		fmt.Println("db is not connected")
		fmt.Println(err.Error())
	}
	return db
}
