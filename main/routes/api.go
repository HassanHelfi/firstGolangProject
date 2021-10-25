package routes

import (
	"crud_echo/controllers/api"
	"github.com/labstack/echo"
	"os"
)

func Routes() {
	e := echo.New()

	e.GET("/user/index", api.Index)
	e.GET("/user/show/:id", api.Show)
	e.POST("/user/store", api.Store)
	e.PUT("/user/update/:id", api.Update)
	e.DELETE("/user/delete/:id", api.Delete)

	e.Logger.Fatal(e.Start(":" +os.Getenv("APP_PORT")))
}
