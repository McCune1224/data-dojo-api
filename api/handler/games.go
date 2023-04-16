package handler

import (
	"github.com/gofiber/fiber/v2"
)

func GetAllGames(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "TODO: Get all games",
	})
}
