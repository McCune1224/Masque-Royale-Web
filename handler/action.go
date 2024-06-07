package handler

import (
	"log"
	"strconv"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/labstack/echo/v4"
	"github.com/mccune1224/betrayal-widget/models"
	"github.com/mccune1224/betrayal-widget/util"
)

func (h *Handler) GetAllActions(c echo.Context) error {
	gameId, err := util.ParseInt32Param(c, "game_id")
	if err != nil {
		return util.BadRequestJson(c, "Invalid Game ID")
	}
	q := models.New(h.Db)

	if c.QueryParam("round") != "" {
		round, err := strconv.Atoi(c.QueryParam("round"))
		if err != nil {
			return util.BadRequestJson(c, "Invalid Round")
		}
		actions, err := q.ListActionsByRoundForGame(c.Request().Context(), models.ListActionsByRoundForGameParams{
			GameID: pgtype.Int4{Int32: gameId, Valid: true},
			Round:  int32(round),
		})
		if err != nil {
			return util.InternalServerErrorJson(c, err.Error())
		}
		return c.JSON(200, actions)
	}

	if c.QueryParam("player_id") != "" {
		playerId, err := strconv.Atoi(c.QueryParam("player_id"))
		if err != nil {
			return util.BadRequestJson(c, "Invalid Player ID")
		}
		actions, err := q.ListActionsByPlayer(c.Request().Context(), int32(playerId))
		if err != nil {
			return util.InternalServerErrorJson(c, err.Error())
		}
		return c.JSON(200, actions)
	}

	actions, err := q.ListActionsByGame(c.Request().Context(), pgtype.Int4{Int32: gameId, Valid: true})
	if err != nil {
		return util.InternalServerErrorJson(c, err.Error())
	}
	return c.JSON(200, actions)
}

func (h *Handler) InsertAction(c echo.Context) error {
	var action models.Action
	err := c.Bind(&action)
	if err != nil {
		log.Println(err)
		return util.BadRequestJson(c, err.Error())
	}
	log.Println(action)
	q := models.New(h.Db)
	action, err = q.CreateAction(c.Request().Context(), models.CreateActionParams{
		GameID:          action.GameID,
		PlayerID:        action.PlayerID,
		PendingApproval: action.PendingApproval,
		Resolved:        action.Resolved,
		Target:          action.Target,
		Context:         action.Context,
		AbilityName:     action.AbilityName,
	})
	if err != nil {
		log.Println(err)
		return util.InternalServerErrorJson(c, err.Error())
	}
	return c.JSON(200, action)
}
