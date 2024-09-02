package utils

import (
	"github.com/gofiber/fiber/v2"
)

// Response is the standard response structure
type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// NewResponse creates a new response object
func NewResponse(code int, message string, data interface{}) *Response {
	return &Response{
		Code:    code,
		Message: message,
		Data:    data,
	}
}

// SendResponse sends a JSON response to the client
func SendResponse(c *fiber.Ctx, code int, message string, data interface{}) error {
	response := NewResponse(code, message, data)
	return c.Status(fiber.StatusOK).JSON(response)
}

// SendSuccess sends a successful JSON response with a default code of 200
func SendSuccess(c *fiber.Ctx, data interface{}) error {
	return SendResponse(c, fiber.StatusOK, "Success", data)
}

// SendError sends an error JSON response with a specified code and message
func SendError(c *fiber.Ctx, code int, message string) error {
	return SendResponse(c, code, message, nil)
}
