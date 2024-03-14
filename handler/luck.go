package handler

import (
	"log"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/mccune1224/betrayal-widget/data"
	"github.com/mccune1224/betrayal-widget/util"
	"github.com/mccune1224/betrayal-widget/views"
)

func (h *Handler) Luck(c echo.Context) error {
	playerNames := []string{}
	players, _ := util.GetPlayers(c)
	for _, player := range players {
		playerNames = append(playerNames, player.P.Name)
		log.Println(player)
	}
	return TemplRender(c, views.Luck(c, playerNames))
}

func (h *Handler) LuckUpdate(c echo.Context) error {
	formPlayer := c.FormValue("player")
	formModifier := c.FormValue("modifier")
	players, _ := util.GetPlayers(c)

	var targetPlayer *data.ComplexPlayer
	for _, player := range players {
		if strings.EqualFold(player.P.Name, formPlayer) {
			targetPlayer = player
			break
		}
	}

	iModifier, err := strconv.Atoi(formModifier)
	if err != nil {

		log.Println(err)
		return c.Redirect(302, "/")
	}
	targetPlayer.P.LuckModifier = iModifier
	err = h.models.Players.Update(&targetPlayer.P)
	if err != nil {
		log.Println(err)
		return c.Redirect(302, "/")
	}

	return TemplRender(c, views.PlayerToken(c, targetPlayer))
}
