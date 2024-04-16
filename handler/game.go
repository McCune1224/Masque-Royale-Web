package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/mccune1224/betrayal-widget/util"
	"github.com/mccune1224/betrayal-widget/view/page"
)

func (h *Handler) CreateGame(c echo.Context) error {
	games, err := h.models.Games.GetAll()
	if err != nil {
		return TemplRender(c, page.Error500(c, err))
	}

	gameName := c.FormValue("name")
	existingGame, _ := h.models.Games.GetByGameID(gameName)

	if existingGame != nil {
		return TemplRender(c, page.Index(c, games, "Game Already Exists"))
	}

	game, err := h.models.Games.InsertGame(gameName, 0)
	if err != nil {
		return TemplRender(c, page.Index(c, games, err.Error()))
	}

	games = append(games, *game)

	return TemplRender(c, page.Index(c, games))
}

func (h *Handler) DeleteGame(c echo.Context) error {
	gameID := util.ParamInt(c, "game_id", -1)
	game, _ := h.models.Games.GetByID(gameID)
	err := h.models.Games.DeleteGame(game.GameID)
	if err != nil {
		TemplRender(c, page.Error500(c, err))
	}

	games, _ := h.models.Games.GetAll()
	return TemplRender(c, page.Index(c, games))
}
