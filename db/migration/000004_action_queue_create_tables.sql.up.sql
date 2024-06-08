CREATE TABLE IF NOT EXISTS actions(
  id serial PRIMARY KEY,
  game_id INT REFERENCES games (id) ON DELETE CASCADE,
  player_id INT REFERENCES players (id) ON DELETE CASCADE,
  pending_approval bool NOT NULL DEFAULT true,
  resolved bool NOT NULL DEFAULT false,
  target VARCHAR(64) NOT NULL,
  context TEXT NOT NULL,
  ability_name VARCHAR(64) NOT NULL,
  round INT NOT NULL,
  priority INT NOT NULL,
  role_id INT REFERENCES roles (id)
);
