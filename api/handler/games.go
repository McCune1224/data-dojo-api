package handler

import (
	"errors"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/mccune1224/data-dojo/api/model"
	"github.com/mccune1224/data-dojo/api/store"
	"gorm.io/gorm"
)

// JSON response for a game
type gameResponse struct {
	ID           uint      `json:"id"`
	Name         string    `gorm:"uniqueIndex" json:"name"`
	ReleaseDate  time.Time `json:"release_date"`
	Developer    string    `json:"developer"`
	Publisher    string    `json:"publisher"`
	Abbreviation string    `json:"abbreviation"`
}

func GetAllGames(c *fiber.Ctx) error {
	// Query the database for all games
	DbGames := []model.Game{}
	gamesResponse := []gameResponse{}
	error := store.DB.Find(&DbGames).Error
	if error != nil {
		return c.Status(500).JSON(Error500Response)
	}

	// Create and return games response
	for i := range DbGames {
		gamesResponse = append(gamesResponse, gameResponse{
			ID:           DbGames[i].ID,
			Name:         DbGames[i].Name,
			Abbreviation: DbGames[i].Abbreviation,
			ReleaseDate:  DbGames[i].ReleaseDate,
			Developer:    DbGames[i].Developer,
			Publisher:    DbGames[i].Publisher,
		})
	}
	return c.JSON(gamesResponse)
}

func GetGameByID(c *fiber.Ctx) error {
	// Get game ID from query
	gameParam, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON(&ErrorResponse{
			Error:            "bad_request",
			ErrorDescription: "please provide a game id",
		})
	}

	// Query the database for the game
	DbGame := model.Game{}
	dbErr := store.DB.Where("id = ?", gameParam).First(&DbGame).Error
	if dbErr != nil && !errors.Is(dbErr, gorm.ErrRecordNotFound) {
		return c.Status(500).JSON(Error500Response)
	} else if errors.Is(dbErr, gorm.ErrRecordNotFound) {
		return c.Status(404).JSON(&ErrorResponse{
			Error:            Error404String,
			ErrorDescription: "Could not find game",
		})
	}

	// Create and return game response
	return c.JSON(
		gameResponse{
			ID:           DbGame.ID,
			Name:         DbGame.Name,
			Abbreviation: DbGame.Abbreviation,
			ReleaseDate:  DbGame.ReleaseDate,
			Developer:    DbGame.Developer,
			Publisher:    DbGame.Publisher,
		})
}

// SearchGames will expect a query paramter called "game" to
// use in an ILIKE search for games of similar pattern
func SearchGames(c *fiber.Ctx) error {
	requestQuery := c.Query("name")
	limitQuery := c.Query("limit")
	limitQueryInt, err := strconv.Atoi(limitQuery)
	if err != nil {
		limitQueryInt = 10
	}

	if requestQuery == "" {
		return c.Status(400).JSON(&ErrorResponse{
			Error:            "bad_request",
			ErrorDescription: "please provide a 'name' query",
		})
	}
	if len(requestQuery) > 30 {
		return c.Status(400).JSON(&ErrorResponse{
			Error:            Error400String,
			ErrorDescription: "query has to be less than 30 characters",
		})
	}
	dbResults := []model.Game{}
	// Can be used to find a game by name or abbreviation
	// May god have mercy on your soul if you have to debug this (I prob will have to)
	err = store.DB.
		Where("name ILIKE ?", "%"+requestQuery+"%").
		Order("name").
		Or("abbreviation ILIKE ?", "%"+requestQuery+"%").
		Limit(limitQueryInt).
		Find(&dbResults).
		Error
	if err != nil {
		return c.Status(500).JSON(Error500Response)
	}
	if dbResults == nil {
		return c.Status(404).JSON(&ErrorResponse{
			Error:            Error404String,
			ErrorDescription: "Could not find games of similar name",
		})
	}
	gamesResponse := []gameResponse{}
	for i := range dbResults {
		gamesResponse = append(gamesResponse, gameResponse{
			ID:           dbResults[i].ID,
			Name:         dbResults[i].Name,
			Abbreviation: dbResults[i].Abbreviation,
			ReleaseDate:  dbResults[i].ReleaseDate,
			Developer:    dbResults[i].Developer,
			Publisher:    dbResults[i].Publisher,
		})
	}
	return c.JSON(gamesResponse)
}
