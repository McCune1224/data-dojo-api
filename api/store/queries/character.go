package queries

import (
	"github.com/mccune1224/data-dojo/api/model"
	"github.com/mccune1224/data-dojo/api/store"
)

func AllCharactersByID(gameID string, characters *[]model.Character) error {
	// Get them by ID
	err := store.DB.Where("game_id = ?", gameID).Find(&characters).Error
	if err != nil {
		return err
	}
	return nil
}

func CharacterByName(name string, character *model.Character) error {
	// Get them by ID
	err := store.DB.Where("name = ?", name).Find(&character).Error
	if err != nil {
		return err
	}
	return nil
}

func CharacterByID(id string, character *model.Character) error {
	// Get them by ID
	err := store.DB.Where("id = ?", id).Find(&character).Error
	if err != nil {
		return err
	}
	// Get their moves
	err = store.DB.Model(&character).Association("Moves").Find(&character.Moves)
	if err != nil {
		return err
	}
	return nil
}

func CharacterMoveByID(id string, character *model.Character) error {
	// Get their moves
	err := store.DB.Model(&character).Association("Moves").Find(&character.Moves)
	if err != nil {
		return err
	}
	return nil
}

func GetCharacterMoves(character *model.Character) error {
	// Get their moves
	err := store.DB.Model(&character).Association("Moves").Find(&character.Moves)
	if err != nil {
		return err
	}
	return nil
}
