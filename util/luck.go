package util

import (
	"github.com/mccune1224/betrayal-widget/data"
)

func BulkCalculateLuck(players []*data.ComplexPlayer) []*data.ComplexPlayer {
	for i := range players {
		target := players[i]
		prev := players[(i-1+len(players))%len(players)]
		next := players[(i+1)%len(players)]
		luck := CalculateLuck(&target.P, &target.R, &prev.P, &prev.R, &next.P, &next.R)
		players[i].P.Luck = luck
	}
	return players
}

func CalculateLuck(tPlayer *data.Player, tRole *data.Role, lPlayer *data.Player, lRole *data.Role, rPlayer *data.Player, rRole *data.Role) int {
	if !lPlayer.Alive && !rPlayer.Alive {
		return 0 + tPlayer.LuckModifier
	}

	if tRole.Name == "Vagabond" {
		return 0 + tPlayer.LuckModifier
	}

	luck := 0

	if !lPlayer.Alive || lRole.Name == "Vagabond" {
		luck += 0
	} else {
		luck += calculateNeighborLuck(tRole.Alignment, lRole.Alignment)
	}

	if !rPlayer.Alive || rRole.Name == "Vagabond" {
		luck += 0
	} else {
		luck += calculateNeighborLuck(tRole.Alignment, rRole.Alignment)
	}
	return luck + tPlayer.LuckModifier
}

func calculateNeighborLuck(targetAlignment string, neighborAlignment string) int {
	switch targetAlignment {
	case "Lawful":
		switch neighborAlignment {
		case "Lawful":
			return 5
		case "Chaotic":
			return 0
		case "Outlander":
			return 3
		}

	case "Chaotic":
		switch neighborAlignment {
		case "Lawful":
			return 0
		case "Chaotic":
			return 5
		case "Outlander":
			return 3
		}

	case "Outlander":
		return 3
	}

	return 0
}
