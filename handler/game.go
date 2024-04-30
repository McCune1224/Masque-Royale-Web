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

func (h *Handler) GamePhaseIncrement(c echo.Context) error {
	game, _ := util.GetGame(c)
	if game.Phase == "Night" {
		game.Phase = "Day"
		game.Round += 1
	} else {
		game.Phase = "Night"
	}
	err := h.models.Games.UpdateProperty(game.ID, "phase", game.Phase)
	if err != nil {
		return TemplRender(c, page.Error500(c, err))
	}
	err = h.models.Games.UpdateProperty(game.ID, "round", game.Round)
	if err != nil {
		return TemplRender(c, page.Error500(c, err))
	}

	return TemplRender(c, page.CycleDashboard(c, game))
}

func (h *Handler) GamePhaseDecrement(c echo.Context) error {
	game, _ := util.GetGame(c)

	if game.Phase == "Day" && game.Round == 1 {
		return TemplRender(c, page.CycleDashboard(c, game))
	}

	if game.Phase == "Day" {
		game.Phase = "Night"
		game.Round -= 1
	} else {
		game.Phase = "Day"
	}
	err := h.models.Games.UpdateProperty(game.ID, "phase", game.Phase)
	if err != nil {
		return TemplRender(c, page.Error500(c, err))
	}
	err = h.models.Games.UpdateProperty(game.ID, "round", game.Round)
	if err != nil {
		return TemplRender(c, page.Error500(c, err))
	}

	return TemplRender(c, page.CycleDashboard(c, game))
}
