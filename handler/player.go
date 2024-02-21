package handler

import (
	"fmt"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/mccune1224/betrayal-widget/data"
	"github.com/mccune1224/betrayal-widget/views"
)

// func (h *Handler) AddPlayer(c echo.Context) error {
// }

func (h *Handler) PlayerDashboard(c echo.Context) error {
	players, err := h.models.Players.GetByGameID(c.Param("game_id"))
	if err != nil {
		log.Println(err)
		return c.Redirect(302, "/")
	}

	game, err := h.models.Games.GetByGameID(c.Param("game_id"))
	if err != nil {
		log.Println(err)
		return c.Redirect(302, "/")
	}

	roles, err := h.models.Roles.GetAll()
	if err != nil {
		log.Println(err)
		return c.Redirect(302, "/")
	}

	return TemplRender(c, views.PlayerDashboard(c, game, players, roles))
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
