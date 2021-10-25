package routes

import (
	"github.com/labstack/echo"
	"net/http"
)

func Routes() {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusCreated, "welcome")
	})

	err := e.Start(":8081")

	if err != nil {
		return
	}
}
