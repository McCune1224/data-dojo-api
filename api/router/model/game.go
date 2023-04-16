package model

import (
	"gorm.io/gorm"
)

type Game struct {
	gorm.Model
	Name        string
	Description string
	ReleaseDate string
	Developer   string
	Publisher   string
	Genre       string
	Platform    string
}
