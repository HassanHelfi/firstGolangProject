package domain

import "github.com/jinzhu/gorm"

type Role struct {
	gorm.Model
	Name        string
	Permissions *[]Permission `gorm:"many2many:permission_role;"`
}
