package handler

import (
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/mccune1224/betrayal-widget/views/components"
)

func (h *Handler) PlayerDropdownModal(c echo.Context) error {
	playerName := c.QueryParam("player")
	path := strings.Split(c.QueryParam("game_id"), "/")
	gameID := path[len(path)-1]

	dbPlayer, err := h.models.Players.GetByGameIDAndName(gameID, playerName)
	if err != nil {
		return err
	}

	return TemplRender(c, components.PlayerModal(dbPlayer))
}
