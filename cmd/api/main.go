package main

import (
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	_ "github.com/joho/godotenv/autoload"
	"github.com/mccune1224/data-dojo/api/router"
	"github.com/mccune1224/data-dojo/api/store")

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	return ":" + port
}


func main() {
	// Spin up DB
	dbConnectErr := store.Connect(os.Getenv("DATABASE_URL"), false)
	if dbConnectErr != nil {
		panic(dbConnectErr)
	}

	app := fiber.New()
	app.Use(logger.New())

	app.Use(limiter.New(limiter.Config{
		Max:               40,
		Expiration:        30 * time.Second,
		LimiterMiddleware: limiter.SlidingWindow{},
	}))

	// Add routes to app
	router.BasicRoutes(app)
	router.APIRoutes(app)

	app.Listen(getPort())
}
