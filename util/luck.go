package util

import (
	"github.com/mccune1224/betrayal-widget/data"
)

func BulkCalculateLuck(players []*data.ComplexPlayer) []*data.ComplexPlayer {
	for i := range players {
		target := players[i]
		prev := players[(i-1+len(players))%len(players)]
		next := players[(i+1)%len(players)]
		leftLuck := CalculateNeighborLuck(target, prev)
		rightLuck := CalculateNeighborLuck(target, next)
		players[i].P.Luck = leftLuck + rightLuck
	}
	return players
}

// func CalculateNeighboringLuck(tPlayer *data.Player, tRole *data.Role, lPlayer *data.Player, lRole *data.Role, rPlayer *data.Player, rRole *data.Role) int {
// 	if !lPlayer.Alive && !rPlayer.Alive {
// 		return 0 + tPlayer.LuckModifier
// 	}
//
// 	if tRole.Name == "Vagabond" {
// 		return 0 + tPlayer.LuckModifier
// 	}
//
// 	luck := 0
//
// 	if !lPlayer.Alive || lRole.Name == "Vagabond" {
// 		luck += 0
// 	} else {
// 		luck += CalculateAlignmentLuck(tRole.Alignment, lRole.Alignment)
// 	}
//
// 	if !rPlayer.Alive || rRole.Name == "Vagabond" {
// 		luck += 0
// 	} else {
// 		luck += CalculateAlignmentLuck(tRole.Alignment, rRole.Alignment)
// 	}
// 	return luck + tPlayer.LuckModifier
// }

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
	if !compared.P.Alive || !target.P.Alive {
		return 0
	}

	if target.R.Name == "Vagabond" || compared.R.Name == "Vagabond" {
		return 0
	}

  var luck int
  if target.P.AlignmentOverride != "" {
    luck = CalculateAlignmentLuck(target.P.AlignmentOverride, compared.R.Alignment)
  } else {
    luck = CalculateAlignmentLuck(target.R.Alignment, compared.R.Alignment)
  }
	return luck
}
