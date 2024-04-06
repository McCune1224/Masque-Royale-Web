package handler

import (
	"log"
	"net/http"

	"github.com/a-h/templ"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/lithammer/fuzzysearch/fuzzy"
	"github.com/mccune1224/betrayal-widget/data"
	"github.com/mccune1224/betrayal-widget/views"
)

var potentialRoles = []string{
	"Agent", "Detective", "Gunman", "Lawyer", "Nurse", "Seraph", "Empress", "Succubus", "Wraith", "Actress", "Assassin", "Highwayman", "Jester", "Sommelier", "Witchdoctor",
}

type Handler struct {
	models *data.Models
}

func NewHandler(db *sqlx.DB) *Handler {
	models := data.NewModels(db)

	return &Handler{
		models: models,
	}
}

// Helper to wrap calling a component's Render method.
func TemplRender(c echo.Context, component templ.Component) error {
	return component.Render(c.Request().Context(), c.Response())
}

func (h *Handler) Index(c echo.Context) error {
	games, err := h.models.Games.GetAll()
	if err != nil {
		log.Println(err)
		return TemplRender(c, views.Index(nil))
	}

	return TemplRender(c, views.Index(games))
}

func (h *Handler) Flashcard(c echo.Context) error {
	var roles []*data.ComplexRole
	roles, err := h.models.Roles.GetAllComplex()
	if err != nil {
		log.Println(err)
		return TemplRender(c, views.Flashcard(c, nil))
	}
	acceptedRoles := []*data.ComplexRole{}
	for _, role := range roles {
		for _, potentialRole := range potentialRoles {
			if role.Name == potentialRole {
				acceptedRoles = append(acceptedRoles, role)
			}
		}
	}

	return TemplRender(c, views.Flashcard(c, acceptedRoles))
}

func (h *Handler) Search(c echo.Context) error {
	var roles []*data.ComplexRole
	search := c.FormValue("search")
	roles, err := h.models.Roles.GetAllComplex()
	if err != nil {
		log.Println(err)
		return TemplRender(c, views.Search(c, nil))
	}

	// log.Println(search, matchingRoleNames)
	activeGameRoles := []*data.ComplexRole{}
	for _, roleName := range potentialRoles {
		for _, dbRole := range roles {
			if dbRole.Name == roleName {
				activeGameRoles = append(activeGameRoles, dbRole)
			}
		}
	}

	bestMatches := []*data.ComplexRole{}
	for _, currRole := range activeGameRoles {
		abilityNames := []string{}
		passiveNames := []string{}
		for _, a := range currRole.Abilities {
			abilityNames = append(abilityNames, a.Name)
		}
		for _, p := range currRole.Passives {
			passiveNames = append(passiveNames, p.Name)
		}
		matchingRoleNames := fuzzy.FindFold(search, []string{currRole.Name})
		matchingAbilityNames := fuzzy.FindFold(search, abilityNames)
		matchingPassiveNames := fuzzy.FindFold(search, passiveNames)
		matchingAlignmentName := fuzzy.FindFold(search, []string{currRole.Alignment})
		roleMatch := len(matchingRoleNames) > 0
		abilityMatch := len(matchingAbilityNames) > 0
		passiveMatch := len(matchingPassiveNames) > 0
		alignmentMatch := len(matchingAlignmentName) > 0

		if roleMatch || abilityMatch || passiveMatch || alignmentMatch {
			bestMatches = append(bestMatches, currRole)
		}
	}

	return TemplRender(c, views.Search(c, bestMatches))
}

func (h *Handler) Auth(c echo.Context) error {
	pw := c.FormValue("password")

	cookie := &http.Cookie{
		Name:  "password",
		Value: pw,
		Path:  "/",
	}
	c.SetCookie(cookie)

	return nil
}
