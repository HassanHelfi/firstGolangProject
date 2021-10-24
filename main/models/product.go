package models

import (
	_ "github.com/labstack/echo"
)

type Product struct {
	Id   int32 `gorm:"unique;not null"`
	Name string `json:"name"`
}