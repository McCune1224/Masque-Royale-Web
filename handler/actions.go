package handler

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/mccune1224/betrayal-widget/views"
)

func (h *Handler) ActionDashboard(c echo.Context) error {
	actions, err := h.models.Actions.GetAll()
	if err != nil {
		log.Println(err)
		return c.Redirect(302, "/")
	}
	return TemplRender(c, views.ActionDashboard(c, actions))
}
