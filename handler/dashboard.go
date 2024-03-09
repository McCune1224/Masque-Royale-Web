package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/mccune1224/betrayal-widget/util"
	"github.com/mccune1224/betrayal-widget/views"
)

func (h *Handler) Dashboard(c echo.Context) error {
	if players, ok := util.GetPlayers(c); ok {
		return TemplRender(c, views.Home(c, players))
	}

	return c.Redirect(302, "/")
}
