package middleware

import "github.com/gofiber/fiber/v2"

func AdminOnly(c *fiber.Ctx) error {
	user := c.Locals("user")
	claims := user.(map[string]interface{})

	if claims["role"] != "admin" {
		return c.Status(403).JSON(fiber.Map{
			"error": "Admin only",
		})
	}

	return c.Next()
}
