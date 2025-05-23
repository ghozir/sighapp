package main

import (
	"log"

	env "github.com/ghozir/sighapp/config"
	"github.com/ghozir/sighapp/database/mongodb"
	"github.com/ghozir/sighapp/features/auth"
	"github.com/ghozir/sighapp/middleware"
	"github.com/gofiber/fiber/v2"
)

func main() {
	env.ConfigInit()
	mongodb.ConnectMongo()
	app := fiber.New()
	// 💥 Pasang middleware error global
	app.Use(middleware.ErrorMiddleware())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("SighApp is alive!")
	})

	authRepo := auth.NewAuthRepository()
	authService := auth.NewAuthService(authRepo)
	auth.AuthRoutes(app, authService)

	log.Fatal(app.Listen(":" + env.Config.AppPort))
}
