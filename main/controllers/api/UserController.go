package api

import (
	"crud_echo/configs"
	"crud_echo/domain"
	"fmt"
	"github.com/labstack/echo"
	"net/http"
	"strconv"
	"time"
)

var user []domain.User

func Index(c echo.Context) error {
	rows := configs.CreateCon().Find(&user)
	err := configs.RedisCon().Set("users_index", rows, 0).Err()
	if err != nil {
		fmt.Println(err)
	}
	user_index, err := configs.RedisCon().Get("user_index").Result()
	if err != nil {
		fmt.Println(err)
	}
	return c.JSON(http.StatusOK, user_index)
}

func Show(c echo.Context) error {
	id := c.Param("id")
	prod := configs.CreateCon().Find(&domain.User{}, id)
	return c.JSON(http.StatusOK, prod)
}

func Store(c echo.Context) error {
	first_name := c.FormValue("first_name")
	full_name := c.FormValue("full_name")
	//email := c.FormValue("email")
	mobile := c.FormValue("mobile")
	status, _ := strconv.ParseInt(c.FormValue("status"), 10, 64)
	psassword := c.FormValue("psassword")

	//roles := []int{c.FormValue("roles")}

	user := configs.CreateCon().Create(&domain.User{
		FirstName: first_name,
		LastName:  full_name,
		//Email:     &email,
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
	//email := c.FormValue("email")
	mobile := c.FormValue("mobile")
	status, _ := strconv.ParseInt(c.FormValue("status"), 10, 64)
	psassword := c.FormValue("psassword")
	person := configs.CreateCon().First(&domain.User{}, user_id).Update(domain.User{
		FirstName: first_name,
		LastName:  last_name,
		//Email:     &email,
		Mobile:    mobile,
		VerifyAt:  time.Now(),
		Status:    uint(status),
		Password:  psassword,
	})

	return c.JSON(http.StatusOK, person)
}

func Delete(c echo.Context) error {

	user_id := c.Param("id")

	person := configs.CreateCon().Delete(&domain.User{}, user_id)

	return c.JSON(http.StatusOK, person)
}
