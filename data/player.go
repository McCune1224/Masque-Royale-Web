package data

type Player struct {
	ID        int    `db:"id"`
	Name      string `db:"name"`
	RoleID    string `db:"role_id"`
	Alive     bool   `db:"alive"`
	GameID    string `db:"game_id"`
	CreatedAt string `db:"created_at"`
	UpdatedAt string `db:"updated_at"`
}
