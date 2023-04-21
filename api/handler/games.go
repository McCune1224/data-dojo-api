package handler

import (
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/mccune1224/data-dojo/api/model"
	"github.com/mccune1224/data-dojo/api/store"
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
	DbGames := []model.Game{}
	gamesResponse := []gameResponse{}
	error := store.DB.Find(&DbGames).Error
	if error != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Could not find games",
		})
	}
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
	gameParam := c.Params("id")

	if gameParam == "" {
		return c.Status(400).JSON(&ErrorResponse{
			Error:             "bad_request",
			Error_Description: "please provide a game id",
		})
	}

	DbGame := model.Game{}
	dbErr := store.DB.First(&DbGame, gameParam).Error
	handleNotFound(c, dbErr)

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
		return c.Status(400).JSON(
			&ErrorResponse{
				Error:             "bad_request",
				Error_Description: "please provide a 'name' query",
			})
	}
	if len(requestQuery) > 50 {
		return c.Status(400).JSON(fiber.Map{
			"error": "query has limit of 50 characters",
		})
	}
	dbResults := []model.Game{}
	// Can be used to find a game by name or abbreviation
	// May god have mercy on your soul if you have to debug this (I prob will have to)
	err = store.DB.
		Where("name ILIKE ?", "%"+requestQuery+"%").
		Or("abbreviation ILIKE ?", "%"+requestQuery+"%").
		Find(&dbResults).
		Limit(limitQueryInt).
		Error
	if err != nil {
		handleNotFound(c, err)
	}
	if dbResults == nil {
		return c.Status(404).JSON(fiber.Map{
			"error": "Could not find any matching",
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
