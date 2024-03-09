package handler

import (
	"fmt"
	"log"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/mccune1224/betrayal-widget/data"
	"github.com/mccune1224/betrayal-widget/util"
	"github.com/mccune1224/betrayal-widget/views"
	"github.com/mccune1224/betrayal-widget/views/components"
)

func (h *Handler) PlayerDashboard(c echo.Context) error {
	players := c.Get("players").([]*data.ComplexPlayer)
	p := []*data.Player{}
	r := []*data.Role{}
	for _, player := range players {
		p = append(p, &player.P)
		r = append(r, &player.R)
	}

	return TemplRender(c, views.PlayerDashboard(c, p, r))
}

func (h *Handler) PlayerAdd(c echo.Context) error {
	log.Println(c.FormParams())

	formPlayerName := c.FormValue("name")
	formRoleName := c.FormValue("role")

	currPlayers, err := h.models.Players.GetByGameID(c.Param("game_id"))
	if err != nil {
		log.Println(err)
		return c.Redirect(302, "/")
	}
	currPlayers = util.OrderPlayers(currPlayers)

	existingRoles, err := h.models.Roles.GetAll()
	if err != nil {
		log.Println(err)
		return c.Redirect(302, "/")
	}

	for _, player := range currPlayers {
		if player.Name == formPlayerName {
			return TemplRender(c, views.PlayerList(c, currPlayers, existingRoles, []string{fmt.Sprintf("Player %s already exists", formPlayerName)}))
		}
	}

	found := false
	selectedRoleID := 0
	for _, role := range existingRoles {
		if role.Name == formRoleName {
			found = true
			selectedRoleID = role.ID
			break
		}
	}

	if !found {
		return TemplRender(c, views.PlayerList(c, currPlayers, existingRoles, []string{"Role does not exist"}))
	}

	newPlayer := &data.Player{
		ID:     selectedRoleID,
		Name:   formPlayerName,
		GameID: c.Param("game_id"),
		RoleID: selectedRoleID,
		Alive:  true,
		Seat:   len(currPlayers) + 1,
	}

	err = h.models.Players.Create(newPlayer)
	if err != nil {
		log.Println(err)
		return c.Redirect(302, "/")
	}

	currPlayers = append(currPlayers, newPlayer)
	h.models.Games.UpdatePlayerCount(c.Param("game_id"), len(currPlayers))
	return TemplRender(c, views.PlayerList(c, currPlayers, existingRoles, nil))
}

func (h *Handler) PlayerRemove(c echo.Context) error {
	player := c.FormValue("player")

	existingPlayers, err := h.models.Players.GetByGameID(c.Param("game_id"))
	if err != nil {
		log.Println(err)
		return c.Redirect(302, "/")
	}

	for _, p := range existingPlayers {
		if p.Name == player {
			err = h.models.Players.Delete(p.ID)
			if err != nil {
				log.Println(err)
				return c.Redirect(302, "/")
			}
			break
		}
	}

	roles, err := h.models.Roles.GetAll()
	if err != nil {
		log.Println(err)
		return c.Redirect(302, "/")
	}

	return TemplRender(c, views.PlayerList(c, existingPlayers, roles, nil))
}

func (h *Handler) PlayerReposition(c echo.Context) error {
	targetSeat, err := strconv.Atoi(c.FormValue("seat"))
	if err != nil {
		log.Println(err)
		return c.Redirect(302, "/")
	}

	playerName := c.FormValue("player")

	players, _ := util.GetPlayers(c)

	if targetSeat < 1 || targetSeat > len(players) {
		return c.Redirect(302, "/")
	}

	targetPlayer := &data.Player{}
	for _, p := range players {
		if p.P.Name == playerName {
			targetPlayer = &p.P
			break
		}
	}

	swapped := false
	for _, p := range players {
		if p.P.Seat == targetSeat && p.P.Name != targetPlayer.Name {
			p.P.Seat = targetPlayer.Seat
			targetPlayer.Seat = targetSeat
			h.models.Players.Update(&p.P)
			h.models.Players.Update(targetPlayer)
			swapped = true
			break
		}
	}

	if !swapped {
		targetPlayer.Seat = targetSeat
		h.models.Players.Update(targetPlayer)
	}
	p := util.ComplexToSimplePlayers(players)

	//  r := []*data.Role{}
	// for _, player := range players {
	// 	r = append(r, &player.R)
	// }
	//
	return TemplRender(c, views.PlayerList(c, p, nil, nil))
}

func (h *Handler) PlayerMenu(c echo.Context) error {
	players, _ := util.GetPlayers(c)
	playerName := c.QueryParam("name")
	var targetPlayer *data.ComplexPlayer
	for _, player := range players {
		if player.P.Name == playerName {
			targetPlayer = player
			break
		}
	}
	return TemplRender(c, components.PlayerMenu(c, targetPlayer, players))
}

func (h *Handler) UpdatePlayerLuckModifier(c echo.Context) error {
	log.Println("HIT----------------------------------")
	player := c.Param("player")
	mod := c.FormValue("modifier")

	players, _ := util.GetPlayers(c)
	var targetPlayer *data.ComplexPlayer
	for _, p := range players {
		if p.P.Name == player {
			targetPlayer = p
			break
		}
	}
	iMod, err := strconv.Atoi(mod)
	if err != nil {
		log.Println(err)
		return c.Redirect(302, "/")
	}

	targetPlayer.P.LuckModifier = iMod
	err = h.models.Players.Update(&targetPlayer.P)
	if err != nil {
		log.Println(err)
		return err
	}
	return TemplRender(c, views.PlayerToken(targetPlayer))
}
