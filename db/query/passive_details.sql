-- name: GetPassiveDetails :one
select *
from passive_details
where id = $1
;

-- name: GetPassiveDetailsByName :one
select *
from passive_details
where name = $1
;

-- name: CreatePassiveDetail :one
INSERT INTO passive_details (
  name, description
) VALUES (
  $1, $2
)
RETURNING *;

-- name: GetAllPassiveDetails :many
select *
from passive_details
;

-- name: GetAllPassiveDetailsByID :many
select *
from passive_details
where id = $1
;

-- name: UpdatePassiveDetail :one
UPDATE passive_details
  SET name = $2,
  description = $3
WHERE id = $1
RETURNING *;

-- name: DeletePassiveDetail :exec
delete from passive_details
where id = $1
;

