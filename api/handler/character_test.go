package handler

import (
	"encoding/json"
	"fmt"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"
	"github.com/mccune1224/data-dojo/api/store"
)

func testApplication() *fiber.App {
	// Load .env file
	godotenv.Load("../../cmd/api/.env", "./.env")
	app := fiber.New()
	// Spin up DB
	dbConnectErr := store.Connect(os.Getenv("DB_URL"), false)
	if dbConnectErr != nil {
		panic(dbConnectErr)
	}
	return app
}

// TestGetAllCharacters tests the GetAllCharacters handler
// Using GGST as the test game (Game ID 2)
func TestGetCharacterByID(t *testing.T) {
	app := testApplication()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
	app.Get("/api/games/:gameID/characters/:id", GetCharacterByID)

	testCases := []struct {
		Description    string
		GameID         string
		characterID    string
		ExpectedStatus int
	}{
		{
			Description: "Valid CharacterID",
			//GGST
			GameID: "2",
			//Zato-1
			characterID:    "11",
			ExpectedStatus: 200,
		}, {
			Description: "Invalid/Non-Existent CharacterID",
			//GGST
			GameID: "2",
			//Non-Existent Character
			characterID:    "42069",
			ExpectedStatus: 404,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.Description, func(t *testing.T) {
			reqUrl := fmt.Sprintf("/api/games/%s/characters/%s", testCase.GameID, testCase.characterID)
			req := httptest.NewRequest(
				"GET",
				reqUrl,
				nil,
			)
			res, err := app.Test(req)
			jsonBody := make(map[string]interface{})
			json.NewDecoder(res.Body).Decode(&jsonBody)
			if err != nil {
				t.Fatal(err)
			}
			if res.StatusCode != testCase.ExpectedStatus {
				t.Fatalf("Expected status code %d, got %d\nReq URL: %s", testCase.ExpectedStatus, res.StatusCode, reqUrl)
			}
			if res.StatusCode == 200 {
				resCharName := jsonBody["character"].(map[string]interface{})["name"].(string)
				if jsonBody["character"] == nil {
					t.Fatalf("Expected a data object, got nil")
				}
				if resCharName != "Zato-1" {
					t.Fatalf("Expected character name to be Zato-1, got %s", resCharName)
				}
			}
		})
	}

}
