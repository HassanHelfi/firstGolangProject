package models

import "github.com/jinzhu/gorm"

type Zipcode struct {
	gorm.Model
	Code      string
	Country   string
	City      string
	Address   string
	Longitude float64
	Latitude  float64
	Users     *[]User `gorm:"many2many:user_zipcode;"`
}
