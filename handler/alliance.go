package handler

import (
	"log"
	"sort"
	"strings"

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
	currAlliances, _ := util.GetAlliances(c)
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
		return TemplRender(c, views.AllianceCards(c, currAlliances))
	}

	if p1 == p2 {
		log.Println("Players cannot be the same")
		return TemplRender(c, views.AllianceCards(c, currAlliances))
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
		return TemplRender(c, views.AllianceCards(c, currAlliances))
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
		return TemplRender(c, views.AllianceCards(c, currAlliances))
	}

	updatedAlliances, err := h.models.Alliances.GetAllByGame(gameID)
	if err != nil {
		log.Println(err)
		return TemplRender(c, views.AllianceCards(c, updatedAlliances))
	}

	return TemplRender(c, views.AllianceCards(c, updatedAlliances))
}

func (h *Handler) UpdateAllianceColor(c echo.Context) error {
	fv := c.FormValue("color")
	split := strings.Split(fv, "-")
	allianceName := split[0]
	color := split[1]

	alliances, _ := util.GetAlliances(c)
	var targetAlliance *data.Alliance
	for _, a := range alliances {
		if a.Name == allianceName {
			targetAlliance = a
			break
		}
	}

	targetAlliance.Color = color
	h.models.Alliances.Update(targetAlliance)

	return TemplRender(c, views.AllianceCards(c, alliances))
}

func (h *Handler) AllianceDelete(c echo.Context) error {
	allianceName := c.QueryParam("name")
	log.Println(allianceName)

	alliances, _ := util.GetAlliances(c)
	var targetAlliance *data.Alliance
	for _, a := range alliances {
		if a.Name == allianceName {
			targetAlliance = a
			break
		}
	}

	h.models.Alliances.Delete(targetAlliance.ID)
	alliances = util.RemoveValue(alliances, targetAlliance)

	return TemplRender(c, views.AllianceCards(c, alliances))
}

func (h *Handler) AllianceChange(c echo.Context) error {
	players, _ := util.GetPlayers(c)
	alliances, _ := util.GetAlliances(c)

	allianceName := c.FormValue("alliance")
	player := c.QueryParam("name")
	log.Println(allianceName, player)

	var newAlliance *data.Alliance
	var oldAlliance *data.Alliance

	for _, a := range alliances {
		if a.Name == allianceName {
			newAlliance = a
		}
	}

	for _, a := range alliances {
		for _, m := range a.Members {
			if m == player {
				oldAlliance = a
				break
			}
		}
	}

	log.Println(oldAlliance, newAlliance)
	if oldAlliance != nil {
		oldAlliance.Members = util.RemoveValue(oldAlliance.Members, player)

		if len(oldAlliance.Members) == 0 {
			h.models.Alliances.Delete(oldAlliance.ID)
		} else {
			h.models.Alliances.Update(oldAlliance)
		}
	}

	newAlliance.Members = append(newAlliance.Members, player)
	h.models.Alliances.Update(newAlliance)

	alliances, err := h.models.Alliances.GetAllByGame(c.Param("game_id"))
	if err != nil {
		log.Println(err)
		return c.String(500, "Error getting alliances")
	}

	sort.Slice(alliances, func(i, j int) bool {
		return alliances[i].Name < alliances[j].Name
	})

	c.Set("alliances", alliances)

	return TemplRender(c, views.Positions(c, players))
}

func (h *Handler) AllianceLeave(c echo.Context) error {
	players, _ := util.GetPlayers(c)
	alliances, _ := util.GetAlliances(c)
	allianceName := c.FormValue("alliance")
	player := c.QueryParam("name")

	log.Println(player, allianceName)

	var oldAlliance *data.Alliance
	for _, a := range alliances {
		if a.Name == allianceName {
			oldAlliance = a
		}
	}
	oldAlliance.Members = util.RemoveValue(oldAlliance.Members, player)

	if len(oldAlliance.Members) == 0 {
		h.models.Alliances.Delete(oldAlliance.ID)
	} else {
		h.models.Alliances.Update(oldAlliance)
	}

	return TemplRender(c, views.Positions(c, players))
}
