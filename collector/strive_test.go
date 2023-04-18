package collector

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/mccune1224/data-dojo/api/model"
)

func TestStriveQuerySerach(t *testing.T) {
	testQueries := []struct {
		Query               string
		ExptectedStatusCode int
		MoveListResponse    []model.Move
	}{
		{
			Query:               "Sol Badguy",
			ExptectedStatusCode: 200,
		},
		{
			Query:               "Happy Chaos",
			ExptectedStatusCode: 200,
		},
		{
			Query:               "Sin Kiske",
			ExptectedStatusCode: 200,
		},
	}
	for _, test := range testQueries {
		var striveQuery GuiltyGearStriveQuery
		err := striveQuery.Fetch(test.Query)
		if err != nil {
			t.Error(err)
		}

		if len(striveQuery.Moves) == 0 {
			t.Error("No moves found")
		}

	}
}

func TestURL(t *testing.T) {
	url := "https://dustloop.com/wiki/api.php?action=cargoquery&tables=MoveData_GGST&fields=chara,input,name,startup,active,recovery,onBlock,onHit,&where=chara=%22I-No%22&format=json"
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("Content-Type", "application/json")
	if err != nil {
		t.Error(err)
	}
	resp, err := client.Do(req)
	if err != nil {
		t.Error(err)
	}
	defer resp.Body.Close()

	var dustloopQuery DustloopQuery
	err = json.NewDecoder(resp.Body).Decode(&dustloopQuery)
	if err != nil {
		t.Fatal(err)
	}
}
