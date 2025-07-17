package middleware

import (
	"os"

	"github.com/alfianvitoanggoro/boilerplate-simple/pkg/response"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func JWTMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		tokenString := c.Get("Authorization")
		if tokenString == "" {
			return response.WriteError(c, fiber.StatusUnauthorized, "Unauthorized", "Missing token")
		}
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
			return []byte(os.Getenv("JWT_SECRET")), nil
		})
		if err != nil || !token.Valid {
			return response.WriteError(c, fiber.StatusUnauthorized, "Unauthorized", "Invalid token")
		}
		return c.Next()
	}
}
