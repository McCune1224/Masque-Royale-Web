package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/mccune1224/betrayal-widget/data"
)

func (h *Handler) GetRandomGame(c echo.Context) error {

	game := &data.Game{}
	err := h.models.Games.DB.Get(game, "SELECT * FROM games ORDER BY random() LIMIT 1")
	if err != nil {
		return c.JSON(500,
			echo.Map{"error": err.Error()},
		)
	}

	return c.JSON(200, game)
}
