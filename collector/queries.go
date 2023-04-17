package collector

import "github.com/mccune1224/data-dojo/api/models"

type DustloopQuery struct {
	Cargoquery []struct {
		// DustLoop queries are weird, every result is put inside a "title" object...
		// But really these are just the moves for a character
		Title models.Move `json:"title"`
	} `json:"cargoquery"`
}
