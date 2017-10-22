package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
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

var ChallNames []string

func main() {
	APIKey = os.Getenv("APIKEY")
	fmt.Println("APIKEY: ", APIKey)
	if APIKey == "" {
		log.Fatal(errors.New("APIKEY not set"))
	}
	players, err := getServerTopPlayers("br1")
	if err != nil {
		log.Fatal(err)
	}
	playerIds := getIdsFromLeagueInfo(players)
	for _, pId := range playerIds {
		fmt.Print(".")
		game, online, err := getPlayerGameAndStatus("br1", pId)
		if err != nil {
			log.Fatal(err)
		}
		if !online {
			continue
		}
		fmt.Println(game)
		errGame := openGame("br1", game)
		if errGame != nil {
			log.Fatal(errGame)
		}
	}
}

// criar variaveis:
// contador de tempo
// token da API
// contador (comeca com 5 min)

// pegar challengers de um server
func getServerTopPlayers(server string) (players *Challenger, err error) {
	requestURL := fmt.Sprintf("https://%s.api.riotgames.com/lol/league/v3/challengerleagues/by-queue/RANKED_SOLO_5x5?api_key=%s", server, APIKey)
	resp, err := http.Get(requestURL)
	defer resp.Body.Close()
	if err != nil {
		return
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	err = json.Unmarshal(body, &players)
	if err != nil {
		return
	}
	return
}

func getIdsFromLeagueInfo(players *Challenger) (playersSlice []string) {
	for _, p := range players.Entries {
		playersSlice = append(playersSlice, p.PlayerOrTeamID)
	}
	return
}

// pegar status de um jogador
func getPlayerGameAndStatus(server, playerId string) (game *GameInfo, status bool, err error) {
	requestURL := fmt.Sprintf("https://%s.api.riotgames.com/lol/spectator/v3/active-games/by-summoner/%s?api_key=%s", server, playerId, APIKey)
	resp, err := http.Get(requestURL)
	defer resp.Body.Close()
	if err != nil {
		return
	}
	if resp.StatusCode != http.StatusOK {
		return
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	err = json.Unmarshal(body, &game)
	if err != nil {
		return
	}
	status = true
	return
}

// abrir jogo
// a cada X segundos verificar o status do jogo
// usando a variavel do contador e depois a cada 30s
func openGame(server string, game *GameInfo) (err error) {
	cmd := fmt.Sprintf("cd \"C:\\Riot Games\\League of Legends\\RADS\\solutions\\lol_game_client_sln\\releases\\\" & for /d %%F in (*) do cd %%F & start \"\" /D \"deploy\" \"League of Legends.exe\" \"8394\" \"LoLLauncher.exe\" \"\" \"replay spectator.%s.lol.riotgames.com:80 %v %v BR1\"", server, game.Observers.EncryptionKey, game.GameID)
	return exec.Command(cmd).Run()
}

// verificar se jogo terminou
func getGameStatus() {

	return
}
