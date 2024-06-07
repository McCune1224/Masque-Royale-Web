-- name: GetAction :one
select *
from actions
where id = $1
;

-- name: CreateAction :one
INSERT INTO actions (
  game_id, player_id, pending_approval, resolved, target, context, ability_name, role_id
) VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8 
)
RETURNING *;

-- name: ListActions :many
select *
from actions
;

-- name: ListActionsByGame :many
select *
from actions
where game_id = $1
;

-- name: GetActionByID :one
select *
from actions
where id = $1
;
-- name: UpdateAction :one
UPDATE actions
  SET game_id = $2,
  player_id = $3,
  pending_approval = $4,
  resolved = $5,
  target = $6,
  context = $7,
  ability_name = $8,
  role_id = $9
WHERE id = $1
RETURNING *;

-- name: DeleteAction :exec
delete from actions
where id = $1
;

-- name: ListActionsByRoundForGame :many
SELECT a.*
FROM actions a
JOIN games g on $1 = a.game_id
WHERE g.round = $2
;
;

-- name: ListActionsByPlayer :many
SELECT a.*
FROM actions a
JOIN players p on p.id = a.player_id
WHERE p.id = $1
;
