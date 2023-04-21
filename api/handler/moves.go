package handler

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/mccune1224/data-dojo/api/model"
	"github.com/mccune1224/data-dojo/api/store"
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
	charID := c.Params("characterID")
	dbMoves := []model.Move{}
	err := store.DB.
		Where("character_id = ?", charID).
		Find(&dbMoves).Error
	// Return appropriate error message if there is an error
	if err != nil {
		return handleNotFound(c, err)
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
	charID := c.Params("characterID")
	moveID := c.Params("id")

	// Query for move with matching characterID and moveID from params
	dbMove := model.Move{}
	err := store.DB.
		Where("character_id = ? AND id = ?", charID, moveID).
		First(&dbMove).Error
	// Handle errors if any
	if err != nil {
		handleNotFound(c, err)
	}
	if dbMove.ID == 0 {
		return c.Status(404).JSON(fiber.Map{
			"error": "No move found with id " + moveID,
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
		handleNotFound(c, err)
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
