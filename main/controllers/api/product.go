package api

import (
	DB "crud_echo/configs"
	"crud_echo/models"
	"github.com/labstack/echo"
	"net/http"
)

var product []models.Product
var message map[string]string

func Index(c echo.Context) error {
	rows := DB.CreateCon().Find(&product)
	return c.JSON(http.StatusOK, rows)
}

func Show(c echo.Context) error{
	id := c.Param("id")
	prod := DB.CreateCon().Find(&models.Product{}, id)
	return c.JSON(http.StatusOK, prod)
}

//func GetProducts(c echo.Context) error {
//	result := models.GetProduct()
//	println("foo")
//	return c.JSON(http.StatusOK, result)
//}
