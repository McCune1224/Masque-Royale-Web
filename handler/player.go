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
	playerID, err := util.ParseInt32Param(c, "player_id")
	if err != nil {
		return util.BadRequestJson(c, "Invalid Player ID")
	}
	q := models.New(h.Db)
	player, err := q.GetPlayer(c.Request().Context(), playerID)
	if err != nil {
		log.Println(err)
	}

	return c.JSON(200, player)
}

func (h *Handler) GetAllPlayers(c echo.Context) error {
	gameId, err := util.ParseInt32Param(c, "game_id")
	if err != nil {
		return util.BadRequestJson(c, "Invalid Game ID")
	}
	q := models.New(h.Db)
	players, err := q.ListPlayersByGame(c.Request().Context(), pgtype.Int4{Int32: gameId, Valid: true})
	if err != nil {
		return util.InternalServerErrorJson(c, err.Error())
	}

	return c.JSON(200, players)
}

func (h *Handler) UpdatePlayer(c echo.Context) error {
	var player models.Player
	err := c.Bind(&player)
	if err != nil {
		log.Println(err)
		return util.BadRequestJson(c, err.Error())
	}
	q := models.New(h.Db)
	player, err = q.UpdatePlayer(c.Request().Context(), models.UpdatePlayerParams{
		ID:                player.ID,
		Name:              player.Name,
		GameID:            player.GameID,
		RoleID:            player.RoleID,
		Alive:             player.Alive,
		AlignmentOverride: player.AlignmentOverride,
		Notes:             player.Notes,
		RoomID:            player.RoomID,
	})
	if err != nil {
		log.Println(err)
		return util.InternalServerErrorJson(c, err.Error())
	}
	return c.JSON(200, player)
}

func (h *Handler) DeletePlayer(c echo.Context) error {
	playerID, err := util.ParseInt32Param(c, "player_id")
	if err != nil {
		return util.BadRequestJson(c, "Invalid Player ID")
	}
	q := models.New(h.Db)
	err = q.DeletePlayer(c.Request().Context(), playerID)
	if err != nil {
		log.Println(err)
		return util.InternalServerErrorJson(c, err.Error())
	}
	return c.JSON(200, "Success")
}

func (h *Handler) GetPlayerActions(c echo.Context) error {
	// gameID, err := util.ParseInt32Param(c, "game_id")
	// if err != nil {
	// 	return util.BadRequestJson(c, "Invalid Game ID")
	// }
	playerID, err := util.ParseInt32Param(c, "player_id")
	if err != nil {
		return util.BadRequestJson(c, "Invalid Player ID")
	}
	q := models.New(h.Db)
	actions, err := q.ListActionsByPlayer(c.Request().Context(), playerID)
	if err != nil {
		return util.InternalServerErrorJson(c, err.Error())
	}
	return c.JSON(200, actions)
}
