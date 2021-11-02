package routes

import (
	"crud_echo/controllers/api"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"os"
)

func Routes() {
	e := echo.New()
	e.GET("/user/index", api.Index)
	e.GET("/user/show/:id", api.Show)
	e.POST("/user/store", api.Store)
	e.PUT("/user/update/:id", api.Update)
	e.DELETE("/user/delete/:id", api.Delete)

	adminGroup := e.Group("/jwt")
	adminGroup.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningMethod: "HS512",
		SigningKey: []byte("test-test"),
	}))
	adminGroup.GET("/user/index", api.Index)

	//login
	e.POST("/login", api.Login)
	e.Logger.Fatal(e.Start(":" +os.Getenv("APP_PORT")))
}
