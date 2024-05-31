CREATE TYPE alignment AS ENUM ('LAWFUL', 'OUTLANDER', 'CHAOTIC');
CREATE TYPE rarity AS ENUM ('COMMON', 'UNCOMMON', 'RARE', 'EPIC', 'LEGENDARY', 'MYTHICAL', 'ROLE_SPECIFIC');


CREATE TABLE IF NOT EXISTS roles(
id serial PRIMARY KEY,
name VARCHAR(64) UNIQUE NOT NULL,
alignment alignment NOT NULL,
ability_ids INT[] DEFAULT '{}',
passive_ids INT[] DEFAULT '{}'
); 


CREATE TABLE IF NOT EXISTS categories (
id serial PRIMARY KEY,
name VARCHAR(64),
priority int
);

CREATE TABLE IF NOT EXISTS ability_details(
id serial PRIMARY KEY,
name VARCHAR(64) UNIQUE NOT NULL,
description TEXT NOT NULL,
default_charges int,
role_id int REFERENCES roles (id),
category_ids INT[] DEFAULT '{}',
rarity rarity NOT NULL,
any_ability bool
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


CREATE TABLE IF NOT EXISTS status_details(
id serial PRIMARY KEY,
name VARCHAR(64) NOT NULL,
description TEXT
);


CREATE TABLE IF NOT EXISTS passive_details(
id serial PRIMARY KEY,
name VARCHAR(64) UNIQUE NOT NULL,
description TEXT NOT NULL
);

