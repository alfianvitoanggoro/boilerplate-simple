package middleware

import (
	"github.com/alfianvitoanggoro/boilerplate-simple/pkg/response"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func RBACMiddleware(roles ...string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		tokenString := c.Get("Authorization")
		token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
			return []byte(""), nil // JWT secret not needed for role check here
		})
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			userRole := claims["role"].(string)
			for _, role := range roles {
				if userRole == role {
					return c.Next()
				}
			}
		}
		return response.WriteError(c, fiber.StatusForbidden, "Forbidden", "Insufficient permissions")
	}
}
