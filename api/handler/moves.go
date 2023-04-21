package handler

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/mccune1224/data-dojo/api/model"
	"github.com/mccune1224/data-dojo/api/store"
	"gorm.io/gorm"
)

// JSON response for a move
type moveResponse struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	Input     string `json:"input"`
	Startup   string `json:"startup"`
	Active    string `json:"active"`
	Recovery  string `json:"recovery"`
	OnBlock   string `json:"on_block"`
	OnHit     string `json:"on_hit"`
	OnCounter string `json:"on_counter"`
}

// Helper to convert from gorm.Model to moveResponse (basically dropping gorm.Model fields)
func (mr *moveResponse) ModelToResponse(m model.Move) {
	mr.ID = m.ID
	mr.Name = m.Name
	mr.Input = m.Input
	mr.Startup = m.Startup
	mr.Active = m.Active
	mr.Recovery = m.Recovery
	mr.OnBlock = m.OnBlock
	mr.OnHit = m.OnHit
	mr.OnCounter = m.OnCounter
}

// Get Moves for a given character by 'characterID' from params
func GetAllMoves(c *fiber.Ctx) error {
	// Query DB for all moves with characterID
	charID, err := c.ParamsInt("characterID")
	if err != nil {
		return c.Status(400).JSON(&ErrorResponse{
			Error:             Error400String,
			ErrorDescription: "please provide a character id",
		})
	}
	dbMoves := []model.Move{}
	err = store.DB.
		Where("character_id = ?", charID).
		Find(&dbMoves).Error
	// Return appropriate error message if there is an error
	if err != nil {
		return c.Status(500).JSON(Error500Response)
	}

	// Convert DB model to response model
	movesResp := []moveResponse{}
	for i := range dbMoves {
		moveResponse := moveResponse{}
		moveResponse.ModelToResponse(dbMoves[i])
		movesResp = append(movesResp, moveResponse)
	}

	// Return response
	return c.JSON(movesResp)
}

// Get single move for a given character by 'characterID' and 'id' from params
func GetMoveByID(c *fiber.Ctx) error {
	charID, err := c.ParamsInt("characterID")
	if err != nil {
		return c.Status(400).JSON(&ErrorResponse{
			Error:             Error400String,
			ErrorDescription: "please provide a character id",
		})
	}
	moveID, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON(&ErrorResponse{
			Error:             Error400String,
			ErrorDescription: "please provide a move id",
		})
	}

	// Query for move with matching characterID and moveID from params
	dbMove := model.Move{}
	err = store.DB.
		Where("character_id = ? AND id = ?", charID, moveID).
		First(&dbMove).Error
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(500).JSON(Error500Response)
		}
		return c.Status(404).JSON(&ErrorResponse{
			Error: fmt.Sprintf("No move found with id %d", moveID),
		})
	}

	// Convert DB model to response model
	moveResp := moveResponse{}
	moveResp.ModelToResponse(dbMove)

	return c.JSON(moveResp)
}

func SearchMoves(c *fiber.Ctx) error {
	// Get query params
	requestQuery := c.Query("name")
	limitQuery := c.Query("limit")
	limitQueryInt, err := strconv.Atoi(limitQuery)
	if err != nil {
		limitQueryInt = 10
	}

	if requestQuery == "" {
		return c.Status(400).JSON(fiber.Map{
			"error": "No search query provided",
		})
	}
	if len(requestQuery) > 20 {
		return c.Status(400).JSON(fiber.Map{
			"error": "Search query too long (max 20 characters)",
		})
	}

	// Query DB for all moves with characterID
	dbMoves := []model.Move{}
	err = store.DB.
		Where("name ILIKE ?", "%"+requestQuery+"%").
		Or("input ILIKE ?", "%"+requestQuery+"%").
		Find(&dbMoves).
		Limit(limitQueryInt).
		Error
	if err != nil {
		return c.Status(500).JSON(Error500Response)
	}

	// Make response JSON and return
	movesResp := []moveResponse{}
	for i := range dbMoves {
		newMoveResponse := moveResponse{}
		newMoveResponse.ModelToResponse(dbMoves[i])
		movesResp = append(movesResp, newMoveResponse)
	}
	return c.JSON(movesResp)
}
