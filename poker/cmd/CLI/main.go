package main

import (
	"log"
	"os"

	"github.com/RyGosling/learn-github/poker"
)

const dbFileName = "game.db.json"

func main() {
	store, err := poker.FileSystemStoreFromFile(dbFileName)
	if err != nil {
		log.Fatal(err)
	}
	game := poker.NewCLI(store, os.Stdin)
	game.PlayPoker()
}
