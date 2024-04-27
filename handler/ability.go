package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/lib/pq"
	"github.com/mccune1224/betrayal-widget/data"
	"github.com/mccune1224/betrayal-widget/util"
	"github.com/mccune1224/betrayal-widget/view/page"
)

func (h *Handler) SearchAbility(c echo.Context) error {
	nameSearch := c.FormValue("name")
	// descriptionSearch := c.FormValue("description")
	categorySearch := c.FormValue("category")

	matchedAbilities := make(map[string]*data.AbilityFlashcardSearch)

	if nameSearch != "" {
		abilities := []*data.Ability{}
		// SELECT name FROM users WHERE levenshtein(name, 'user_input') <= 3;
		err := h.models.Abilities.Select(&abilities, "SELECT * from abilities WHERE name ILIKE $1", nameSearch)
		if err != nil {
			return TemplRender(c, page.Error500(c, err))
		}
		if len(abilities) != 1 {
			err := h.models.Abilities.Select(&abilities, "SELECT * FROM abilities WHERE levenshtein(name, $1) <=4 LIMIT 5", nameSearch)
			if err != nil {
				return TemplRender(c, page.Error500(c, err))
			}
		}

		for i := range abilities {
			associatedRoleNames := pq.StringArray{}
			err = h.models.Roles.Select(&associatedRoleNames, "SELECT name from roles WHERE $1 = ANY(ability_ids)", abilities[i].ID)
			if err != nil {
				return TemplRender(c, page.Error500(c, err))
			}
			flashcard := &data.AbilityFlashcardSearch{}
			flashcard.Ability = *abilities[i]
			flashcard.AssociatedRoles = associatedRoleNames
			matchedAbilities[abilities[i].Name] = flashcard
		}
	}

	// if descriptionSearch != "" {
	// 	abilities := []*data.Ability{}
	// 	err := h.models.Abilities.Select(&abilities, "SELECT * FROM abilities WHERE description ILIKE '%' || $1 || '%'", nameSearch)
	// 	if err != nil {
	// 		return TemplRender(c, page.Error500(c, err))
	// 	}
	// 	for i := range abilities {
	// 		matchedAbilities[abilities[i].Name] = abilities[i]
	// 	}
	// }

	// TODO: Match based off ability
	if categorySearch != "" {
	}

	searchResults := util.MaptoSlice(matchedAbilities)
	return TemplRender(c, page.AbilityResults(c, searchResults))
}
