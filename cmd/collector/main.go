package main

import (
	"flag"
	"os"

	_ "github.com/joho/godotenv/autoload"
	"github.com/mccune1224/data-dojo/collector"
)

type GameScraper interface {
	Fetch(searchTerm string) error
	Write(dbURL string) error
}

func gameQuery(gs GameScraper, searchTerm string, dsn string) {
	// Fetch data from the game
	err := gs.Fetch(searchTerm)
	if err != nil {
		panic(err)
	}

	// Write data to db
	err = gs.Write(dsn)
	if err != nil {
		panic(err)
	}
}

func main() {
	// Parse flag args
	game := flag.String("game", "GGST", "The game to scrape data from")
	searchTerm := flag.String("search", "Ky Kiske", "The search term to use")
	flag.Parse()

	DB_URL := os.Getenv("DB_URL")
	if *game == "GGST" {
		ggstScraper := collector.GuiltyGearStriveQuery{}
		gameQuery(&ggstScraper, *searchTerm, DB_URL)
	}
}
