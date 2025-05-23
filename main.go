package main

import (
	"log"

	env "github.com/ghozir/sighapp/config"
	"github.com/ghozir/sighapp/database/mongodb"
	"github.com/ghozir/sighapp/features/auth"
	"github.com/ghozir/sighapp/features/problem"
	"github.com/ghozir/sighapp/middleware"
	"github.com/gofiber/fiber/v2"
)

func main() {
	env.ConfigInit()
	mongodb.ConnectMongo()
	app := fiber.New()
	app.Use(middleware.ErrorHandler())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("SighApp is alive!")
	})

	authRepo := auth.NewAuthRepository()
	app.Use(middleware.JWTAuth(authRepo))

	authService := auth.NewAuthService(authRepo)
	auth.AuthRoutes(app, authService)

	problemService := problem.NewProblemService(problem.NewProblemRepository())
	problem.ProblemRoutes(app, problemService)

	log.Fatal(app.Listen(":" + env.Config.AppPort))
}
