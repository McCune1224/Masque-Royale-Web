package handler

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/mccune1224/betrayal-widget/view/page"
)

func (h *Handler) IndexPage(c echo.Context) error {
	games, err := h.models.Games.GetAll()
	if err != nil {
		log.Println(err)
		return TemplRender(c, page.Error500(c, err))
	}

	return TemplRender(c, page.Index(c, games))
}
