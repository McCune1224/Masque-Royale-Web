package data

import (
	"github.com/jmoiron/sqlx"
)

type Models struct {
	Games   GameModel
	Players PlayerModel
}

func NewModels(db *sqlx.DB) *Models {
	return &Models{
		Games:   GameModel{DB: db},
		Players: PlayerModel{DB: db},
	}
}
