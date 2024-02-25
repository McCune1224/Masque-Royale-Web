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
	game_id := c.Param("game_id")
	game, err := h.models.Games.GetByGameID(game_id)
	if err != nil || game == nil {
		log.Println(err)
		return c.Redirect(302, "/")
	}
	players, err := h.models.Players.GetByGameID(game_id)
	if err != nil {
		log.Println(err)
		return c.Redirect(302, "/")
	}
	c.Set("players", players)
	c.Set("game_id", game_id)

	return TemplRender(c, views.Seating(c, util.OrderPlayers(players)))
}

func (h *Handler) SwapSeats(c echo.Context) error {
	players, err := h.models.Players.GetByGameID(c.Param("game_id"))
	if err != nil {
		log.Println(err)
		return c.Redirect(302, "/")
	}

	var player1 *data.Player
	var player2 *data.Player

	p1 := strings.TrimSpace(strings.Split(c.FormValue("player1"), "-")[0])
	p2 := strings.TrimSpace(strings.Split(c.FormValue("player2"), "-")[0])

	for _, player := range players {
		if player.Name == p1 {
			player1 = player
		}
		if player.Name == p2 {
			player2 = player
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

	err = h.models.Players.Update(player1)
	if err != nil {
		log.Println(err)
		return c.Redirect(302, "/")
	}
	err = h.models.Players.Update(player2)
	if err != nil {
		log.Println(err)
		return c.Redirect(302, "/")
	}

	updatedPlayers, _ := h.models.Players.GetByGameID(c.Param("game_id"))
	return TemplRender(c, views.Positions(c, util.OrderPlayers(updatedPlayers)))
}
