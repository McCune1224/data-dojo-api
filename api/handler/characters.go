package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mccune1224/data-dojo/api/model"
	"github.com/mccune1224/data-dojo/api/store"
)

// JSON response for a character
type characterResponse struct {
	ID     uint   `json:"id"`
	Name   string `json:"name"`
	GameID uint   `json:"game_id"`
}

func GetAllCharacters(c *fiber.Ctx) error {
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
	err := store.DB.
		Where("id = ? and game_id = ?", characterIDParam, gameIDParam).
		First(&dbChar).Error
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
	requestQuery := c.Query("name")
	dbResults := []model.Character{}
	err := store.DB.
		Where("name ILIKE ?", "%"+requestQuery+"%").
		Find(&dbResults).Error

	if err != nil {
		handleNotFound(c, err)
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
