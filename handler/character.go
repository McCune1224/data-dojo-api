package handler

import (
	"gorm.io/gorm"
)

type Character struct {
	gorm.Model
	Name        string
	Description string
	MoveSet     string
}
