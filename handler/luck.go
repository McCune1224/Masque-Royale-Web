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
	players = util.BulkCalculateLuck(players)
	for _, p := range players {
		err := h.models.Players.Update(&p.P)
		if err != nil {
			log.Println(err)
			return c.Redirect(302, "/")
		}
	}

	return TemplRender(c, views.Luck(c, players))
}
