package main

import (
	"log"
	"os"
	"test/learn-github/poker"
)

const dbFileName = "game.db.json"

func main() {
	store, err := poker.FileSystemStoreFromFile(dbFileName)
	if err != nil {
		log.Fatalf(err)
	}
	game := poker.NewCLI(store, os.Stdin)
	game.PlayPoker()
}
