package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mccune1224/data-dojo/api/handler"
)

// BasicRoutes adds routes that are not specific to a game to the app
func BasicRoutes(app *fiber.App) {
	// Root route
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	// Helper for checking if the server is up
	app.Get("/ping", func(c *fiber.Ctx) error {
		return c.SendString("pong")
	})
}

func APIRoutes(app *fiber.App) {
	api := app.Group("/api")

	// Helper for showing all routes
	api.Get("/routes", func(c *fiber.Ctx) error {
		return c.JSON(app.GetRoutes(false))
	})
	// Games
	games := api.Group("/games")
	games.Get("/search", handler.SearchGames)
	games.Get("/:id", handler.GetGameByID)
	games.Get("/", handler.GetAllGames)

	// Characters
	characters := games.Group(":gameID/characters")
	characters.Get("/search", handler.SearchCharacters)
	characters.Get("/:id", handler.GetCharacterByID)
	characters.Get("/", handler.GetAllCharacters)

	// Moves
	moves := characters.Group(":characterID/moves")
	moves.Get("/search", handler.SearchMoves)
	moves.Get("/:id", handler.GetMoveByID)
	moves.Get("/", handler.GetAllMoves)
}
