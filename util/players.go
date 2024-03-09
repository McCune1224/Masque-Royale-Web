package util

import (
	"github.com/labstack/echo/v4"
	"github.com/mccune1224/betrayal-widget/data"
)

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

type PlayerContext struct {
	echo.Context
}

func (pc *PlayerContext) GetPlayers() ([]*data.ComplexPlayer, bool) {
	players, ok := pc.Get("players").([]*data.ComplexPlayer)
	return players, ok
}

func (pc *PlayerContext) SetPlayers(players []*data.ComplexPlayer) {
	pc.Set("players", players)
}

func GetPlayerContext(c echo.Context) *PlayerContext {
	return c.(*PlayerContext)
}

func GetPlayers(c echo.Context) ([]*data.ComplexPlayer, bool) {
	players, ok := c.Get("players").([]*data.ComplexPlayer)
	return players, ok
}

func ComplexToSimplePlayers(players []*data.ComplexPlayer) []*data.Player {
	simplePlayers := make([]*data.Player, len(players))
	for i, player := range players {
		simplePlayers[i] = &player.P
	}
	return simplePlayers
}

func GetNeighbor(target *data.ComplexPlayer, players []*data.ComplexPlayer, side string) *data.ComplexPlayer {
	ti := -1
	for i := range players {
		if players[i] == target {
			ti = i
			break
		}
	}

	if side == "next" {
		return players[(ti+1)%len(players)]
	} else {
		return players[(ti-1+len(players))%len(players)]
	}
}
