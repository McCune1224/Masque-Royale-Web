CREATE TABLE IF NOT EXISTS actions(
  id serial PRIMARY KEY,
  game_id INT REFERENCES games (id),
  player_id INT REFERENCES players (id),
  pending_approval bool NOT NULL DEFAULT true,
  resolved bool NOT NULL DEFAULT false,
  target VARCHAR(64) NOT NULL,
  context TEXT NOT NULL,
  ability_name VARCHAR(64) NOT NULL,
  role_id INT REFERENCES roles (id)
);
