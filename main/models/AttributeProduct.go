package models

import (
	"gorm.io/plugin/soft_delete"
	"time"
)

type AttributeProduct struct {
	AttributeID uint   `gorm:"primaryKey"`
	ProductID   uint   `gorm:"primaryKey"`
	Value       string `gorm:"type:text"`
	CreatedAt   time.Time
	DeletedAt   soft_delete.DeletedAt `gorm:"uniqueIndex:deleted_at"`
}
