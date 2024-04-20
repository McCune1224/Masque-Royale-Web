package util

import (
	"slices"
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

// Split basepath url and confirm if within path
func IsResourcePath(c echo.Context, target string) bool {
	resources := strings.Split(c.Path(), "/")
	targetResourceIndex := slices.Index(resources, target)
	return targetResourceIndex != -1
}

// Will return path type based on if admin, players, or base
func ResourcePathType(c echo.Context) string {
	if IsResourcePath(c, "admin") {
		return "admin"
	}

	if IsResourcePath(c, "players") {
		return "players"
	}

	return "base"
}
