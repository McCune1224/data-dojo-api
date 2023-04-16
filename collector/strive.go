package internal

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
)

type DustloopQuery struct {
	Cargoquery []struct {
		Title struct {
			Chara    string `json:"chara"`
			Input    string `json:"input"`
			Name     string `json:"name"`
			Startup  string `json:"startup"`
			Active   string `json:"active"`
			Recovery string `json:"recovery"`
			OnBlock  string `json:"onBlock"`
			OnHit    string `json:"onHit"`
		} `json:"title"`
	} `json:"cargoquery"`
}

func StriveScraper(characterName string) (*DustloopQuery, error) {
	// Replace whitespaces with %20 for the URL

	characterName = strings.ReplaceAll(characterName, " ", "%20")
	quotedCharacterName := fmt.Sprintf("\"%s\"", characterName)
	baseURL := "https://dustloop.com/wiki/api.php?action=cargoquery&tables=MoveData_GGST&fields=chara,input,name,startup,active,recovery,onBlock,onHit,&where=chara=%s&format=json"
	reqURL := fmt.Sprintf(baseURL, quotedCharacterName)
	client := &http.Client{}
	req, err := http.NewRequest("GET", reqURL, nil)
	req.Header.Add("Content-Type", "application/json")
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, errors.New(fmt.Sprintf("Wanted 200 status, got %d", resp.StatusCode))
	}
	defer resp.Body.Close()
	var dustloopQuery DustloopQuery
	err = json.NewDecoder(resp.Body).Decode(&dustloopQuery)
	if err != nil {
		return nil, err
	}

	return &dustloopQuery, nil
}
