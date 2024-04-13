package util

import (
	"errors"

	"github.com/labstack/echo/v4"
	"github.com/mccune1224/betrayal-widget/data"
)

// Attempt to fetch the game, return error if no value found
func GetGame(c echo.Context) (*data.Game, error) {
	game, ok := c.Get("game").(*data.Game)
	if !ok {
		return nil, errors.New("invalid game id " + c.Path())
	}

	return game, nil
}
