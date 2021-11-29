package domain

import "github.com/jinzhu/gorm"

type Shipper struct {
	gorm.Model
	Name string
	Fee  int
}
