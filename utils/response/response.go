package response

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func successResponse(httpCode int, message string, code string, data interface{}, pagination interface{}) fiber.Map {
	res := fiber.Map{
		"status":   true,
		"message":  message,
		"code":     code,
		"data":     data,
		"eTag":     uuid.NewString(),
		"httpCode": httpCode,
	}

	if pagination != nil {
		res["pagination"] = pagination
	}

	return res
}

// 200 OK
func OK(c *fiber.Ctx, message string, code string, data interface{}) error {
	return c.Status(http.StatusOK).JSON(successResponse(http.StatusOK, message, code, data, nil))
}

// 201 Created
func Created(c *fiber.Ctx, message string, code string, data interface{}) error {
	return c.Status(http.StatusCreated).JSON(successResponse(http.StatusCreated, message, code, data, nil))
}

// 200 OK with pagination
func OKWithPagination(c *fiber.Ctx, message string, code string, data interface{}, pagination interface{}) error {
	return c.Status(http.StatusOK).JSON(successResponse(http.StatusOK, message, code, data, pagination))
}

// 204 No content
func NoContent(c *fiber.Ctx) error {
	return c.SendStatus(http.StatusNoContent)
}
