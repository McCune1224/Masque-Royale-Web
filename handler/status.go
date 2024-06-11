package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/mccune1224/betrayal-widget/models"
)

func (h *Handler) GetAllStatuses(c echo.Context) error {
	q := models.New(h.Db)
	statuses, err := q.ListStatusDetails(c.Request().Context())
	if err != nil {
		return err
	}
	return c.JSON(200, statuses)
}
