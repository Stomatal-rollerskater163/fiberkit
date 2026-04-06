package middlewares

import (
	"strings"

	"github.com/gofiber/fiber/v3"
)

func Auth(ctx fiber.Ctx) error {
	authHeader := ctx.Get(fiber.HeaderAuthorization)
	if !strings.HasPrefix(authHeader, "Bearer ") {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "missing bearer token",
		})
	}

	return ctx.Next()
}
