package domain

import (
	"time"
	"gorm.io/plugin/soft_delete"
)

type ProductVendor struct {
	ProductID uint `gorm:"primaryKey"`
	VendorID  uint `gorm:"primaryKey"`
	Quantity  uint
	Price     uint
	CreatedAt time.Time
	DeletedAt soft_delete.DeletedAt `gorm:"uniqueIndex:deleted_at"`
}
