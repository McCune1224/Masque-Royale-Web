package handler

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/mccune1224/betrayal-widget/util"
	"github.com/mccune1224/betrayal-widget/views"
)

func (h *Handler) AllianceDashboard(c echo.Context) error {
	gameID := c.Param("game_id")

	players, _ := util.GetPlayers(c)
	alliances, err := h.models.Alliances.GetAllByGame(gameID)
	if err != nil {
		log.Println(err)
		return c.String(500, "Error getting alliances")
	}
	return TemplRender(c, views.AllianceDashboard(c, alliances, players))
}
