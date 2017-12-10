package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

type GameInfo struct {
	GameID            int    `json:"gameId"`
	GameMode          string `json:"gameMode"`
	GameType          string `json:"gameType"`
	GameQueueConfigID int    `json:"gameQueueConfigId"`
	Participants      []struct {
		TeamID       int    `json:"teamId"`
		SummonerName string `json:"summonerName"`
		SummonerID   int    `json:"summonerId"`
	} `json:"participants"`
	Observers struct {
		EncryptionKey string `json:"encryptionKey"`
	} `json:"observers"`
	PlatformID    string `json:"platformId"`
	GameStartTime int64  `json:"gameStartTime"`
	GameLength    int    `json:"gameLength"`
}

// criar variaveis:
// contador de tempo
// token da API
// contador (comeca com 5 min)

// pegar challengers de um server
func getServerTopPlayers(server string) (players *Challenger, err error) {
	requestURL := fmt.Sprintf("https://%s.api.riotgames.com/lol/league/v3/challengerleagues/by-queue/RANKED_SOLO_5x5?api_key=%s", server, APIKey)
	resp, err := http.Get(requestURL)
	if err != nil {
		return
	}
	defer resp.Body.Close()
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
	if err != nil {
		return
	}
	defer resp.Body.Close()
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
	wd, _ := os.Getwd()
	defer os.Chdir(wd)
	return launchGame(server, game)
}

// cmd := fmt.Sprintf("cd \"C:\\Riot Games\\League of Legends\\RADS\\solutions\\lol_game_client_sln\\releases\\\" & for /d %%F in (*) do cd %%F & start \"\" /D \"deploy\" \"League of Legends.exe\" \"8394\" \"LoLLauncher.exe\" \"\" \"replay spectator.%s.lol.riotgames.com:80 %v %v BR1\"", server, game.Observers.EncryptionKey, game.GameID)
func launchGame(server string, game *GameInfo) (err error) {
	path, err := os.Open("C:\\Riot Games\\League of Legends\\RADS\\solutions\\lol_game_client_sln\\releases\\")
	if err != nil {
		return
	}
	r, err := findLatestRelease(path)
	if err != nil {
		return
	}
	deployPath := fmt.Sprintf("C:\\Riot Games\\League of Legends\\RADS\\solutions\\lol_game_client_sln\\releases\\%s\\deploy", r)
	err = os.Chdir(deployPath)
	if err != nil {
		return
	}
	// ABRIR JOGO
	return
}

func findLatestRelease(path *os.File) (release string, err error) {
	files, err := path.Readdir(0)
	if err != nil {
		return "", err
	}
	var t fileWithTime
	for _, f := range files {
		if checkAfter(t.time, f.ModTime()) && f.IsDir() {
			t.time = f.ModTime()
			t.fileName = f.Name()
		}
	}
	return t.fileName, err
}

func checkAfter(start, check time.Time) (b bool) {
	return check.After(start)
}

// verificar se jogo terminou
func getGameStatus() {
	return
}
