CREATE TYPE alignment AS ENUM ('LAWFUL', 'OUTLANDER', 'CHAOTIC');
CREATE TYPE rarity AS ENUM ('COMMON', 'UNCOMMON', 'RARE', 'EPIC', 'LEGENDARY', 'MYTHICAL', 'ROLE_SPECIFIC');


CREATE TABLE IF NOT EXISTS roles (
id serial PRIMARY KEY,
name VARCHAR(64) UNIQUE NOT NULL,
alignment alignment NOT NULL
); 


CREATE TABLE IF NOT EXISTS categories (
id serial PRIMARY KEY,
name VARCHAR(64),
priority int
);

CREATE TABLE IF NOT EXISTS ability_details (
id serial PRIMARY KEY,
name VARCHAR(64) UNIQUE NOT NULL,
description TEXT NOT NULL,
default_charges int,
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

CREATE TABLE IF NOT EXISTS role_abilites_join(
  role_id INT NOT NULL,
  ability_id INT NOT NULL,
  PRIMARY KEY (role_id, ability_id)
);

CREATE TABLE IF NOT EXISTS role_passives_join(
  role_id INT NOT NULL,
  passive_id INT NOT NULL,
  PRIMARY KEY (role_id, passive_id)
);


-- FIXME: There's no uniform system for categories, so for right now I'm just hardcoding the name and associated priority here
INSERT INTO categories 
(name, priority) 
VALUES 
('Alteration', 1),
('Reactive', 2),
('Redirection', 3),
('Vote Redirection', 3),
('Investigation', 4),
('Protection', 5),
('Visit', 6), 
('Blocking', 6),
('Vote', 7), 
('Immunity', 7),
('Vote Immunity', 7),
('Vote Manipulation', 8),
('Support', 9),
('Debuff', 10),
('Theft', 11),
('Healing', 12),
('Destruction', 13),
('Killing', 14);


-- WARNING: This is so infrequently used that it is not worth creating a create for it, but it is used in the game so it needs to be here
INSERT INTO status_details (name, description) VALUES
('Cursed', 'If it isn’t removed within three cycles, you will die.'),
('Empowered', '"You can use any one of your abilities for 2 nights, even if you had previously ran out of charges. Does not consume a charge on use, but still requires an action to use. 
You can only use one AA or Ability with zero charges per cycle. Upon using a killing ability while Empowered, you lose the Empowered status."'),
('Drunk', '25% chance to target a random person instead when using an item or an ability. Wears off after two nights. This effect stacks. (Drunk one stack=25%, Drunk two stack=50%, etc.)'),
('Restrained', 'Cannot use abilities. Permanent until cured.'),
('Blackmailed', 'You can’t talk or vote at that Elimination Phase. Removes after a cycle.'),
('Despaired', 'You vote for yourself at Elimination. Permanent until cured.'),
('Madness', 'When inflicted with madness you must make efforts to present yourself as the role you’ve been made mad about. Anything deviating from that will count as breaking madness, and will result in death. This status lasts two nights unless otherwise stated.');

