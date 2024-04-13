package util

import (
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/mccune1224/betrayal-widget/data"
)

func GetGameIDURL(c echo.Context) string {
	// get the full url of the current path
	url := c.Request().URL.String()
	// split the url by the / and join up until dashboard/:gameid
	return strings.Join(strings.Split(url, "/")[:4], "/")
}

func GamePath(c echo.Context) string {
	url := c.Request().URL.String()
	return strings.Join(strings.Split(url, "/")[:3], "/")
}

func PlayerPath(c echo.Context, playerID string) string {
	gameURL := GamePath(c)

	return gameURL + "/players/" + playerID
}

func GetPlayerUpdateURL(c echo.Context, target *data.ComplexPlayer) string {
	return GetGameIDURL(c) + "/menu/update/" + target.P.Name
}
