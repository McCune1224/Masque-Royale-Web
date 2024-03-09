package handler

import (
	"fmt"
	"log"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/mccune1224/betrayal-widget/views/create"
)

func (h *Handler) CreateGame(c echo.Context) error {
	return TemplRender(c, create.GameCreate(""))
}

func (h *Handler) GenerateGame(c echo.Context) error {
	gameID := strings.ToLower(c.QueryParam("game_id"))
	if gameID == "" {
		return TemplRender(c, create.GameCreate("Game ID is required"))
	}
	existingID, _ := h.models.Games.GetByGameID(gameID)
	if existingID != nil {
		return TemplRender(c, create.GameCreate(fmt.Sprintf("Game ID %s already exists", gameID)))
	}

	_, err := h.models.Games.InsertGame(gameID, 0)
	if err != nil {
		return err
	}

	c.Set("game_id", gameID)
	return c.Redirect(302, "/games/dashboard/"+gameID)
}

func (h *Handler) JoinGame(c echo.Context) error {
	gameID := c.Param("game_id")
	if gameID == "" {
		return c.HTML(400, "game_id is required")
	}

	_, err := h.models.Games.GetByGameID(gameID)
	if err != nil {
		log.Println(err)
		return c.HTML(500, "<div>Error getting game</div>")
	}
  players, _ := h.models.Players.GetAllComplexByGameID(gameID)
  c.Set("players", players)

	return c.Redirect(302, "/games/dashboard/"+gameID)
}

func (h *Handler) DeleteGame(c echo.Context) error {
	gameID := c.Param("game_id")
	if gameID == "" {
		return c.HTML(400, "game_id is required")
	}
	err := h.models.Games.DeleteGame(gameID)
	if err != nil {
		return err
	}
	return c.Redirect(302, "/")
}
