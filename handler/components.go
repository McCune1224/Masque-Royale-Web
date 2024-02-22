package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/mccune1224/betrayal-widget/views/components"
)

func (h *Handler) PlayerDropdownModal(c echo.Context) error {
  playerName := c.QueryParam("player")

  dbPlayer, err := h.models.Players.GetByGameIDAndName(c.Param("game_id"), playerName)
  if err != nil {
    return err
  }

	return TemplRender(c, components.PlayerModal(dbPlayer))
}
