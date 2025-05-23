package auth

import (
	"github.com/ghozir/sighapp/features/auth/dto"
	"github.com/ghozir/sighapp/utils"
	"github.com/ghozir/sighapp/utils/exception"
	"github.com/ghozir/sighapp/utils/response"
	"github.com/gofiber/fiber/v2"
)

func LoginHandler(service Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req dto.LoginRequest

		if err := c.BodyParser(&req); err != nil {
			return exception.BadRequest("invalid request body")
		}

		if errs := utils.ValidateStruct(req); errs != nil {
			return exception.UnprocessableEntityWithData("validation failed", errs)
		}

		result, err := service.LoginUser(c, req)
		if err != nil {
			return err
		}

		return response.OK(c, "Login success", "LOGIN_SUCCESS", result)
	}
}

func GetToken(service Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return response.OK(c, "Login success", "LOGIN_SUCCESS", c.Locals("user"))
	}
}
