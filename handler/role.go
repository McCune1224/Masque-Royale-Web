package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/mccune1224/betrayal-widget/data"
	"github.com/mccune1224/betrayal-widget/view/page"
)

func (h *Handler) SearchRole(c echo.Context) error {
	nameSearch := c.FormValue("name")
	descriptionSearch := c.FormValue("description")

	searchResults := []*data.ComplexRole{}
	if nameSearch != "" {
		roles := []*data.ComplexRole{}
		err := h.models.Roles.DB.Select(&roles, "SELECT * from roles WHERE name ILIKE $1", nameSearch)
		if err != nil {
			return TemplRender(c, page.Error500(c, err))
		}
		if len(roles) != 1 {
			// SELECT * FROM abilities WHERE levenshtein(name, $1) <=4 LIMIT 5
			err := h.models.Roles.DB.Select(&roles, "SELECT * FROM roles WHERE levenshtein(name, $1) <=4 LIMIT 5", nameSearch)
			if err != nil {
				return TemplRender(c, page.Error500(c, err))
			}
		}
		searchResults = append(searchResults, roles...)
	}

	if descriptionSearch != "" {
	}

	return TemplRender(c, page.PlayerRoleResults(c, searchResults))
}
