package api

import (
	"crud_echo/configs"
	"crud_echo/models"
	"github.com/labstack/echo"
	"net/http"
	"strconv"
	"time"
)

var user []models.User

func Index(c echo.Context) error {
	rows := configs.CreateCon().Find(&user)
	return c.JSON(http.StatusOK, rows)
}

func Show(c echo.Context) error {
	id := c.Param("id")
	prod := configs.CreateCon().Find(&models.Product{}, id)
	return c.JSON(http.StatusOK, prod)
}

func Store(c echo.Context) error {
	first_name := c.FormValue("first_name")
	full_name := c.FormValue("full_name")
	email := c.FormValue("email")
	mobile := c.FormValue("mobile")
	status, _ := strconv.ParseInt(c.FormValue("status"), 10, 64)
	psassword := c.FormValue("psassword")

	//roles := []int{c.FormValue("roles")}

	user:= configs.CreateCon().Create(&models.User{
		FirstName: first_name,
		LastName:  full_name,
		Email:     &email,
		Mobile:    mobile,
		VerifyAt:  time.Now(),
		Status:    uint(status),
		Password:  psassword,
	})

	return c.JSON(http.StatusOK, user)
}
func Update(c echo.Context) error {

	user_id := c.Param("id")

	first_name := c.FormValue("first_ame")
	last_name := c.FormValue("full_name")
	email := c.FormValue("email")
	mobile := c.FormValue("mobile")
	status, _ := strconv.ParseInt(c.FormValue("status"), 10, 64)
	psassword := c.FormValue("psassword")
	person := configs.CreateCon().First(&models.User{}, user_id).Update(models.User{
		FirstName: first_name,
		LastName:  last_name,
		Email:     &email,
		Mobile:    mobile,
		VerifyAt:  time.Now(),
		Status:    uint(status),
		Password:  psassword,
	})

	return c.JSON(http.StatusOK, person)
}

func Delete(c echo.Context) error {

	user_id := c.Param("id")

	person := configs.CreateCon().Delete(&models.User{}, user_id)

	return c.JSON(http.StatusOK, person)
}
