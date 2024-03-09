package handler

import (
	"log"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/mccune1224/betrayal-widget/data"
	"github.com/mccune1224/betrayal-widget/util"
	"github.com/mccune1224/betrayal-widget/views"
)

func (h *Handler) SeatingDashboard(c echo.Context) error {
	if players, ok := util.GetPlayers(c); ok {
		return TemplRender(c, views.Seating(c, util.OrderComplexPlayers(players)))
	}
	return c.Redirect(302, "/")
}

func (h *Handler) SwapSeats(c echo.Context) error {
	players, _ := util.GetPlayers(c)

	var player1 *data.Player
	var player2 *data.Player

	p1 := strings.TrimSpace(strings.Split(c.FormValue("player1"), "-")[0])
	p2 := strings.TrimSpace(strings.Split(c.FormValue("player2"), "-")[0])

	for _, player := range players {
		if player.P.Name == p1 {
			player1 = &player.P
		}
		if player.P.Name == p2 {
			player2 = &player.P
		}
	}

	if player1 == nil || player2 == nil {
		log.Println("Invalid player")
		return c.Redirect(302, "/")
	}

	seat1 := player1.Seat
	seat2 := player2.Seat

	player1.Seat = seat2
	player2.Seat = seat1

	err := h.models.Players.Update(player1)
	if err != nil {
		log.Println(err)
		return c.Redirect(302, "/")
	}
	err = h.models.Players.Update(player2)
	if err != nil {
		log.Println(err)
		return c.Redirect(302, "/")
	}

	updatedPlayers, _ := h.models.Players.GetAllComplexByGameID(c.Param("game_id"))
	updatedPlayers = util.BulkCalculateLuck(updatedPlayers)
	for _, p := range updatedPlayers {
		err := h.models.Players.Update(&p.P)
		if err != nil {
			log.Println(err)
			return c.Redirect(302, "/")
		}
	}

	return TemplRender(c, views.Positions(c, util.OrderComplexPlayers(updatedPlayers)))
}
