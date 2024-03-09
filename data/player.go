package data

import (
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

type Player struct {
	ID           int    `db:"id"`
	Name         string `db:"name"`
	GameID       string `db:"game_id"`
	RoleID       int    `db:"role_id"`
	Alive        bool   `db:"alive"`
	Seat         int    `db:"seat"`
	Luck         int    `db:"luck"`
	LuckModifier int    `db:"luck_modifier"`
	CreatedAt    string `db:"created_at"`
	UpdatedAt    string `db:"updated_at"`
}

// ComplexPlayer is a player with a role
type ComplexPlayer struct {
	P Player `db:"players"`
	R Role   `db:"roles"`
}

// WARNING: May god have mercy on my soul for this abomination
type playerRoleJoin struct {
	PlayerID       int           `db:"player_id"`
	PlayerName     string        `db:"player_name"`
	PlayerGameID   string        `db:"player_game_id"`
	PlayerRoleID   int           `db:"player_role_id"`
	PlayerAlive    bool          `db:"player_alive"`
	PlayerSeat     int           `db:"player_seat"`
	PlayerLuck     int           `db:"player_luck"`
	PlayerLuckMod  int           `db:"player_luck_modifier"`
	PlayerCreated  string        `db:"player_created_at"`
	PlayerUpdated  string        `db:"player_updated_at"`
	RoleID         int           `db:"role_id"`
	RoleName       string        `db:"role_name"`
	RoleAlignment  string        `db:"role_alignment"`
	RoleAbilityIDs pq.Int32Array `db:"role_ability_ids"`
	RolePassiveIDs pq.Int32Array `db:"role_passive_ids"`
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

func (m *PlayerModel) GetByID(id int) (*Player, error) {
	player := &Player{}
	err := m.DB.Get(player, "SELECT * FROM players WHERE id = $1", id)
	if err != nil {
		return nil, err
	}
	return player, nil
}

func (m *PlayerModel) GetByName(name string) (*Player, error) {
	player := &Player{}
	err := m.DB.Get(player, "SELECT * FROM players WHERE name ILIKE $1", name)
	if err != nil {
		return nil, err
	}
	return player, nil
}

func (m *PlayerModel) GetByGameIDAndName(gameID string, name string) (*Player, error) {
	player := &Player{}
	err := m.DB.Get(player, "SELECT * FROM players WHERE game_id = $1 AND name ILIKE $2", gameID, name)
	if err != nil {
		return nil, err
	}
	return player, nil
}

func (m *PlayerModel) Create(player *Player) error {
	_, err := m.DB.NamedExec("INSERT INTO players (name, game_id, role_id, alive, seat, luck, luck_modifier) VALUES (:name, :game_id, :role_id, :alive, :seat, :luck, :luck_modifier)", player)
	if err != nil {
		return err
	}
	return nil
}

func (m *PlayerModel) Update(player *Player) error {
	_, err := m.DB.NamedExec("UPDATE players SET name = :name, game_id = :game_id, role_id = :role_id, alive = :alive, seat = :seat, luck = :luck, luck_modifier = :luck_modifier WHERE id = :id", player)
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

func (m *PlayerModel) GetRole(roleID int) (*Role, error) {
	role := &Role{}
	err := m.DB.Get(role, "SELECT * FROM roles WHERE id = $1", roleID)
	if err != nil {
		return nil, err
	}
	return role, nil
}

func (m *PlayerModel) GetComplexByGameID(gameID string) ([]*ComplexPlayer, error) {
	playerQuery := []*playerRoleJoin{}
	players := []*ComplexPlayer{}
	query := `SELECT p.id AS player_id, p.name AS player_name, p.game_id AS player_game_id, p.role_id AS player_role_id, p.alive AS player_alive, p.seat AS player_seat, p.luck AS player_luck, p.luck_modifier AS player_luck_modifier, p.created_at AS player_created_at, p.updated_at AS player_updated_at, r.id AS role_id, r.name AS role_name, r.alignment AS role_alignment, r.ability_ids AS role_ability_ids, r.passive_ids AS role_passive_ids FROM players p JOIN roles r ON p.role_id = r.id WHERE p.game_id = $1`
	err := m.DB.Select(&playerQuery, query, gameID)
	if err != nil {
		return nil, err
	}

	for _, p := range playerQuery {
		player := &Player{
			ID:           p.PlayerID,
			Name:         p.PlayerName,
			GameID:       p.PlayerGameID,
			RoleID:       p.PlayerRoleID,
			Alive:        p.PlayerAlive,
			Seat:         p.PlayerSeat,
			Luck:         p.PlayerLuck,
			LuckModifier: p.PlayerLuckMod,
			CreatedAt:    p.PlayerCreated,
			UpdatedAt:    p.PlayerUpdated,
		}
		role := &Role{
			ID:         p.RoleID,
			Name:       p.RoleName,
			Alignment:  p.RoleAlignment,
			AbilityIDs: p.RoleAbilityIDs,
			PassiveIDs: p.RolePassiveIDs,
		}
		players = append(players, &ComplexPlayer{
			P: *player,
			R: *role,
		})
	}

	return players, nil
}
