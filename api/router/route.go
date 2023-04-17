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

    app.Get("/test", func(c *fiber.Ctx) error {
        return c.SendString("test")
    })
}

func GameRoutes(app *fiber.App) {
	games := app.Group("/games")

	// Get all games in the database
	games.Get("/", handler.GetAllGames)
}
