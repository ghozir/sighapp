package middleware

import (
	"github.com/ghozir/sighapp/utils/exception"
	"github.com/gofiber/fiber/v2"
)

func ErrorHandler() fiber.Handler {
	return func(c *fiber.Ctx) error {
		err := c.Next()
		if err != nil {
			if e, ok := err.(*exception.HttpError); ok {
				return c.Status(e.Code).JSON(e)
			}

			internal := exception.InternalServerError("Internal server error")
			return c.Status(internal.Code).JSON(internal)
		}
		return nil
	}
}
