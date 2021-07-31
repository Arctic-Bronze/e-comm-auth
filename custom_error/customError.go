package custom_error

import "github.com/gofiber/fiber/v2"

func GetDefaultInternalError(message string) *fiber.Error {
	return &fiber.Error{Code: 500, Message: message}
}
