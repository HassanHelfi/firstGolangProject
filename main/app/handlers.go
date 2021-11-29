package app

import (
	"crud_echo/service"
	"fmt"
	"github.com/labstack/echo"
	"net/http"
	"strconv"
)

type UserHandlers struct {
	service service.UserServices
}

func (uh *UserHandlers) GetAllUsers(c echo.Context) error {
	users, _ := uh.service.GetAllUsers()
	err := c.JSON(http.StatusOK, users)

	if err != nil {
		return nil
	}
	return nil
}

func (uh *UserHandlers) GetUser(c echo.Context) error {
	id := c.Param("id")
	fmt.Println(id)
	users, _ := uh.service.GetUser(id)
	err := c.JSON(http.StatusOK, users)

	if err != nil {
		return nil
	}
	return nil
}

func (uh *UserHandlers) Store(c echo.Context) error {
	first_name := c.FormValue("first_name")
	full_name := c.FormValue("full_name")
	//email := c.FormValue("email")
	mobile := c.FormValue("mobile")
	status, _ := strconv.ParseInt(c.FormValue("status"), 10, 64)
	password := c.FormValue("psassword")

	user, _ := uh.service.StoreUser(first_name, full_name, mobile, status, password)
	err := c.JSON(http.StatusOK, user)

	if err != nil {
		return nil
	}
	return nil
}

type ProductHandlers struct {
	service service.ProductServices
}

func (ph *ProductHandlers) GetAllProduct(c echo.Context) error {
	prod, _ := ph.service.GetProduct()
	err := c.JSON(http.StatusOK, prod)
	if err != nil {
		return nil
	}
	return nil
}

func (ph *ProductHandlers) StoreProduct(c echo.Context) error {
	name := c.FormValue("name")
	disc := c.FormValue("disc")
	prod, _ := ph.service.StoreProduct(name,disc)
	err := c.JSON(http.StatusOK, prod)
	if err != nil {
		return nil
	}
	return nil
}

func (ph *ProductHandlers) GetProduct(c echo.Context) error {
	id := c.Param("product_id")
	fmt.Println(id)
	prod, _ := ph.service.GetOneProduct(id)
	err := c.JSON(http.StatusOK, prod)
	if err != nil {
		return nil
	}
	return nil
}