-- name: GetRole :one
select *
from roles
where id = $1
limit 1
;

-- name: CreateRole :one
INSERT INTO roles (
 name, alignment, ability_ids, passive_ids
  )
VALUES ( $1, $2, $3, $4) RETURNING *;


-- name: ListRoles :many
select *
from roles
;

