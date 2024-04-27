package handler

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/lib/pq"
	"github.com/mccune1224/betrayal-widget/data"
	"github.com/mccune1224/betrayal-widget/util"
	"github.com/mccune1224/betrayal-widget/view/page"
)

func (h *Handler) SearchRole(c echo.Context) error {
	search := c.FormValue("search")

	abilitiesIds := pq.Int64Array{}
	passiveIds := pq.Int64Array{}
	search = "%" + search + "%"
	err := h.models.Roles.DB.Select(&abilitiesIds, "SELECT id FROM abilities WHERE description ILIKE $1", search)
	err = h.models.Roles.DB.Select(&passiveIds, "SELECT id FROM passives WHERE description ILIKE $1", search)
	abIDS := pq.Int64Array{}
	paIDS := pq.Int64Array{}
	searchIDs := pq.Int64Array{}
	err := h.models.Roles.DB.Select(&abIDS, "SELECT id FROM abilities WHERE description ILIKE '%' || $1 || '%'", search)
	if err != nil {
		log.Println("HIT", err)
		return TemplRender(c, page.Error500(c, err))
	}

	log.Println(abilitiesIds, passiveIds, game.GameID)
	err = h.models.Roles.DB.Select(&paIDS, "SELECT id FROM passives WHERE description ILIKE '%' || $1 || '%'", search)
	if err != nil {
		log.Println("HIT", err)
		return TemplRender(c, page.Error500(c, err))
	}

	searchIDs = append(searchIDs, abIDS...)
	searchIDs = append(searchIDs, paIDS...)

	matchingRoles := []*data.SearchComplexRoleResult{}
	for _, currSearchID := range searchIDs {
		var roleID int
		err = h.models.Roles.DB.Get(&roleID, "SELECT id from roles WHERE $1 = ANY(ability_ids) OR $1 = ANY(passive_ids)", currSearchID)
		if err != nil {
			return TemplRender(c, page.Error500(c, err))
		}

		if roleID != 0 {
			role, err := h.models.Roles.GetComplexByID(roleID)
			if err != nil {
				return TemplRender(c, page.Error500(c, err))
			}
			temp := &data.SearchComplexRoleResult{
				CR:          role,
				MatchedName: search,
			}
			matchingRoles = util.UniqueInsert(matchingRoles, temp)
		}
	}

	log.Println(searchIDs, search)

	return TemplRender(c, page.PlayerRoleResults(c, matchingRoles, search))
}
