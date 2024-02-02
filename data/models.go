package data

import (
	"github.com/jmoiron/sqlx"
)

type Models struct {
	Games GameModel
}

func NewModels(db *sqlx.DB) *Models {
	return &Models{
		Games: GameModel{DB: db},
	}
}
