package main

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