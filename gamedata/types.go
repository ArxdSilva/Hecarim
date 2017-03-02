package gamedata

// Regions ...
var Regions = []struct {
	Country, Short string
}{
	{"brazil", "BR"},
	{"europe_north_east", "EUNE"},
	{"europe_west", "EUW"},
	{"japan", "JP"},
	{"korea", "KR"},
	{"latin_america_north", "LAN"},
	{"latin_america_south", "LAS"},
	{"north_america", "NA"},
	{"oceania", "OCE"},
	{"pbe", "PBE"},
	{"russia", "RU"},
	{"turkey", "TR"},
}

var gameMode = []struct {
	Name, Short string
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
	PlayerSt playerStatus
	QueueSt  queueStatus
}
