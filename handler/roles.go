package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/mccune1224/betrayal-widget/models"
)

func (h *Handler) GetAllRoles(c echo.Context) error {

	q := models.New(h.Db)

	roles, err := q.ListRoles(c.Request().Context())
	if err != nil {
		return c.JSON(500, err)
	}

	return c.JSON(200, roles)
}
