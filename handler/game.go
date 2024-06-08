package handler

import (
	"log"

	"github.com/jackc/pgerrcode"
	"github.com/labstack/echo/v4"
	"github.com/mccune1224/betrayal-widget/models"
	"github.com/mccune1224/betrayal-widget/util"
)

func (h *Handler) GetRandomGame(c echo.Context) error {
	q := models.New(h.Db)
	game, err := q.GetRandomGame(c.Request().Context())
	if err != nil {
		return util.InternalServerErrorJson(c, err.Error())
	}
	return c.JSON(200, game)
}

func (h *Handler) GetAllGames(c echo.Context) error {
	q := models.New(h.Db)
	games, err := q.ListGames(c.Request().Context())
	if err != nil {
		return util.InternalServerErrorJson(c, err.Error())
	}
	return c.JSON(200, games)
}

func (h *Handler) GetGameByID(c echo.Context) error {
	gameId, err := util.ParseInt32Param(c, "game_id")
	if err != nil {
		return util.BadRequestJson(c, "Invalid Game ID")
	}

	q := models.New(h.Db)
	game, err := q.GetGame(c.Request().Context(), int32(gameId))
	if err != nil {
		return util.InternalServerErrorJson(c, err.Error())
	}
	return c.JSON(200, game)
}

func (h *Handler) InsertGame(c echo.Context) error {

	gm := &models.CreateGameParams{}
	err := c.Bind(gm)
	if err != nil {
		return util.BadRequestJson(c, "Invalid game name")
	}

	if gm.Name == "" {
		util.BadRequestJson(c, "Missing game name")
	}

	//default case if day is not provided
	if gm.Phase == "" {
		gm.Phase = "Day"
	}

	q := models.New(h.Db)
	game, err := q.CreateGame(c.Request().Context(), *gm)
	if err != nil {
		if util.ErrorContains(err, pgerrcode.UniqueViolation) {
			util.BadRequestJson(c, "Game name already exists")
		} else {
			log.Println(err)
			return util.InternalServerErrorJson(c, err.Error())
		}
	}

	return c.JSON(200, game)

}

func (h *Handler) DeleteGame(c echo.Context) error {
	gameId, err := util.ParseInt32Param(c, "game_id")
	if err != nil {
		return util.BadRequestJson(c, "Invalid Game ID")
	}
	q := models.New(h.Db)
	err = q.DeleteGame(c.Request().Context(), int32(gameId))
	if err != nil {
		return util.InternalServerErrorJson(c, err.Error())
	}
	return c.JSON(200, "Success")
}

func (h *Handler) UpdateGame(c echo.Context) error {
	var game models.Game
	err := c.Bind(&game)
	if err != nil {
		log.Println(err)
		return util.BadRequestJson(c, err.Error())
	}
	q := models.New(h.Db)
	game, err = q.UpdateGame(c.Request().Context(), models.UpdateGameParams{
		ID:        game.ID,
		Name:      game.Name,
		Phase:     game.Phase,
		Round:     game.Round,
		PlayerIds: game.PlayerIds,
	})
	if err != nil {
		log.Println(err)
		return util.InternalServerErrorJson(c, err.Error())
	}
	return c.JSON(200, game)
}
