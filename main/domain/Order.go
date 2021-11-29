package domain

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Order struct {
	gorm.Model
	Name          string
	Tax           int
	Discount      int
	TotalPrice    uint
	ShipDate      time.Time
	ShipperID     uint
	Shipper       Shipper
	UserID        uint
	User          User
	OrderStatusID uint
	OrderStatus   OrderStatus `gorm:"foreignKey:OrderStatusID;"`
}
