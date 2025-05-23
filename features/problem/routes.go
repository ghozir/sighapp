package problem

import (
	"github.com/gofiber/fiber/v2"
)

func ProblemRoutes(app *fiber.App, service Service) {
	problem := app.Group("/problem")
	problem.Post("/", InsertProblemData(service))
	problem.Post("/anonim", InsertProblemData(service))
}
