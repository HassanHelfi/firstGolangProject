package main

import (
	"crud_echo/migrations"
	"crud_echo/routes"
	"github.com/joho/godotenv"
)

func main() {
	//load .env file
	godotenv.Load(".env")
	// auto migrate
	migrations.Migrate()
	// routes
	routes.Routes()
}
