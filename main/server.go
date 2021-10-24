package main

import (
	"github.com/hassan/crud_echo/routes"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"os"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.PUT, echo.DELETE},
	}))

	routes.Routes(e)
	err := e.Start(":" + os.Getenv("APP_PORT"))
	if err != nil {
		return 
	}
	//e.Logger.Fatal(e.Start("8001"))

}
