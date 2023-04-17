package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mccune1224/data-dojo/api/handler"
)

// BasicRoutes adds routes that are not specific to a game to the app
func BasicRoutes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.SendString("OK")
	})

	app.Get("/ping", func(c *fiber.Ctx) error {
		return c.SendString("pong")
	})
}

func APIRoutes(app *fiber.App) {
	api := app.Group("/api")

	// Games
	games := api.Group("/games")
	games.Get("/", handler.GetAllGames)
	games.Get("/:game", handler.GetGameByID)

	// Characters
	characters := games.Group(":gameID/characters")
	characters.Get("/", handler.GetAllCharacters)
    characters.Get("/:id", handler.GetCharacterByID)
}
