package data

import "github.com/jmoiron/sqlx"

type Game struct {
	ID          int    `db:"id"`
	Name        string `db:"name"`
	PlayerCount int    `db:"player_count"`
	// Phase       string `db:"phase"`
	// Round       int    `db:"round"`
	CreatedAt string `db:"created_at"`
	UpdatedAt string `db:"updated_at"`
}

type GameModel struct {
	*sqlx.DB
}
