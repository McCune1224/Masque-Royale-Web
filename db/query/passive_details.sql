-- name: GetPassiveDetail :one
select *
from passive_details
where id = $1
;

-- name: CreatePassiveDetail :one
INSERT INTO passive_details (
  name, description
) VALUES (
  $1, $2
)
RETURNING *;

-- name: GetAllPassiveDetails :many
SELECT * FROM passive_details;

-- name: GetAllPassiveDetailsByID :many
SELECT * FROM passive_details
WHERE id = $1;

-- name: UpdatePassiveDetail :one
UPDATE passive_details
  SET name = $2,
  description = $3
WHERE id = $1
RETURNING *;

-- name: DeletePassiveDetail :exec
DELETE FROM passive_details
WHERE id = $1;
