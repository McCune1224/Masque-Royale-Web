package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/mccune1224/betrayal-widget/models"
)

func (h *Handler) GetAllRooms(c echo.Context) error {

	q := models.New(h.Db)
	rooms, err := q.ListRooms(c.Request().Context())
	if err != nil {
		return echo.NewHTTPError(500, err.Error())
	}

	return c.JSON(200, rooms)
}

func (h *Handler) GetRoomByID(c echo.Context) error {
	return c.JSON(200, "Success")
}

func (h *Handler) GetRoomByName(c echo.Context) error {
	return c.JSON(200, "Success")
}
