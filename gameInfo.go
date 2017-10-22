package main

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
