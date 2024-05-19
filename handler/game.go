package handler

import (
	"strconv"

	"github.com/labstack/echo/v4"
)

func (h *Handler) GetRandomGame(c echo.Context) error {
	game, err := h.models.Games.GetRandomGame()
	if err != nil {
		return c.JSON(500,
			echo.Map{"message": err.Error()},
		)
	}
	return c.JSON(200, game)
}

func (h *Handler) GetAllGames(c echo.Context) error {
	games, err := h.models.Games.GetAllGames()
	if err != nil {
		return c.JSON(500,
			echo.Map{"message": err.Error()},
		)
	}

	return c.JSON(200, games)
}

func (h *Handler) GetGameByID(c echo.Context) error {
	gameIdParam := c.Param("game_id")
	gameId, err := strconv.Atoi(gameIdParam)
	if err != nil {
		return c.JSON(400,
			echo.Map{"message": "Invalid Game ID"},
		)
	}

	game, err := h.models.Games.GetGameByID(gameId)
	if err != nil {
		return c.JSON(500, echo.Map{"message": err.Error()})
	}

	return c.JSON(200, game)
}
