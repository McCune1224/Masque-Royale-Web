package handler

import (
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
	return &Handler{
		models: data.NewModels(db),
	}
}

// Helper to wrap calling a component's Render method.
func TemplRender(c echo.Context, component templ.Component) error {
	return component.Render(c.Request().Context(), c.Response())
}

func (h *Handler) Index(c echo.Context) error {
	cookie, _ := c.Cookie("access_token")
	if cookie != nil {
		return c.Redirect(302, "/dashboard")
	}

	return TemplRender(c, views.Index(false))
}
