package internal

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/mccune1224/data-dojo/api/models"
)

func TestStriveScraper(t *testing.T) {
	testQueries := []string{
		"Ky Kiske",
		"Sol Badguy",
		"May",
		"Millia Rage",
		"Potemkin",
		"Zato-1",
	}
	for _, query := range testQueries {
		queryResult, err := StriveScraper(query)
		if err != nil {
			t.Error(err)
		}
		if len(queryResult.Cargoquery) == 0 {
			t.Error("queryResult is empty")
		}
		newCharacter := models.Character{
			Name: query,
		}
		for _, move := range queryResult.Cargoquery {
			newMove := models.Move{
				Name:     move.Title.Name,
				Input:    move.Title.Input,
				Startup:  move.Title.Startup,
				Active:   move.Title.Active,
				Recovery: move.Title.Recovery,
				OnBlock:  move.Title.OnBlock,
				OnHit:    move.Title.OnHit,
			}
			newCharacter.Moves = append(newCharacter.Moves, newMove)
		}

		// write to file
		newFile, err := os.Create(fmt.Sprintf("%s.json", query))
		if err != nil {
			log.Fatal(err)
		}
		defer newFile.Close()
		json.NewEncoder(newFile).Encode(newCharacter)

	}
}
