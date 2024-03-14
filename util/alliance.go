package util

import (
	"github.com/labstack/echo/v4"
	"github.com/mccune1224/betrayal-widget/data"
)

func GetAlliances(c echo.Context) ([]*data.Alliance, bool) {
	alliances, ok := c.Get("alliances").([]*data.Alliance)
	return alliances, ok
}

func AllianceContainsPlayer(a *data.Alliance, p *data.Player) bool {
	for _, member := range a.Members {
		if member == p.Name {
			return true
		}
	}
	return false
}

func PlayerWithinAlliance(p *data.Player, alliances []*data.Alliance) *data.Alliance {
  for _, alliance := range alliances {
    if AllianceContainsPlayer(alliance, p) {
      return alliance
    }
  }
  return nil
}
