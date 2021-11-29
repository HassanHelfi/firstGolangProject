package domain

import "github.com/jinzhu/gorm"

type Permission struct {
	gorm.Model
	Name  string
	Roles *[]Role `gorm:"many2many:permission_role;"`
}
