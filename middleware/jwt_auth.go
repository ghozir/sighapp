package middleware

import (
	"strings"

	"github.com/ghozir/sighapp/features/auth"
	"github.com/ghozir/sighapp/utils"
	"github.com/ghozir/sighapp/utils/exception"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var whiteList = map[string]bool{
	"/":               true,
	"/auth/login":     true,
	"/problem/anonim": true,
}

func JWTAuth(repo auth.Repository) fiber.Handler {
	return func(c *fiber.Ctx) error {
		if whiteList[c.Path()] {
			return BasicAuth()(c)
		}
		authHeader := c.Get("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			return exception.BadRequest("Missing or invalid token")
		}

		token := strings.TrimPrefix(authHeader, "Bearer ")
		userID, jti, err := utils.DecodeJWT(token)
		if err != nil {
			return exception.Unauthorized("Invalid or expired token")
		}

		session, err := repo.FindOneToken(bson.M{
			"userId": userID,
			"jti":    jti,
		})
		if err != nil {
			return exception.Unauthorized("Session invalid or expired")
		}

		id, err := primitive.ObjectIDFromHex(session.UserID)
		if err != nil {
			return exception.BadRequest("Invalid user ID format")
		}

		user, err := repo.FindOneUser(bson.M{"_id": id})
		if err != nil {
			return exception.Unauthorized("User Not Found")
		}
		c.Locals("user", user.ToSafeUser())
		return c.Next()
	}
}
