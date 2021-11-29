package domain

import "github.com/jinzhu/gorm"

type Comment struct {
	gorm.Model
	Content   string `gorm:"type:text"`
	Status    uint
	ParentID  int `gorm:"TYPE:integer REFERENCES comments"`
	Parent    *Comment
	UserID    int
	User      User
	ProductID int
	Product   Product
}
