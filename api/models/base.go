package models

import "gorm.io/gorm"

// Base model for all games
type Game struct {
	gorm.Model
	Name        string
	ReleaseDate string
	Developer   string
	Publisher   string
	Genre       string
	Platform    string

	// Relationships
	Characters []Character
}

// Base model for all characters
type Character struct {
	gorm.Model
	Name  string
	Moves []Move
}

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
	CharacterID uint
}
