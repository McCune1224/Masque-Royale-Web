package handler

import (
	"errors"
	"log"

	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5/pgconn"
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
		log.Println("AM I HERE")
		return models.CreatePlayerParams{}, err
	}

	var params models.CreatePlayerParams
	if err := c.Bind(pcv); err != nil {
		log.Println(err.Error())
		return params, err
	}

	params = models.CreatePlayerParams{
		Name:   pcv.Name,
		GameID: pgtype.Int4{Int32: gameId, Valid: true},
		RoleID: pgtype.Int4{Int32: pcv.RoleID, Valid: true},
		Alive:  true,
		RoomID: pgtype.Int4{Int32: pcv.RoomID, Valid: true},
	}

	return params, nil
}

func (h *Handler) InsertPlayer(c echo.Context) error {
	pcv := new(playerCreateValidator)
	dbPlayerCreate, err := pcv.Validate(c)
	if err != nil {
		log.Println(err)
		return util.BadRequestJson(c, err.Error())
	}
	q := models.New(h.Db)

	role, err := q.GetRole(c.Request().Context(), dbPlayerCreate.RoleID.Int32)
	if err != nil {
		return util.InternalServerErrorJson(c, err.Error())
	}

	dbPlayerCreate.Alignment = models.NullAlignment{
		Alignment: role.Alignment,
		Valid:     true,
	}

	player, err := q.CreatePlayer(c.Request().Context(), dbPlayerCreate)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			switch pgErr.Code {
			case pgerrcode.UniqueViolation:
				return util.BadRequestJson(c, "Player already exists")
			default:
				return util.InternalServerErrorJson(c, err.Error())
			}

		}
		log.Println(err)
		return util.InternalServerErrorJson(c, err.Error())
	}

	// create player notes
	_, err = q.UpsertPlayerNote(c.Request().Context(), models.UpsertPlayerNoteParams{
		PlayerID: player.ID,
		Note:     "...",
	})
	if err != nil {
		log.Println(err)
		return util.InternalServerErrorJson(c, err.Error())
	}

	// create player abilities
	roleAbilities, err := q.GetRoleAbilityDetails(c.Request().Context(), role.ID)
	if err != nil {
		log.Println(err)
		return util.InternalServerErrorJson(c, err.Error())
	}
	for _, roleAbility := range roleAbilities {
		_, err = q.CreatePlayerAbility(c.Request().Context(), models.CreatePlayerAbilityParams{
			PlayerID:         player.ID,
			AbilityDetailsID: roleAbility.ID,
			Charges:          roleAbility.DefaultCharges.Int32,
		})
		if err != nil {
			log.Println(err)
			return util.InternalServerErrorJson(c, err.Error())
		}
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

func (h *Handler) GetAllGamePlayers(c echo.Context) error {
	gameId, err := util.ParseInt32Param(c, "game_id")
	if err != nil {
		return util.BadRequestJson(c, "Invalid Game ID")
	}
	q := models.New(h.Db)
	players, err := q.ListPlayersByGame(c.Request().Context(), pgtype.Int4{Int32: gameId, Valid: true})

	// if err != nil {
	// 	var pgErr *pgconn.PgError
	// 	if errors.As(err, &pgErr) {
	// 		switch pgerrcode.(pgErr.Code) {
	//
	// 		}
	// 	}
	// 	return util.InternalServerErrorJson(c, err.Error())
	// }

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
		ID:     player.ID,
		Name:   player.Name,
		GameID: player.GameID,
		RoleID: player.RoleID,
		Alive:  player.Alive,
		RoomID: player.RoomID,
	})
	if err != nil {
		log.Println(err)
		return util.InternalServerErrorJson(c, err.Error())
	}
	return c.JSON(200, player)
}

func (h *Handler) GetPlayerNotes(c echo.Context) error {
	playerID, err := util.ParseInt32Param(c, "player_id")
	if err != nil {
		return util.BadRequestJson(c, "Invalid Player ID")
	}
	q := models.New(h.Db)
	notes, err := q.GetPlayerNote(c.Request().Context(), playerID)
	if err != nil {
		log.Println(err)
		return util.InternalServerErrorJson(c, err.Error())
	}
	return c.JSON(200, notes)
}

func (h *Handler) GetPlayerAbilities(c echo.Context) error {
	playerID, err := util.ParseInt32Param(c, "player_id")
	if err != nil {
		return util.BadRequestJson(c, "Invalid Player ID")
	}
	q := models.New(h.Db)
	abilities, err := q.ListPlayerAbilitiesJoin(c.Request().Context(), playerID)
	if err != nil {
		log.Println(err)
		return util.InternalServerErrorJson(c, err.Error())
	}
	return c.JSON(200, abilities)
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
