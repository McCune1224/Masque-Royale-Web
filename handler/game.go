package handler

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/mccune1224/betrayal-widget/views"
)

func (h *Handler) CreateGame(c echo.Context) error {
	return TemplRender(c, views.GameCreate(""))
}

func (h *Handler) GenerateGame(c echo.Context) error {
	gameID := c.QueryParam("game_id")
	if gameID == "" {
		return TemplRender(c, views.GameCreate("Game ID is required"))
	}
	existingID, _ := h.models.Games.GetByGameID(gameID)
	if existingID != nil {
		return TemplRender(c, views.GameCreate(fmt.Sprintf("Game ID %s already exists", gameID)))
	}

	playerCount := c.QueryParam("player_count")
	if playerCount == "" {
		return TemplRender(c, views.GameCreate("Player Count is required"))
	}
	iPlayerCount, err := strconv.Atoi(playerCount)
	if err != nil {
		return TemplRender(c, views.GameCreate("Player Count must be a number"))
	}

	_, err = h.models.Games.InsertGame(gameID, iPlayerCount)
	if err != nil {
		return err
	}

	return TemplRender(c, views.Dashboard(c))
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
	c.SetCookie(&http.Cookie{
		Name:     "game_id",
		Value:    gameID,
		Path:     "/",
		SameSite: http.SameSiteNoneMode,
		Secure:   true,
	})

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
