package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"time"
)

var APIKey string

type Challenger struct {
	Name     string `json:"name"`
	Tier     string `json:"tier"`
	Queue    string `json:"queue"`
	LeagueID string `json:"leagueId"`
	Entries  []struct {
		PlayerOrTeamID   string `json:"playerOrTeamId"`
		PlayerOrTeamName string `json:"playerOrTeamName"`
		LeaguePoints     int    `json:"leaguePoints"`
		Rank             string `json:"rank"`
		Wins             int    `json:"wins"`
		Losses           int    `json:"losses"`
		Inactive         bool   `json:"inactive"`
	} `json:"entries"`
}

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
	players, err := getServerTopPlayers("ru")
	if err != nil {
		log.Fatal(err)
	}
	playerIds := getIdsFromLeagueInfo(players)
	for _, pId := range playerIds {
		fmt.Print(".")
		game, online, err := getPlayerGameAndStatus("ru", pId)
		if err != nil {
			log.Fatal(err)
		}
		if !online {
			continue
		}
		// fmt.Printf("\n%+v\n\n", game)
		errGame := openGame("ru", game)
		if errGame != nil {
			log.Fatal(errGame)
		}
		time.Sleep(1 * time.Second)
	}
}
