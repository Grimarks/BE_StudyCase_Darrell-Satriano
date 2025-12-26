package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func AdminOnly(c *fiber.Ctx) error {
	user := c.Locals("user")

	claims := user.(jwt.MapClaims)

	if claims["role"] != "admin" {
		return c.Status(403).JSON(fiber.Map{
			"error": "Admin only",
		})
	}

	return c.Next()
}
