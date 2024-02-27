package handler

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/mccune1224/betrayal-widget/util"
	"github.com/mccune1224/betrayal-widget/views"
)

func (h *Handler) Luck(c echo.Context) error {
	players, err := h.models.Players.GetComplexByGameID(c.Param("game_id"))
	if err != nil {
		log.Println(err)
		return c.Redirect(302, "/")
	}

	players = util.OrderComplexPlayers(players)

	for i := range players {
		target := players[i]
		// get previous index and next index (wrapping around if on the edge)
		prev := players[(i-1+len(players))%len(players)]
		next := players[(i+1)%len(players)]
		luck := util.CalculateLuck(&target.P, &target.R, &prev.P, &prev.R, &next.P, &next.R)

		log.Println("LUCK %s - %d", target.P.Name, luck)
	}

	return TemplRender(c, views.Luck(c, players))
}
