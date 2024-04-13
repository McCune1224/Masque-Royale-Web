package handler

import (
	"log"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/mccune1224/betrayal-widget/util"
	"github.com/mccune1224/betrayal-widget/view/page"
)

func (h *Handler) IndexPage(c echo.Context) error {
	games, err := h.models.Games.GetAll()
	if err != nil {
		log.Println(err)
		return TemplRender(c, page.Error500(c, err))
	}

	return TemplRender(c, page.Index(c, games))
}

func (h *Handler) GameDashboardPage(c echo.Context) error {
	gID, err := strconv.Atoi(c.Param("game_id"))
	if err != nil {
		return TemplRender(c, page.Error500(c, err))
	}

	game, err := h.models.Games.GetByID(gID)
	if err != nil {
		return TemplRender(c, page.Error500(c, err))
	}

	return TemplRender(c, page.GameDashboard(c, game))
}

func (h *Handler) JoinGamePage(c echo.Context) error {
	gID, err := strconv.Atoi(c.Param("game_id"))
	if err != nil {
		return TemplRender(c, page.Error500(c, err))
	}

	game, err := h.models.Games.GetByID(gID)
	if err != nil {
		return TemplRender(c, page.Error500(c, err))
	}

	players, err := h.models.Players.GetAllByGameID(game.GameID)
	if err != nil {
		return TemplRender(c, page.Error500(c, err))
	}

	return TemplRender(c, page.JoinGame(c, game, players))
}

func (h *Handler) PlayerDashboardPage(c echo.Context) error {
	playerName := c.Param("player_name")
	game, err := util.GetGame(c)
	if err != nil {
		return TemplRender(c, page.Error500(c, err))
	}
	player, err := h.models.Players.GetByGameNameAndPlayerName(game.GameID, playerName)
	if err != nil {
		return TemplRender(c, page.Error500(c, err))
	}

	log.Println(player.RoleID)
	foo, err := h.models.Roles.Get(player.RoleID)
	if err != nil {
		return TemplRender(c, page.Error500(c, err))
	}
	log.Println(foo)

	role, err := h.models.Roles.GetComplexByID(player.RoleID)
	if err != nil {
		return TemplRender(c, page.Error500(c, err))
	}
	log.Println("HIT HERE")

	return TemplRender(c, page.PlayerDashboard(c, player, role))
}
