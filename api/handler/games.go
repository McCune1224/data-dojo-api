package handler

import (
	"errors"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/mccune1224/data-dojo/api/model"
	"github.com/mccune1224/data-dojo/api/store"
	"gorm.io/gorm"
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
	gameParam := c.Params("id")

	if gameParam == "" {
		return c.Status(400).JSON(fiber.Map{
			"error": "Please provide a game ID",
		})
	}

	DbGame := model.Game{}
	dbErr := store.DB.First(&DbGame, gameParam).Error
	if dbErr != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Could not find game " + gameParam,
			"error":   dbErr.Error(),
			"param":   gameParam,
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

func SearchGames(c *fiber.Ctx) error {
	requestQuery := c.Params("query")
	if requestQuery == "" {
		return c.Status(400).JSON(fiber.Map{
			"error": "Please provide a search term",
		})
	}
	dbResults := []model.Game{}
	// Can be used to find a game by name or abbreviation
	// May god have mercy on your soul if you have to debug this (I prob will have to)
	err := store.DB.
		Where("name ILIKE ?", "%"+requestQuery+"%").
		Or("abbreviation ILIKE ?", "%"+requestQuery+"%").
		Find(&dbResults).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(404).JSON(fiber.Map{})
		}
		return c.Status(500).JSON(fiber.Map{
			"error": "Could not search games",
		})
	}
	if dbResults == nil {
		return c.Status(404).JSON(fiber.Map{
			"error": "Could not find any matching",
		})
	}
	gamesResponse := []GameResponse{}
	for i := range dbResults {
		gamesResponse = append(gamesResponse, GameResponse{
			ID:           dbResults[i].ID,
			Name:         dbResults[i].Name,
			Abbreviation: dbResults[i].Abbreviation,
			ReleaseDate:  dbResults[i].ReleaseDate,
			Developer:    dbResults[i].Developer,
			Publisher:    dbResults[i].Publisher,
		})
	}
	return c.JSON(fiber.Map{
		"games": gamesResponse,
	})

}
