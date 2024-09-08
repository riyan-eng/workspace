package api

import (
	"github.com/gofiber/fiber/v2"
)

func (m *ServiceServer) Index(c *fiber.Ctx) error {
	return c.Status(200).JSON(fiber.Map{
		"message": "Welcome to my channel",
	})
}
