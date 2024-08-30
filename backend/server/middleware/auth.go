package middleware

import (
	"github.com/gofiber/fiber/v2"
	"strings"
)

// AuthMiddleware 是一个简单的鉴权中间件
func AuthMiddleware(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")
	if !strings.HasPrefix(authHeader, "Bearer ") {
		return c.Status(fiber.StatusUnauthorized).SendString("未授权")
	}
	return c.Next()
}
