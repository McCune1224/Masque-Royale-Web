package handler

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/lib/pq"
	"github.com/mccune1224/betrayal-widget/data"
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

func (h *Handler) AllianceCreate(c echo.Context) error {
	gameID := c.Param("game_id")

	p1 := c.FormValue("player1")
	p2 := c.FormValue("player2")
	color := c.FormValue("color")
	allianceName := c.FormValue("name")

	if color == "" {
		color = "white"
	}

	if allianceName == "" {
		log.Println("Alliance name is required")
		return c.Redirect(302, "/")
	}

	if p1 == p2 {
		log.Println("Players cannot be the same")
		return c.Redirect(302, "/")
	}

	var player1 *data.ComplexPlayer
	var player2 *data.ComplexPlayer
	players, _ := util.GetPlayers(c)
	for _, p := range players {
		if p.P.Name == p1 {
			player1 = p
		}
		if p.P.Name == p2 {
			player2 = p
		}
	}

	if player1 == nil || player2 == nil {
		log.Println("Could not find players")
		return c.Redirect(302, "/")
	}

	alliance := &data.Alliance{
		Name:        allianceName,
		Description: "",
		Members:     pq.StringArray{player1.P.Name, player2.P.Name},
		GameID:      gameID,
		Color:       color,
	}
	err := h.models.Alliances.Create(alliance)
	if err != nil {
		log.Println(err)
		return c.String(500, "Error creating alliance")
	}

	updatedAlliances, err := h.models.Alliances.GetAllByGame(gameID)
	if err != nil {
		log.Println(err)
		return c.Redirect(302, "/")
	}

	return TemplRender(c, views.AllianceCards(c, updatedAlliances))
}
