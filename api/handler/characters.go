package handler

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/mccune1224/data-dojo/api/model"
	"github.com/mccune1224/data-dojo/api/store"
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

func GetAllCharacters(c *fiber.Ctx) error {
	type characterResponse struct {
		ID     uint   `json:"id"`
		Name   string `json:"name"`
		GameID uint   `json:"game_id"`
	}
	gameID := c.Params("gameID")
	if gameID == "" {
		return c.Status(400).JSON(fiber.Map{
			"error": "Please provide a game ID",
		})
	}

	dbChars := []model.Character{}
	err := store.DB.
		Where("game_id = ?", gameID).
		Find(&dbChars).Error

	if err != nil {
		return handleNotFound(c, err)
	}

	charResponse := []characterResponse{}
	for i := range dbChars {

		charResponse = append(charResponse, characterResponse{
			ID:     dbChars[i].ID,
			Name:   dbChars[i].Name,
			GameID: dbChars[i].GameID,
		})
	}

	return c.JSON(fiber.Map{
		"characters": charResponse,
	})
}

func GetCharacterByID(c *fiber.Ctx) error {
	type characterResponse struct {
		ID     uint   `json:"id"`
		Name   string `json:"name"`
		GameID uint   `json:"game_id"`
	}
	gameIDParam := c.Params("gameID")
	if gameIDParam == "" {
		return c.Status(400).JSON(fiber.Map{
			"error": "Please provide a game ID",
		})
	}

	characterIDParam := c.Params("id")
	if characterIDParam == "" {
		return c.Status(400).JSON(fiber.Map{
			"error": "Please provide a character ID",
		})
	}

	dbChar := model.Character{}
	err := store.DB.Where("id = ? and game_id = ?", characterIDParam, gameIDParam).First(&dbChar).Error
	if err != nil {
		return handleNotFound(c, err)
	}

	return c.JSON(fiber.Map{
		"character": struct {
			ID     uint   `json:"id"`
			Name   string `json:"name"`
			GameID uint   `json:"game_id"`
		}{
			ID:     dbChar.ID,
			Name:   dbChar.Name,
			GameID: dbChar.GameID,
		},
	})
}

func SearchCharacters(c *fiber.Ctx) error {
	type characterResponse struct {
		ID     uint   `json:"id"`
		Name   string `json:"name"`
		GameID uint   `json:"game_id"`
	}
	requestQuery := c.Query("name")
	dbResults := []model.Character{}
	err := store.DB.
		// Forgive me father for I have sinned and used ILIKE
		// Shoutout to the postgresql gods for making this possible
		Where("name ILIKE ?", "%"+requestQuery+"%").
		Find(&dbResults).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(404).JSON(fiber.Map{
				"error": "No characters found with query " + requestQuery,
			})
		} else {
			return c.Status(500).JSON(fiber.Map{
				"error":  "Internal server error",
				"reason": err.Error(),
			})
		}
	}

	charactersResponse := []characterResponse{}
	for i := range dbResults {
		charactersResponse = append(charactersResponse, characterResponse{
			ID:     dbResults[i].ID,
			Name:   dbResults[i].Name,
			GameID: dbResults[i].GameID,
		})
	}

	return c.JSON(fiber.Map{
		"characters": charactersResponse,
	})

}
