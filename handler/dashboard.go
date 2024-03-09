package handler

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/mccune1224/betrayal-widget/util"
	"github.com/mccune1224/betrayal-widget/views"
)

func (h *Handler) Dashboard(c echo.Context) error {
	game_id := c.Param("game_id")
	game, err := h.models.Games.GetByGameID(game_id)
	if err != nil || game == nil {
		log.Println(err)
		return c.Redirect(302, "/")
	}
	players, err := h.models.Players.GetAllComplexByGameID(game_id)
	if err != nil {
		log.Println(err)
		return c.Redirect(302, "/")
	}
	c.Set("players", players)
	c.Set("game_id", game_id)

	return TemplRender(c, views.Home(c, util.OrderComplexPlayers(players)))
}
