package model

import (
	"time"

	"gorm.io/gorm"
)

// Base model for all games
type Game struct {
	gorm.Model
	Name         string `gorm:"uniqueIndex"`
	Abbreviation string `gorm:"uniqueIndex"`
	ReleaseDate  time.Time
	Developer    string
	Publisher    string

	Characters []Character
}

// Character belongs to a Game and has many Moves
type Character struct {
	gorm.Model
	Name  string `gorm:"uniqueIndex"`
	Moves []Move

	Game   Game
	GameID uint
}

// Move belongs to a Character
type Move struct {
	gorm.Model
	Name      string
	Input     string
	Startup   string
	Active    string
	Recovery  string
	OnBlock   string
	OnHit     string
	OnCounter string

	// Relationships
	Character   Character
	CharacterID uint
}
