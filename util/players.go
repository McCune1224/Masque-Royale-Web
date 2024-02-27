package util

import "github.com/mccune1224/betrayal-widget/data"

func OrderPlayers(players []*data.Player) []*data.Player {
	orderedPlayers := make([]*data.Player, len(players))
	for _, player := range players {
		orderedPlayers[player.Seat-1] = player
	}
	return orderedPlayers
}

func OrderComplexPlayers(players []*data.ComplexPlayer) []*data.ComplexPlayer {
	orderedPlayers := make([]*data.ComplexPlayer, len(players))
	for _, player := range players {
		orderedPlayers[player.P.Seat-1] = player
	}
	return orderedPlayers
}
