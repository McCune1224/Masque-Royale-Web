CREATE TABLE IF NOT EXISTS action(
  id serial PRIMARY KEY,
  ability_details_id INT REFERENCES ability_details (id),
  player_id INT REFERENCES players (id),
  game_id INT REFERENCES games (id),
  pending_approval bool NOT NULL DEFAULT true,
  resolved bool NOT NULL DEFAULT false,
  target VARCHAR(64) NOT NULL,
  context TEXT NOT NULL
);
