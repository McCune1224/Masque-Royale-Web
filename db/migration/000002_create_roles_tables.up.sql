DO $$
BEGIN
IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'alignment') THEN
CREATE TYPE alignment AS ENUM ('LAWFUL', 'OUTLANDER', 'CHAOTIC');
END IF;
END $$;
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
role_id int REFERENCES roles (id),
category_ids INT[] DEFAULT '{}',
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

