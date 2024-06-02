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

-- name: ListPlayersByGame :many
select *
from players
where game_id = $1
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

-- name: UpdatePlayerNotes :one
UPDATE players
  SET notes = $2
WHERE id = $1
RETURNING *;

-- name: UpdatePlayerAlive :one
UPDATE players
  SET alive = $2
WHERE id = $1
RETURNING *;

-- name: UpdatePlayerRole :one
UPDATE players
  SET role_id = $2
WHERE id = $1
RETURNING *;

-- name: UpdatePlayerAlignmentOverride :one
UPDATE players
  SET alignment_override = $2
WHERE id = $1
RETURNING *;

-- name: UpdatePlayerRoom :one
UPDATE players
  SET room_id = $2
WHERE id = $1
RETURNING *;


-- name: DeletePlayer :exec
delete from players
where id = $1;
