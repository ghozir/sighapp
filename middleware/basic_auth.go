package middleware

import (
	"encoding/base64"
	"strings"

	env "github.com/ghozir/sighapp/config"
	"github.com/ghozir/sighapp/utils/exception"
	"github.com/gofiber/fiber/v2"
)

func BasicAuth() fiber.Handler {
	return func(c *fiber.Ctx) error {
		auth := c.Get("Authorization")
		if !strings.HasPrefix(auth, "Basic ") {
			return exception.Unauthorized("Missing Basic Auth header")
		}

		payload := strings.TrimPrefix(auth, "Basic ")
		decoded, err := base64.StdEncoding.DecodeString(payload)
		if err != nil {
			return exception.Unauthorized("Invalid base64 encoding in Authorization header")
		}

		parts := strings.SplitN(string(decoded), ":", 2)
		if len(parts) != 2 {
			return exception.Unauthorized("Invalid auth format, expected username:password")
		}

		username, password := parts[0], parts[1]

		if username != env.Config.BasicAuthUsername || password != env.Config.BasicAuthPassword {
			return exception.Unauthorized("Invalid credentials")
		}

		return c.Next()
	}
}
