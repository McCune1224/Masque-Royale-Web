package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/mccune1224/betrayal-widget/models"
	"github.com/mccune1224/betrayal-widget/util"
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
	roomID, err := util.ParseInt32Param(c, "room_id")
	if err != nil {
		return util.BadRequestJson(c, "Invalid Room ID")
	}
	q := models.New(h.Db)
	room, err := q.GetRoom(c.Request().Context(), roomID)
	if err != nil {
		return util.InternalServerErrorJson(c, err.Error())
	}

	return c.JSON(200, room)
}

func (h *Handler) GetRoomByName(c echo.Context) error {
	return c.JSON(200, "Success")
}
