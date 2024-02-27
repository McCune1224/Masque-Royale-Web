package util

import "github.com/mccune1224/betrayal-widget/data"

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
	case "LAWFUL":
		switch neighborAlignment {
		case "LAWFUL":
			return 5
		case "CHAOTIC":
			return 0
		case "OUTLANDER":
			return 3
		}

	case "CHAOTIC":
		switch neighborAlignment {
		case "LAWFUL":
			return 0
		case "CHAOTIC":
			return 5
		case "OUTLANDER":
			return 3
		}

	case "OUTLANDER":
		return 3
	}

	return 0
}
