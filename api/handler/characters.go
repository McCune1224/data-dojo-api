package handler

import (
	"errors"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/mccune1224/data-dojo/api/model"
	"github.com/mccune1224/data-dojo/api/store"
	"gorm.io/gorm"
)

// JSON response for a character
type characterResponse struct {
	ID     uint   `json:"id"`
	Name   string `json:"name"`
	GameID uint   `json:"game_id"`
}

func GetAllCharacters(c *fiber.Ctx) error {
	// Get game ID from query
	gameID, err := c.ParamsInt("gameID")
	if err != nil {
		return c.Status(400).JSON(&ErrorResponse{
			Error:            Error400String,
			ErrorDescription: "please provide a game id",
		})
	}
	// Search DB for all characters of associated game
	dbChars := []model.Character{}
	err = store.DB.
		Where("game_id = ?", gameID).
		Find(&dbChars).Error
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(500).JSON(Error500Response)
		}
		return c.Status(404).JSON(&ErrorResponse{
			Error:            Error404String,
			ErrorDescription: "Could not find characters",
		})
	}

	// Create and return characters response
	charResponse := []characterResponse{}
	for i := range dbChars {
		charResponse = append(charResponse, characterResponse{
			ID:     dbChars[i].ID,
			Name:   dbChars[i].Name,
			GameID: dbChars[i].GameID,
		})
	}
	return c.JSON(charResponse)
}

func GetCharacterByID(c *fiber.Ctx) error {
	// Get game and character ID from query
	gameIDParam, err := c.ParamsInt("gameID")
	if err != nil {
		return c.Status(400).JSON(&ErrorResponse{
			Error:            Error400String,
			ErrorDescription: "please provide a game id",
		})
	}
	characterIDParam, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON(&ErrorResponse{
			Error:            Error400String,
			ErrorDescription: "please provide a character id",
		})
	}

	// Search DB for character of associated game
	dbChar := model.Character{}
	err = store.DB.
		Where("id = ? and game_id = ?", characterIDParam, gameIDParam).
		First(&dbChar).Error
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(500).JSON(Error500Response)
		}
		return c.Status(404).JSON(&ErrorResponse{
			Error:            Error404String,
			ErrorDescription: "Could not find character",
		})
	}

	return c.JSON(&characterResponse{
		ID:     dbChar.ID,
		Name:   dbChar.Name,
		GameID: dbChar.GameID,
	})
}

func SearchCharacters(c *fiber.Ctx) error {
	requestQuery := c.Query("name")
	limitQuery := c.Query("limit")
	limitQueryInt, err := strconv.Atoi(limitQuery)
	if err != nil {
		limitQueryInt = 10
	}
	if requestQuery == "" {
		return c.Status(400).JSON(&ErrorResponse{
			Error:            Error400String,
			ErrorDescription: "Please provide a name to search for",
		},
		)
	}
	dbResults := []model.Character{}
	err = store.DB.
		Where("name ILIKE ?", "%"+requestQuery+"%").
		Find(&dbResults).
		Limit(limitQueryInt).
		Error
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(500).JSON(Error500Response)
		}
		return c.Status(404).JSON(&ErrorResponse{
			Error:            Error404String,
			ErrorDescription: "Could not find characters",
		})
	}

	charactersResponse := []characterResponse{}
	for i := range dbResults {
		charactersResponse = append(charactersResponse, characterResponse{
			ID:     dbResults[i].ID,
			Name:   dbResults[i].Name,
			GameID: dbResults[i].GameID,
		})
	}

	return c.JSON(charactersResponse)
}
