package problem

import (
	problemdto "github.com/ghozir/sighapp/features/problem/dto"
	"github.com/ghozir/sighapp/utils"
	"github.com/ghozir/sighapp/utils/exception"
	"github.com/ghozir/sighapp/utils/response"
	"github.com/gofiber/fiber/v2"
)

func InsertProblemData(service Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req problemdto.InsertProblemRequest

		if err := c.BodyParser(&req); err != nil {
			return exception.BadRequest("invalid request body")
		}

		if errs := utils.ValidateStruct(req); errs != nil {
			return exception.UnprocessableEntityWithData("validation failed", errs)
		}

		result, err := service.InsertProblem(c, req)
		if err != nil {
			return err
		}

		return response.OK(c, "Login success", "LOGIN_SUCCESS", result)
	}
}
