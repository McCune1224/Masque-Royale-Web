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
		return c.JSON(500, echo.Map{"message": err.Error()})
	}
	return c.JSON(200, game)
}

func (h *Handler) GetAllGames(c echo.Context) error {
	q := models.New(h.Db)
	games, err := q.ListGames(c.Request().Context())
	if err != nil {
		return c.JSON(500,
			echo.Map{"message": err.Error()},
		)
	}
	return c.JSON(200, games)
}

func (h *Handler) GetGameByID(c echo.Context) error {
	gameId, err := util.ParseInt32Param(c, "game_id")
	if err != nil {
		return c.JSON(400, echo.Map{"message": "Invalid Game ID"})
	}

	q := models.New(h.Db)
	game, err := q.GetGame(c.Request().Context(), int32(gameId))
	if err != nil {
		return c.JSON(500, echo.Map{"message": err.Error()})
	}

	return c.JSON(200, game)
}

func (h *Handler) InsertGame(c echo.Context) error {

	gm := &models.CreateGameParams{}
	err := c.Bind(gm)
	if err != nil {
		return c.JSON(400, echo.Map{"message": "Invalid Game Name"})
	}

	if gm.Name == "" {
		return c.JSON(400, echo.Map{"message": "Missing Game Name"})
	}

	q := models.New(h.Db)
	game, err := q.CreateGame(c.Request().Context(), *gm)
	if err != nil {
		pgerr := util.ParsePgError(err)
		if pgerr != nil {
			switch pgerr.Code {
			case pgerrcode.UniqueViolation:
				c.JSON(400, echo.Map{"message:": "Game Name already exists"})
			default:
				log.Println(err)
				return c.JSON(500, echo.Map{"error": err.Error()})
			}

		} else {
			return c.JSON(500, echo.Map{"error": err.Error()})
		}

	}
	return c.JSON(200, game)

}

// func (h *Handler) UpdateGame(c echo.Context) error {}
