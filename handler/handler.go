package handler

import (
	"github.com/a-h/templ"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/mccune1224/betrayal-widget/data"
)

var CurrentGameRoles = []string{
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
