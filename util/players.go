package util

import "github.com/mccune1224/betrayal-widget/data"

func OrderPlayers(players []*data.Player) []*data.Player {
	orderedPlayers := make([]*data.Player, len(players))
	for _, player := range players {
		orderedPlayers[player.Seat-1] = player
	}
	return orderedPlayers
}
