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
	var bothGames []Game
	var windowsGames []Game
	for id, game := range gamesMap {
		game.ID = id
		games = append(games, game)

		if game.Mac && !game.Linux && !game.Windows {
			macGames = append(macGames, game)
		}
		if game.Linux && !game.Mac && !game.Windows {
			linuxGames = append(linuxGames, game)
		}
		if game.Linux && game.Mac && !game.Windows {
			bothGames = append(bothGames, game)
		}
		if game.Windows && !game.Linux && !game.Mac {
			windowsGames = append(windowsGames, game)
		}

		fmt.Printf("Processed game ^-^: %s (ID: %s)\n", game.Name, game.ID)
	}

	fmt.Printf("\nProcessed %d games total\n", len(games))
	fmt.Printf("\nProcessed %d mac exclusive games\n", len(macGames))
	for id, game := range macGames {
		fmt.Printf("MacOS %d: %s %s\n", id, game.ID, game.Name)
	}
	fmt.Printf("\nProcessed %d linux exclusive games\n", len(linuxGames))
	for id, game := range linuxGames {
		fmt.Printf("Linux %d: %s %s\n", id, game.ID, game.Name)
	}
	fmt.Printf("\nProcessed %d games that are only on both linux and mac\n", len(bothGames))
	for id, game := range bothGames {
		fmt.Printf("Both %d: %s %s\n", id, game.ID, game.Name)
	}
	fmt.Printf("\nProcessed %d windows exclusive games\n", len(windowsGames))
}
