package configs

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"os"
)

func CreateCon() *gorm.DB {
	dbCon, err := gorm.Open("postgres", dbConfig())
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("db is connected")
	}

	err = dbCon.Error
	fmt.Println(err)
	if err != nil {
		fmt.Println("db is not connected")
		fmt.Println(err.Error())
	}
	return dbCon
}

func dbConfig() string {
	conf := "user=" + os.Getenv("DB_USER") +
		" " + "port=" + os.Getenv("DB_PORT") +
		" " + "password=" + os.Getenv("DB_PASSWORD") +
		" " + "dbname=" + os.Getenv("DB_NAME") +
		" " + "sslmode=" + os.Getenv("DB_SSLMODE")
	return conf
}
