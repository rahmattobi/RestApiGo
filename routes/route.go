package routes

import (
	"auth-jwt/handlers"
	"auth-jwt/middleware"

	"github.com/gofiber/fiber/v2"
)

func RouteInit(c *fiber.App) {

	c.Post("/login", handlers.Login)

	c.Get("/user", middleware.Auth, handlers.HandlerUser)
	c.Get("/user/:id", middleware.Auth, handlers.HandlerUserGetById)
	c.Post("/user", middleware.Auth, handlers.HandlerUserInput)
	c.Put("/user/:id", middleware.Auth, handlers.HandlerUserUpdate)
	c.Delete("/user/:id", middleware.Auth, handlers.HandlerUserDelete)
}
