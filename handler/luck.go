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
		prev := players[(i-1+len(players))%len(players)]
		next := players[(i+1)%len(players)]
		luck := util.CalculateLuck(&target.P, &target.R, &prev.P, &prev.R, &next.P, &next.R)
		log.Printf("LUCK %s - %d - %s (%s)", target.P.Name, luck, target.R.Name, target.R.Alignment[:1])
		players[i].P.Luck = luck
		err := h.models.Players.Update(&players[i].P)
		if err != nil {
			log.Println(err)
			return c.Redirect(302, "/")
		}
	}

	return TemplRender(c, views.Luck(c, players))
}
