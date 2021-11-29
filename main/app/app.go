package app

import (
	"crud_echo/domain"
	"crud_echo/migrations"
	"crud_echo/service"
	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	"os"
)

func Start() {
	//load .env file
	godotenv.Load(".env")
	// auto migrate
	migrations.Migrate()
	migrations.Migrate()
	// routes
	e := echo.New()
	uh := UserHandlers{service.AllUserServices(domain.AllUserRepositoryDB())}
	ph := ProductHandlers{service.AllProductServices(domain.AllProductRepositoryDB())}
	e.GET("/users", uh.GetAllUsers)
	e.GET("/users/:id", uh.GetUser)
	e.POST("/users/store", uh.Store)
	///
	e.GET("/products", ph.GetAllProduct)
	e.GET("/products/:product_id", ph.GetProduct)
	e.POST("/products/store", ph.StoreProduct)
	//routes.Routes()
	e.Logger.Fatal(e.Start(":" +os.Getenv("APP_PORT")))

}
