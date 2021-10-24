package routes

import (
	"github.com/hassan/crud_echo/controllers/api"
	"github.com/labstack/echo"
	"net/http"
)

func Routes(e *echo.Echo) {
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusCreated, "welcome")
	})
	e.GET("/test", api.GetProducts)
}
