package data

type Player struct {
	ID        int    `db:"id"`
	Name      string `db:"name"`
	GameID    string `db:"game_id"`
	RoleID    string `db:"role_id"`
	Alive     bool   `db:"alive"`
	Seat      int    `db:"seat"`
	CreatedAt string `db:"created_at"`
	UpdatedAt string `db:"updated_at"`
}
