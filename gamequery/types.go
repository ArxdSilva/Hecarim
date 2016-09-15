package gamequery

import "fmt"

// Region ...
var Region = []struct {
	country, short string
}{
	{"brazil", "br"},
	{"europe_north_east", "eune"},
	{"europe_west", "euw"},
	{"japan", "jp"},
	{"korea", "kr"},
	{"latin_america_north", "lan"},
	{"latin_america_south", "las"},
	{"north_america", "na"},
	{"oceania", "oce"},
	{"pbe", "pbe"},
	{"russia", "ru"},
	{"turkey", "tr"},
}

var gameMode = []struct {
	name, short string
}{
	{"aram", "ARAM"},
	{"ascension", "ASCENSION"},
	{"classic", "CLASSIC"},
	{"showdown", "FIRSTBLOOD"},
	{"poro_king", "KINGPORO"},
	{"dominion", "ODIN"},
	{"one_for_all", "ONEFORALL"},
	{"tutorial", "TUTORIAL"},
	{"nexus_siege", "SIEGE"},
}

var gameType = []struct {
	name, short string
}{
	{"custom", "CUSTOM_GAME"},
	{"tutorial", "TUTORIAL_GAME"},
	{"matched", "MATCHED_GAME"},
}

type queueStatus struct {
	gameMode  string
	gameType  string
	inGameMap string
	mapSide   string
	role      string
	lane      string
}

type playerStatus struct {
	summonerName string
	region       string
	season       string
	tier         string
	division     string
	platform     string
	inGame       bool
}

// GameData ...
var GameData struct {
	playerSt playerStatus
	queueSt  queueStatus
}

// SetRegion ...
func SetRegion(r string) (string, error) {
	for _, reg := range Region {
		if r == reg.country {
			GameData.playerSt.region = r
			return reg.short, nil
		}
		continue
	}
	if GameData.playerSt.region == "" {
		err := fmt.Errorf("Region %s not found", r)
		return "", err
	}
	return "", nil
}
