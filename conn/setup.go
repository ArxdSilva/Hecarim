package conn

import (
	"fmt"

	"github.com/arxdsilva/Hecarim/gamedata"
)

var RiotApiKey, RiotTournamentKey, Region string

func SetApiKey(key string) {
	RiotApiKey = key
	return
}

func SetTournamentKey(key string) {
	RiotTournamentKey = key
	return
}

// SetRegion set's a region to be
func SetRegion(r string) error {
	for _, reg := range gamedata.Regions {
		if r == reg.Short {
			Region = r
			return nil
		}
	}
	return fmt.Errorf("Could not set region to %v, not a valid region.\n", r)
}
