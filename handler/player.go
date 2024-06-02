package handler

import (
	"log"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/labstack/echo/v4"
	"github.com/mccune1224/betrayal-widget/models"
	"github.com/mccune1224/betrayal-widget/util"
)

type playerCreateValidator struct {
	Name   string `json:"name" validate:"required"`
	RoleID int32  `json:"role_id" validate:"required"`
	RoomID int32  `json:"room_id" validate:"required"`
}

func (pcv *playerCreateValidator) Validate(c echo.Context) (models.CreatePlayerParams, error) {
	gameId, err := util.ParseInt32Param(c, "game_id")
	if err != nil {
		return models.CreatePlayerParams{}, err
	}

	var params models.CreatePlayerParams
	if err := c.Bind(pcv); err != nil {
		log.Println(err.Error())
		return params, err
	}
	params = models.CreatePlayerParams{
		Name:              pcv.Name,
		GameID:            pgtype.Int4{Int32: gameId, Valid: true},
		RoleID:            pgtype.Int4{Int32: pcv.RoleID, Valid: true},
		Alive:             true,
		AlignmentOverride: pgtype.Text{String: "", Valid: true},
		Notes:             "...",
		RoomID:            pgtype.Int4{Int32: pcv.RoomID, Valid: true},
	}

	return params, nil
}

func (h *Handler) InsertPlayer(c echo.Context) error {
	pcv := new(playerCreateValidator)
	dbPlayerCreate, err := pcv.Validate(c)
	if err != nil {
		return util.BadRequestJson(c, err.Error())
	}
	log.Println(dbPlayerCreate)
	q := models.New(h.Db)
	player, err := q.CreatePlayer(c.Request().Context(), dbPlayerCreate)
	if err != nil {
		log.Println(err)
		return util.InternalServerErrorJson(c, err.Error())
	}

	return c.JSON(200, player)
}

func (h *Handler) GetPlayerByID(c echo.Context) error {
	return c.JSON(200, "Success")
}

func (h *Handler) GetAllPlayers(c echo.Context) error {
	return c.JSON(200, "Success")
}
