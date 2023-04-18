package handler

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/mccune1224/data-dojo/api/model"
	"github.com/mccune1224/data-dojo/api/store"
	"github.com/mccune1224/data-dojo/api/store/queries"
	"gorm.io/gorm"
)

type CharacterResponse struct {
	ID     uint           `json:"id"`
	Name   string         `json:"name"`
	Moves  []MoveResponse `json:"moves"`
	GameID uint           `json:"gameID"`
}

func (cr *CharacterResponse) ModelToResponse(c model.Character) {
	cr.ID = c.ID
	cr.Name = c.Name
	cr.GameID = c.GameID

	for i := range c.Moves {
		move := MoveResponse{}
		move.ModelToResponse(c.Moves[i])
		cr.Moves = append(cr.Moves, move)
	}
}

func GetAllCharacters(c *fiber.Ctx) error {
	gameID := c.Params("gameID")
	if gameID == "" {
		return c.Status(400).JSON(fiber.Map{
			"error": "Please provide a game ID",
		})
	}

	dbChars := []model.Character{}
	err := queries.AllCharactersByID(gameID, &dbChars)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(404).JSON(fiber.Map{
				"error": "No characters found for game " + gameID,
			})
		} else {
			return c.Status(500).JSON(fiber.Map{
				"error":  "Internal server error",
				"reason": err.Error(),
			})
		}
	}

	chars := []CharacterResponse{}
	for i := range dbChars {
		chars = append(chars, CharacterResponse{
			ID:     dbChars[i].ID,
			Name:   dbChars[i].Name,
			GameID: dbChars[i].GameID,
		})
	}

	return c.JSON(fiber.Map{
		"characters": chars,
	})
}

func GetCharacterByID(c *fiber.Ctx) error {
	gameID := c.Params("gameID")
	if gameID == "" {
		return c.Status(400).JSON(fiber.Map{
			"error": "Please provide a game ID",
		})
	}

	characterID := c.Params("id")
	if characterID == "" {
		return c.Status(400).JSON(fiber.Map{
			"error": "Please provide a character ID",
		})
	}

	dbChar := model.Character{}
    err := store.DB.Where("id = ? and game_id = ?", characterID, gameID).First(&dbChar).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(404).JSON(fiber.Map{
				"error": "No character found with ID " + characterID,
			})
		} else {
			return c.Status(500).JSON(fiber.Map{
				"error":  "Internal server error",
				"reason": err.Error(),
			})
		}
	}

	err = queries.GetCharacterMoves(&dbChar)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error":  "Internal server error",
			"reason": err.Error(),
		})
	}

	movesResponse := []MoveResponse{}
	for i := range dbChar.Moves {
		movesResponse = append(movesResponse, MoveResponse{
			ID:       dbChar.Moves[i].ID,
			Name:     dbChar.Moves[i].Name,
			Input:    dbChar.Moves[i].Input,
			Startup:  dbChar.Moves[i].Startup,
			Active:   dbChar.Moves[i].Active,
			Recovery: dbChar.Moves[i].Recovery,
			OnBlock:  dbChar.Moves[i].OnBlock,
			OnHit:    dbChar.Moves[i].OnHit,
		})
	}
	return c.JSON(fiber.Map{
		"character": CharacterResponse{
			ID:     dbChar.ID,
			Name:   dbChar.Name,
			GameID: dbChar.GameID,
			Moves:  movesResponse,
		},
	})
}
