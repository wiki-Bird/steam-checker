package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Game struct {
	ID      string `json:"-"`
	Name    string `json:"name"`
	Windows bool   `json:"windows"`
	Mac     bool   `json:"mac"`
	Linux   bool   `json:"linux"`
}

func main() {
	fmt.Println("starting to calculate :3c...")

	file, err := os.Open("games.json")
	if err != nil {
		fmt.Println("Error opening games file :< ", err)
		return
	}
	defer file.Close()
	fmt.Println("File opened successfully")

	var gamesMap map[string]Game
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&gamesMap)

	if err != nil {
		fmt.Println("Error decoding JSON :<", err)
		return
	}
	fmt.Printf("Found %d games in the JSON c: \n", len(gamesMap))

	var games []Game
	var macGames []Game
	var linuxGames []Game
	for id, game := range gamesMap {
		game.ID = id
		games = append(games, game)

		if game.Mac && !game.Linux && !game.Windows {
			macGames = append(macGames, game)
		}
		if game.Linux && !game.Mac && !game.Windows {
			linuxGames = append(linuxGames, game)
		}

		fmt.Printf("Processed game ^-^: %s (ID: %s)\n", game.Name, game.ID)
	}

	fmt.Printf("\nProcessed %d games total\n", len(games))
	fmt.Printf("Processed %d mac exclusive games\n", len(macGames))
	fmt.Printf("Processed %d linux exclusive games\n", len(linuxGames))
}
