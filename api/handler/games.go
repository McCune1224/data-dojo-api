package handler

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/mccune1224/data-dojo/api/model"
	"github.com/mccune1224/data-dojo/api/store"
)

type GameResponse struct {
	ID           uint      `json:"id"`
    Name         string    `gorm:"uniqueIndex" json:"name"`
	ReleaseDate  time.Time `json:"release_date"`
	Developer    string    `json:"developer"`
	Publisher    string    `json:"publisher"`
	Abbreviation string    `json:"abbreviation"`
}

func GetAllGames(c *fiber.Ctx) error {
	DbGames := []model.Game{}
	gamesResponse := []GameResponse{}
	error := store.DB.Find(&DbGames).Error
	if error != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Could not find games",
		})
	}
	for i := range DbGames {
		gamesResponse = append(gamesResponse, GameResponse{
			ID:           DbGames[i].ID,
			Name:         DbGames[i].Name,
			Abbreviation: DbGames[i].Abbreviation,
			ReleaseDate:  DbGames[i].ReleaseDate,
			Developer:    DbGames[i].Developer,
			Publisher:    DbGames[i].Publisher,
		})
	}
	return c.JSON(fiber.Map{
		"games": gamesResponse,
	})
}

func GetGameByID(c *fiber.Ctx) error {
	gameParam := c.Params("game")
	if gameParam == "" {
		return c.Status(400).JSON(fiber.Map{
			"error": "Please provide a game ID",
		})
	}

	DbGame := model.Game{}
	idError := store.DB.First(&DbGame, gameParam).Error
	if idError != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Could not find game " + gameParam,
			"error":   idError.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"game": GameResponse{
			ID:           DbGame.ID,
			Name:         DbGame.Name,
			Abbreviation: DbGame.Abbreviation,
			ReleaseDate:  DbGame.ReleaseDate,
			Developer:    DbGame.Developer,
			Publisher:    DbGame.Publisher,
		},
	})
}
