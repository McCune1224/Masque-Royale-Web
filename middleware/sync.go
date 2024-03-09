package middleware

import (
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/mccune1224/betrayal-widget/data"
	"github.com/mccune1224/betrayal-widget/util"
)

type SyncMiddleware struct {
	models *data.Models
}

func NewSyncMiddleware(db *sqlx.DB) *SyncMiddleware {
	models := data.NewModels(db)
	return &SyncMiddleware{
		models: models,
	}
}

func (s *SyncMiddleware) SyncPlayerInfo(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if c.Param("game_id") == "" {
			return next(c)
		}
		players, err := s.models.Players.GetAllComplexByGameID(c.Param("game_id"))
		if err != nil {
			return err
		}
		players = util.OrderComplexPlayers(players)
		// pc := util.PlayerContext{Context: c}
		// pc.SetPlayers(players)
		c.Set("players", players)
		return next(c)
	}
}
