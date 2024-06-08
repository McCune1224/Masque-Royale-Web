CREATE TABLE IF NOT EXISTS players(
id serial PRIMARY KEY,
name VARCHAR(64) UNIQUE NOT NULL,
game_id INT REFERENCES games (id) ON DELETE CASCADE,
role_id INT REFERENCES roles (id),
alive bool NOT NULL,
alignment_override VARCHAR(64),
notes TEXT NOT NULL,
room_id INT REFERENCES rooms (id)
);


CREATE TABLE IF NOT EXISTS player_inventories(
player_id serial UNIQUE NOT NULL ,
ability_name VARCHAR(64) UNIQUE NOT NULL,
ability_quantity int,
PRIMARY KEY(player_id, ability_name)
);


CREATE TABLE IF NOT EXISTS abilities(
id serial PRIMARY KEY,
ability_details_id int REFERENCES ability_details (id),
player_inventory_id int REFERENCES player_inventories (player_id)
);
