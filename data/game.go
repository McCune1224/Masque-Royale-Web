package data

import (
	"github.com/jmoiron/sqlx"
)

type Game struct {
	Id          int    `db:"id"`
	Name        string `db:"name"`
	Phase       string `db:"phase"`
	Round       int    `db:"round"`
	PlayerCount int    `db:"player_count"`
	CreatedAt   string `db:"created_at"`
	UpdatedAt   string `db:"updated_at"`
}

type GameModel struct {
	*sqlx.DB
}

func (gm *GameModel) GetRandomGame() (*Game, error) {
	game := &Game{}
	err := gm.Get(game, "SELECT * FROM games ORDER BY random() LIMIT 1")
	if err != nil {
		return nil, err
	}
	return game, err
}

func (gm *GameModel) GetAllGames() ([]*Game, error) {
	games := []*Game{}
	err := gm.Select(&games, "SELECT * from games")
	if err != nil {
		return nil, err
	}
	return games, nil
}

func (gm *GameModel) GetGameByID(id int) (*Game, error) {
	game := &Game{}
	err := gm.Get(game, "SELECT * FROM games WHERE id = $1", id)
	if err != nil {
		return nil, err
	}
	return game, nil
}
