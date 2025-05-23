package auth

import (
	"github.com/gofiber/fiber/v2"
)

func AuthRoutes(app *fiber.App, service Service) {
	auth := app.Group("/auth")
	auth.Get("/token", GetToken(service))
	auth.Post("/login", LoginHandler(service))
}
