-- CREATE TABLE IF NOT EXISTS players(
-- id serial PRIMARY KEY,
-- name VARCHAR(64) UNIQUE NOT NULL,
-- game_id INT REFERENCES games (id),
-- role_id INT REFERENCES roles (id),
-- alive bool NOT NULL,
-- alignment_override VARCHAR(64),
-- notes TEXT NOT NULL,
-- room_id INT REFERENCES rooms (id)
-- );
--
--
-- CREATE TABLE IF NOT EXISTS player_inventories(
-- player_id serial UNIQUE NOT NULL ,
-- ability_name VARCHAR(64) UNIQUE NOT NULL,
-- ability_quantity int,
-- PRIMARY KEY(player_id, ability_name)
-- );
--
--
-- CREATE TABLE IF NOT EXISTS abilities(
-- id serial PRIMARY KEY,
-- ability_details_id int REFERENCES ability_details (id),
-- player_inventory_id int REFERENCES player_inventories (player_id)
-- );
--

-- name: GetPlayer :one
select *
from players
where id = $1
limit 1
;

-- name: CreatePlayer :one
INSERT INTO players (
 name, game_id, role_id, alive, alignment_override, notes, room_id
  )
VALUES ( $1, $2, $3, $4, $5, $6, $7 ) RETURNING *;


-- name: ListPlayers :many
select *
from players
;

-- name: GetPlayerByID :one
select *
from players
where id = $1
;

-- name: GetAllPlayers :many
select *
from players
;

-- name: GetPlayerByName :one
select *
from players
where name = $1
;

-- name: UpdatePlayer :one
UPDATE players
  SET name = $2,
  game_id = $3,
  role_id = $4,
  alive = $5,
  alignment_override = $6,
  notes = $7,
  room_id = $8
WHERE id = $1
RETURNING *;

--Name: UpdatePlayerNotes :one
UPDATE players
  SET notes = $2
WHERE id = $1
RETURNING *;

--Name: UpdatePlayerRoom :one
UPDATE players
  SET room_id = $2
WHERE id = $1
RETURNING *;

-- name: DeletePlayer :exec
delete from players
where id = $1;
