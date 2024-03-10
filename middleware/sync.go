package middleware

import (
	"log"

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
		diff := util.BulkCalculateLuck(players)
		for i := range players {
			if players[i].P.Luck != diff[i].P.Luck {
				err := s.models.Players.Update(&diff[i].P)
				if err != nil {
					log.Println(err)
					return c.Redirect(302, "/")
				}
			}
		}
		players = util.OrderComplexPlayers(players)
		// pc := util.PlayerContext{Context: c}
		// pc.SetPlayers(players)
		c.Set("players", players)
		return next(c)
	}
}
