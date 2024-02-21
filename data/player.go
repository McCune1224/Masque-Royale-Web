package data

import "github.com/jmoiron/sqlx"

type Player struct {
	ID        int    `db:"id"`
	Name      string `db:"name"`
	GameID    string `db:"game_id"`
	RoleID    int    `db:"role_id"`
	Alive     bool   `db:"alive"`
	Seat      int    `db:"seat"`
	CreatedAt string `db:"created_at"`
	UpdatedAt string `db:"updated_at"`
}

type PlayerModel struct {
	DB *sqlx.DB
}

func (m *PlayerModel) GetByGameID(gameID string) ([]*Player, error) {
	players := []*Player{}
	err := m.DB.Select(&players, "SELECT * FROM players WHERE game_id = $1", gameID)
	if err != nil {
		return nil, err
	}
	return players, nil
}

func (m *PlayerModel) Create(player *Player) error {
	_, err := m.DB.NamedExec("INSERT INTO players (name, game_id, role_id, alive, seat) VALUES (:name, :game_id, :role_id, :alive, :seat)", player)
	if err != nil {
		return err
	}
	return nil
}

func (m *PlayerModel) Update(player *Player) error {
	_, err := m.DB.NamedExec("UPDATE players SET name = :name, game_id = :game_id, role_id = :role_id, alive = :alive, seat = :seat WHERE id = :id", player)
	if err != nil {
		return err
	}
	return nil
}

func (m *PlayerModel) Delete(id int) error {
	_, err := m.DB.Exec("DELETE FROM players WHERE id = $1", id)
	if err != nil {
		return err
	}
	return nil
}
