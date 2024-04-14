package handler

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/mccune1224/betrayal-widget/data"
	"github.com/mccune1224/betrayal-widget/util"
	"github.com/mccune1224/betrayal-widget/view/page"
)

func (h *Handler) AddPlayerToGame(c echo.Context) error {
	playerName := c.FormValue("player")
	roleName := c.FormValue("role")
	game, _ := util.GetGame(c)

	players, err := h.models.Players.GetAllComplexByGameID(game.GameID)
	if err != nil {
		return TemplRender(c, page.Error500(c, err))
	}

	for _, v := range players {
		if strings.EqualFold(v.P.Name, playerName) {
			return TemplRender(c, page.PlayerList(c, players, fmt.Sprintf("player %s already added", playerName)))
		}
	}

	role, err := h.models.Roles.GetByName(roleName)
	if err != nil {
		return TemplRender(c, page.PlayerList(c, players, err.Error()))
	}

	newPlayer := &data.Player{
		ID:                0,
		Name:              playerName,
		GameID:            game.GameID,
		RoleID:            role.ID,
		Alive:             true,
		Seat:              0,
		Luck:              0,
		LuckModifier:      0,
		LuckStatus:        "",
		AlignmentOverride: "",
	}

	err = h.models.Players.Create(newPlayer)
	if err != nil {
		return TemplRender(c, page.Error500(c, err))
	}

	cp, _ := h.models.Players.GetComplexByGameIDAndName(newPlayer.GameID, newPlayer.Name)

	h.models.Games.Update(game)

	players = append(players, cp)

	game.PlayerCount = len(players)
	h.models.Games.Update(game)
	return TemplRender(c, page.PlayerList(c, players))
}

func (h *Handler) DeletePlayerFromGame(c echo.Context) error {
	playerID := util.ParamInt(c, "player_id", -1)
	game, _ := util.GetGame(c)
	h.models.Players.Delete(playerID)
	players, _ := h.models.Players.GetAllComplexByGameID(game.GameID)

	game.PlayerCount = len(players)
	h.models.Games.Update(game)
	return TemplRender(c, page.PlayerList(c, players))
}

func (h *Handler) MarkPlayerDead(c echo.Context) error {
	playerID := util.ParamInt(c, "player_id", -1)
	game, _ := util.GetGame(c)
	// TODO: FIX ME
	player, err := h.models.Players.GetByID(playerID)
	if err != nil {
		return TemplRender(c, page.Error500(c, err))
	}
	player.Alive = !player.Alive
	err = h.models.Players.UpdateDeathStatus(player.ID, player.Alive)
	if err != nil {
		return TemplRender(c, page.Error500(c, err))
	}

	cp, _ := h.models.Players.GetComplexByGameIDAndName(game.GameID, player.Name)

	return TemplRender(c, page.PlayerCard(c, cp))
}

func (h *Handler) SubmitPlayerAction(c echo.Context) error {
	playerID := util.ParamInt(c, "player_id", -1)
	context := strings.TrimSpace(c.FormValue("context"))
	target := strings.TrimSpace(c.FormValue("target"))
	abilityID := c.FormValue("ability")
	iAbilityID, _ := strconv.Atoi(abilityID)
	game, _ := util.GetGame(c)

	pa := &data.PlayerAction{
		ActionID:    iAbilityID,
		PlayerID:    playerID,
		Target:      target,
		Description: context,
	}

	err := h.models.Actions.InsertPlayerAction(pa)
	if err != nil {
		return TemplRender(c, page.Error500(c, err))
	}

	pa, err = h.models.Actions.GetPlayerActionByActionID(iAbilityID)
	if err != nil {
		return TemplRender(c, page.Error500(c, err))
	}

	actionList, err := h.models.Actions.GetActionList(game.GameID)
	if err != nil {
		return TemplRender(c, page.Error500(c, err))
	}

	actionList.ActionIds = append(actionList.ActionIds, int64(pa.ID))
	err = h.models.Actions.UpdateActionList(actionList)
	if err != nil {
		return TemplRender(c, page.Error500(c, err))
	}

	allPlayerActions, err := h.models.Actions.GetAllActionsForPlayer(pa.PlayerID)
	if err != nil {
		return TemplRender(c, page.Error500(c, err))
	}

	return TemplRender(c, page.PendingActions(c, allPlayerActions))
}
