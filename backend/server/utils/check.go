package utils

import "github.com/gofiber/fiber/v2"

func CheckIsAdminReq(c *fiber.Ctx) bool {
	user, err := GetUserByReq(c)
	if err != nil || user == nil {
		return false
	}
	return user.Role == "admin"
}
