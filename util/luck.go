package util

import (
	"github.com/mccune1224/betrayal-widget/data"
)

func BulkCalculateLuck(players []*data.ComplexPlayer) []*data.ComplexPlayer {
	for i := range players {
		target := players[i]
		prev := players[(i-1+len(players))%len(players)]
		next := players[(i+1)%len(players)]
		luck := CalculateNeighboringLuck(&target.P, &target.R, &prev.P, &prev.R, &next.P, &next.R)
		players[i].P.Luck = luck
	}
	return players
}

func CalculateNeighboringLuck(tPlayer *data.Player, tRole *data.Role, lPlayer *data.Player, lRole *data.Role, rPlayer *data.Player, rRole *data.Role) int {
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
		luck += CalculateAlignmentLuck(tRole.Alignment, lRole.Alignment)
	}

	if !rPlayer.Alive || rRole.Name == "Vagabond" {
		luck += 0
	} else {
		luck += CalculateAlignmentLuck(tRole.Alignment, rRole.Alignment)
	}
	return luck + tPlayer.LuckModifier
}

func CalculateAlignmentLuck(targetAlignment string, comparedAlignment string) int {
	switch targetAlignment {
	case "Lawful":
		switch comparedAlignment {
		case "Lawful":
			return 5
		case "Chaotic":
			return 0
		case "Outlander":
			return 3
		}

	case "Chaotic":
		switch comparedAlignment {
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

func CalculateNeighborLuck(target *data.ComplexPlayer, compared *data.ComplexPlayer) int {
	if !compared.P.Alive {
		return 0
	}

	if target.R.Name == "Vagabond" || compared.R.Name == "Vagabond" {
		return 0
	}

	luck := 0

	if !compared.P.Alive || compared.R.Name == "Vagabond" {
		luck += 0
	} else {
		luck += CalculateAlignmentLuck(target.R.Alignment, compared.R.Alignment)
	}
	return luck
}
