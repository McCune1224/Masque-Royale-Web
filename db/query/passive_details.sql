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

