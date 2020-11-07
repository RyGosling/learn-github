package main

import (
	"log"
	"net/http"

	"github.com/RyGosling/learn-github/poker"
)

const dbFileName = "game.db.json"

func main() {
	store, err := poker.FileSystemStoreFromFile(dbFileName)
	if err != nil {
		log.Fatal(err)
	}
	server := poker.NewPlayerServer(store)
	if err := http.ListenAndServe(":5000", server); err != nil {
		log.Fatalf("could not listen on port 5000 %v", err)
	}
}
