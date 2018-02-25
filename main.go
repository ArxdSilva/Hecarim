package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"time"
)

var APIKey string

type fileWithTime struct {
	fileName string
	time     time.Time
}

func main() {
	APIKey = os.Getenv("APIKEY")
	fmt.Println("APIKEY: ", APIKey)
	if APIKey == "" {
		log.Fatal(errors.New("APIKEY not set"))
	}
	players, err := getServerTopPlayers("br")
	if err != nil {
		log.Fatal(err)
	}
	playerIds := getIdsFromPlayers(players)
	for _, pID := range playerIds {
		fmt.Print(".")
		game, online, err := getPlayerGameAndStatus("br", pID)
		if err != nil {
			log.Fatal(err)
		}
		if !online {
			continue
		}
		// fmt.Printf("\n%+v\n\n", game)
		errGame := openGame("br", game)
		if errGame != nil {
			log.Fatal(errGame)
		}
		time.Sleep(1 * time.Second)
	}
}
