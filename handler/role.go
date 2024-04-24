package handler

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/lib/pq"
	"github.com/mccune1224/betrayal-widget/util"
	"github.com/mccune1224/betrayal-widget/view/page"
)

func (h *Handler) SearchRole(c echo.Context) error {
	game, _ := util.GetGame(c)
	search := c.FormValue("search")

	abilitiesIds := pq.Int64Array{}
	passiveIds := pq.Int64Array{}
	search = "%" + search + "%"
	err := h.models.Roles.DB.Select(&abilitiesIds, "SELECT id FROM abilities WHERE description ILIKE $1", search)
	err = h.models.Roles.DB.Select(&passiveIds, "SELECT id FROM passives WHERE description ILIKE $1", search)
	if err != nil {
		log.Println("HIT", err)
		return TemplRender(c, page.Error500(c, err))
	}

	log.Println(abilitiesIds, passiveIds, game.GameID)

	return nil
}
