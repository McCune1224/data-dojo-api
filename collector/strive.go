package collector

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/url"

	"github.com/mccune1224/data-dojo/api/models"
	"github.com/mccune1224/data-dojo/api/store"
)

type GuiltyGearStriveQuery struct {
	Character string        `json:"character"`
	Moves     []models.Move `json:"moves"`
}

// Fetch the moves for a character from DustLoop
func (ggstq *GuiltyGearStriveQuery) Fetch(characterName string) error {
	encodedQuery := url.QueryEscape(characterName)
	quotedEncodedName := fmt.Sprintf("\"%s\"", encodedQuery)
	baseURL := "https://dustloop.com/wiki/api.php?action=cargoquery&tables=MoveData_GGST&fields=chara,input,name,startup,active,recovery,onBlock,onHit,&where=chara=%s&format=json"
	reqURL := fmt.Sprintf(baseURL, quotedEncodedName)
	client := &http.Client{}
	req, err := http.NewRequest("GET", reqURL, nil)
	req.Header.Add("Content-Type", "application/json")
	if err != nil {
		return err
	}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return errors.New(fmt.Sprintf("Wanted 200 status, got %d", resp.StatusCode))
	}

	var dustloopQuery DustloopQuery
	err = json.NewDecoder(resp.Body).Decode(&dustloopQuery)
	if err != nil {
		return err
	}

	if len(dustloopQuery.Cargoquery) == 0 {
		return errors.New("No results found")
	}

	if dustloopQuery.Cargoquery[0].Title.Input == "" {
		return errors.New("No results found")
	}

	ggstq.Character = characterName
	for _, currMove := range dustloopQuery.Cargoquery {
		ggstq.Moves = append(ggstq.Moves, currMove.Title)
	}

	return nil
}

func (ggstq *GuiltyGearStriveQuery) Write(dsn string) error {
	if len(ggstq.Moves) == 0 {
		return errors.New("No moves to write")
	}
	log.Println(ggstq.Moves)
	if ggstq.Character == "" {
		return errors.New("No Character name provided")
	}
	err := store.Connect(dsn, true)
	if err != nil {
		log.Fatal(err)
	}

	// // check if game exists first
	// ggst := models.Game{
	// 	Name:        "Guilty Gear Strive",
	// 	Developer:   "Arc System Works",
	// 	Publisher:   "Bandai Namco Entertainment",
	// 	ReleaseDate: time.Date(2021, time.June, 8, 0, 0, 0, 0, time.UTC),
	// }

	// gameCreate := store.DB.Create(&ggst)
	// if gameCreate.Error != nil {
	// 	return gameCreate.Error
	// }

	game := models.Game{}
	err = store.DB.Where("name = ?", "Guilty Gear Strive").First(&game).Error
	if err != nil {
		return err
	}

	// insert character into db
	characterEntry := models.Character{
		Name:   ggstq.Character,
		Moves:  ggstq.Moves,
		GameID: game.ID,
	}
	err = store.DB.Create(&characterEntry).Error
	if err != nil {
		return err
	}

	fmt.Println(characterEntry)

	return nil
}
