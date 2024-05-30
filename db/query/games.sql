-- name: GetGame :one
select *
from games
where id = $1
limit 1
;

-- name: GetGameByName :one
select *
from games
where name = $1
limit 1
;

-- name: ListGames :many
select *
from games
order by name
;

-- name: GetRandomGame :one
select *
from games
order by random()
limit 1
;

-- name: UpdateGame :one
UPDATE games
  SET name = $2,
  phase = $3,
  round = $4,
  player_ids = $5
WHERE id = $1
RETURNING *;

-- name: CreateGame :one
INSERT INTO games (
  name, phase, round, player_ids
) VALUES (
  $1, $2, $3, $4
)
RETURNING *;

-- name: DeleteGame :exec
delete from games
where id = $1
;

