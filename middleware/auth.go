package middleware

import (
	"slices"
	"strings"

	"github.com/Abhishek2010dev/movie-management-system/models"
	"github.com/Abhishek2010dev/movie-management-system/utils"
	"github.com/gofiber/fiber/v3"
)

const AuthPayloadKey = "auth-payload-key"

func AuthMiddleware(secretKey []byte) fiber.Handler {
	return func(c fiber.Ctx) error {
		auth := c.Get(fiber.HeaderAuthorization)
		if auth == "" || !strings.HasPrefix(auth, "Bearer ") {
			return fiber.NewError(fiber.StatusUnauthorized, "Missing or invalid Authorization header")
		}
		tokenString := strings.TrimPrefix(auth, "Bearer ")
		claims, err := utils.VerifyToken(secretKey, tokenString)
		if err != nil {
			return fiber.NewError(fiber.StatusUnauthorized, err.Error())
		}
		c.Locals(AuthPayloadKey, *claims)
		return c.Next()
	}
}

func RequireRoles(roles ...models.UserRole) fiber.Handler {
	return func(c fiber.Ctx) error {
		claims := fiber.Locals[utils.Claims](c, AuthPayloadKey)
		if slices.Contains(roles, claims.Role) {
			return c.Next()
		}
		return fiber.NewError(fiber.StatusForbidden, "You are not allowed to access this route")
	}
}
