CREATE TABLE IF NOT EXISTS players(
id serial PRIMARY KEY,
name VARCHAR(64) UNIQUE NOT NULL,
game_id INT REFERENCES games (id) ON DELETE CASCADE,
role_id INT REFERENCES roles (id),
alive bool NOT NULL,
alignment alignment,
room_id INT REFERENCES rooms (id)
);

CREATE TABLE IF NOT EXISTS player_notes(
player_id serial REFERENCES players (id) ON DELETE CASCADE,
note TEXT NOT NULL,
PRIMARY KEY(player_id)
);

CREATE TABLE IF NOT EXISTS player_abilities(
player_id serial REFERENCES players (id) ON DELETE CASCADE,
ability_details_id serial REFERENCES ability_details (id),
charges int NOT NULL DEFAULT 1,
PRIMARY KEY(player_id, ability_details_id)
);

CREATE TABLE IF NOT EXISTS player_statuses(
player_id serial REFERENCES players (id) ON DELETE CASCADE,
status_id serial REFERENCES status_details (id),
stack int NOT NULL DEFAULT 0,
round_given int NOT NULL,
PRIMARY KEY(player_id, status_id)
);
