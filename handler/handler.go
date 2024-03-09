package handler

import (
	"log"

	"github.com/a-h/templ"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/mccune1224/betrayal-widget/data"
	"github.com/mccune1224/betrayal-widget/views"
)

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
