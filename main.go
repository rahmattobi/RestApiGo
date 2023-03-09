package main

import (
	"auth-jwt/database"
	"auth-jwt/database/migration"
	"auth-jwt/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	// Initial Database
	database.DatabaseInit()
	// DB Migration
	migration.RunMigration()
	// Initial Routes
	routes.RouteInit(app)

	app.Listen(":8000")
}
