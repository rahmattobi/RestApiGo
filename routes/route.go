package routes

import (
	"auth-jwt/handlers"

	"github.com/gofiber/fiber/v2"
)

func RouteInit(c *fiber.App) {
	c.Get("/user", handlers.HandlerUser)
	c.Get("/user/:id", handlers.HandlerUserGetById)
	c.Post("/user", handlers.HandlerUserInput)
	c.Put("/user/:id", handlers.HandlerUserUpdate)
	c.Delete("/user/:id", handlers.HandlerUserDelete)
}
