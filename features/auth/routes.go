package auth

import (
	"github.com/ghozir/sighapp/middleware"
	"github.com/gofiber/fiber/v2"
)

func AuthRoutes(app *fiber.App, service Service) {
	auth := app.Group("/auth")
	auth.Get("/token", middleware.JWTAuth(NewAuthRepository()), GetToken(service))
	auth.Post("/login", middleware.BasicAuth(), LoginHandler(service))
}
