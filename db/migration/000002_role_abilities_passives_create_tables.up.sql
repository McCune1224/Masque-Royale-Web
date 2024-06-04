CREATE TYPE alignment AS ENUM ('LAWFUL', 'OUTLANDER', 'CHAOTIC');
CREATE TYPE rarity AS ENUM ('COMMON', 'UNCOMMON', 'RARE', 'EPIC', 'LEGENDARY', 'MYTHICAL', 'ROLE_SPECIFIC');

CREATE TABLE IF NOT EXISTS roles (
id serial PRIMARY KEY,
name VARCHAR(64) UNIQUE NOT NULL,
alignment alignment NOT NULL
); 

CREATE TABLE IF NOT EXISTS ability_details (
id serial PRIMARY KEY,
name VARCHAR(64) UNIQUE NOT NULL,
description TEXT NOT NULL,
default_charges int,
category_ids INT[] DEFAULT '{}',
rarity rarity NOT NULL,
priority int,
any_ability bool
);

CREATE TABLE IF NOT EXISTS any_ability_details (
id serial PRIMARY KEY,
name VARCHAR(64) UNIQUE NOT NULL,
shorthand VARCHAR(20) UNIQUE NOT NULL,
description TEXT NOT NULL,
category_ids INT[] DEFAULT '{}',
rarity rarity NOT NULL,
priority int
);

CREATE TABLE IF NOT EXISTS passive_details(
id serial PRIMARY KEY,
name VARCHAR(64) UNIQUE NOT NULL,
description TEXT NOT NULL
);


CREATE TABLE IF NOT EXISTS role_abilities_join(
  role_id INT NOT NULL,
  ability_id INT NOT NULL,
  PRIMARY KEY (role_id, ability_id)
);

CREATE TABLE IF NOT EXISTS role_passives_join(
  role_id INT NOT NULL,
  passive_id INT NOT NULL,
  PRIMARY KEY (role_id, passive_id)
);


CREATE TABLE IF NOT EXISTS categories (
id serial PRIMARY KEY,
name VARCHAR(64),
priority int
);


-- FIXME: There's no uniform system for categories, so for right now I'm just
-- hardcoding the name and associated priority here
INSERT INTO categories 
(name, priority) 
VALUES 
('ALTERATION', 1),
('REACTIVE', 2),
('REDIRECTION', 3),
('VISIT REDIRECTION', 3),
('VOTE REDIRECTION', 3),
('INVESTIGATION', 4),
('PROTECTION', 5),
('VISIT', 6), 
('VISITING', 6), 
('VISIT BLOCKING', 6),
('BLOCKING', 6),
('VOTE', 7), 
('IMMUNITY', 7),
('VOTE IMMUNITY', 7),
('VOTE MANIPULATION', 8),
('SUPPORT', 9),
('DEBUFF', 10),
('THEFT', 11),
('HEALING', 12),
('DESTRUCTION', 13),
('KILLING', 14),
('POSITIVE', 20),
('NEGATIVE', 20),
('NEUTRAL', 20),
('NON-VISITING', 20),
('NON VISITING', 20),
('INSTANT', 20),
('NIGHT', 20);


CREATE TABLE IF NOT EXISTS status_details(
id serial PRIMARY KEY,
name VARCHAR(64) NOT NULL,
description TEXT
);


-- WARNING: This is so infrequently used that it is not worth creating a create for
-- it, but it is used in the game so it needs to be here
INSERT INTO status_details (name, description) VALUES
('Cursed', 'If it isn’t removed within three cycles, you will die.'),
('Empowered', '"You can use any one of your abilities for 2 nights, even if you had previously ran out of charges. Does not consume a charge on use, but still requires an action to use. 
You can only use one AA or Ability with zero charges per cycle. Upon using a killing ability while Empowered, you lose the Empowered status."'),
('Drunk', '25% chance to target a random person instead when using an item or an ability. Wears off after two nights. This effect stacks. (Drunk one stack=25%, Drunk two stack=50%, etc.)'),
('Restrained', 'Cannot use abilities. Permanent until cured.'),
('Blackmailed', 'You can’t talk or vote at that Elimination Phase. Removes after a cycle.'),
('Despaired', 'You vote for yourself at Elimination. Permanent until cured.'),
('Madness', 'When inflicted with madness you must make efforts to present yourself as the role you’ve been made mad about. Anything deviating from that will count as breaking madness, and will result in death. This status lasts two nights unless otherwise stated.');


CREATE TABLE IF NOT EXISTS rooms(
id serial PRIMARY KEY,
name VARCHAR(64) UNIQUE NOT NULL,
description TEXT NOT NULL
);

INSERT INTO rooms (name, description) VALUES
('Grand Hall', 'Once per game, swap rooms with another player, they cannot move for a cycle.'),
('Tea Room', 'Once per game, choose a room. You learn everyone currently occupying it.'),
('Library', 'The first time a player uses an ability or AA on you, you learn what role they are (but not who they are).'),
('Conservatory', 'At game start, pick an alignment. You learn how many of that alignment are in play.'),
('Dining Hall', 'The first time you inflict a negative status on someone, you learn if it was successful, and why not, if it wasn''t.'),
('Wine Cellar', 'At game start, a random player learns you and your role.'),
('Scullery', 'Once per game, you may reroll care package or Power Drop.'),
('Butler''s Pantry', 'At game start, pick a status immunity to add to your inventory.'),
('Gallery', 'At game start, choose a player. Learn their alignment.'),
('Gift Wrapping', 'Once per game, you can vote twice.'),
('Servant''s Quarters', 'At game start, choose a role. Learn if they are in play.'),
('Panic Room', 'Survive elimination from a tied vote once.');

