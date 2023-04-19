package handler

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func handleNotFound(c *fiber.Ctx, err error) error {
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return c.Status(404).JSON(fiber.Map{
			"error": "No records found",
		})
	} else {
		return c.Status(500).JSON(fiber.Map{
			"error":  "Internal server error",
			"reason": err.Error(),
		})
	}
}
