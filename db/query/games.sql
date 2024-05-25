-- name: GetGame :one
SELECT * FROM games
WHERE id = $1 LIMIT 1;

-- name: GetGameByName :one
SELECT * FROM games
WHERE name = $1 LIMIT 1;

-- name: ListGames :many
SELECT * FROM games
ORDER BY name;

-- name: GetRandomGame :one
SELECT * FROM games
ORDER BY random() LIMIT 1;

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
DELETE FROM games
WHERE id = $1;
