package handler

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/mccune1224/betrayal-widget/util"
	"github.com/mccune1224/betrayal-widget/view/page"
)

func (h *Handler) SearchRole(c echo.Context) error {
	game, _ := util.GetGame(c)
	search := c.FormValue("search")

	names := []string{}
	search = "%" + search + "%"
	err := h.models.Roles.DB.Select(&names, "SELECT name FROM abilities WHERE description LIKE $1", search)
	if err != nil {
		log.Println("HIT", err)
		return TemplRender(c, page.Error500(c, err))
	}

	log.Println(names, game.GameID)

	return nil
}
