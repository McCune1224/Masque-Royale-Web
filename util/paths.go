package util

import (
	"strings"

	"github.com/labstack/echo/v4"
)

func GetGameIDURL(c echo.Context) string {
	// get the full url of the current path
	url := c.Request().URL.String()

	// split the url by the / and join up until dashboard/:gameid
	return strings.Join(strings.Split(url, "/")[:4], "/")
}
