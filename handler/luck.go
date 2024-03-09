package handler

import (
	"log"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/mccune1224/betrayal-widget/views"
)

func (h *Handler) Luck(c echo.Context) error {
	playerNames, err := h.models.Players.GetPlayerNames(c.Param("game_id"))
	if err != nil {
		log.Println(err)
		return c.Redirect(302, "/")
	}
	return TemplRender(c, views.Luck(c, playerNames))
}

func (h *Handler) LuckUpdate(c echo.Context) error {
	formPlayer := c.FormValue("player")
	formModifier := c.FormValue("modifier")
	player, err := h.models.Players.GetComplexByGameID(c.Param("game_id"), formPlayer)
	if err != nil {
		log.Println(err)
		return c.Redirect(302, "/")
	}
	iModifier, err := strconv.Atoi(formModifier)
	if err != nil {

		log.Println(err)
		return c.Redirect(302, "/")
	}
	player.P.LuckModifier = iModifier
	err = h.models.Players.Update(&player.P)
	if err != nil {
		log.Println(err)
		return c.Redirect(302, "/")
	}

  return TemplRender(c, views.PlayerToken(player))
}
