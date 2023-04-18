package handler

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/mccune1224/data-dojo/api/model"
	"github.com/mccune1224/data-dojo/api/store"
	"gorm.io/gorm"
)

type MoveResponse struct {
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

// Helper to drop gorm.Model fields from the response
func (mr *MoveResponse) ModelToResponse(m model.Move) {
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
	charID := c.Params("characterID")
	dbMoves := []model.Move{}
	err := store.DB.Where("character_id = ?", charID).Find(&dbMoves).Error
	// Return appropriate error message if there is an error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(404).JSON(fiber.Map{
				"error": "No moves found for character " + charID,
			})
		} else {
			return c.Status(500).JSON(fiber.Map{
				"error":  "Internal server error",
				"reason": err.Error(),
			})
		}
	}

	// Convert DB model to response model
	moves := []MoveResponse{}
	for i := range dbMoves {
		moveResponse := MoveResponse{}
		moveResponse.ModelToResponse(dbMoves[i])
		moves = append(moves, moveResponse)
	}

	// Return response
	return c.JSON(fiber.Map{
		"moves": moves,
	})
}

// Get single move for a given character by 'characterID' and 'id' from params
func GetMoveByID(c *fiber.Ctx) error {
	charID := c.Params("characterID")
	moveID := c.Params("id")

	// Query for move with matching characterID and moveID from params
	dbMove := model.Move{}
	err := store.DB.Where("character_id = ? AND id = ?", charID, moveID).First(&dbMove).Error
	// Handle errors if any
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(404).JSON(fiber.Map{
				"error": "No move found with id " + moveID,
			})
		} else {
			return c.Status(500).JSON(fiber.Map{
				"error":  "Internal server error",
				"reason": err.Error(),
			})
		}
	}
	if dbMove.ID == 0 {
		return c.Status(404).JSON(fiber.Map{
			"error": "No move found with id " + moveID,
		})
	}

	// Convert DB model to response model
	responseMove := MoveResponse{}
	responseMove.ModelToResponse(dbMove)

	return c.JSON(fiber.Map{
		"move": responseMove,
	})
}
