-- name: GetAbilityDetail :one
select *
from ability_details
where id = $1
;

-- name: CreateAbilityDetail :one
INSERT INTO ability_details (
  name, description, role_id, category_ids, any_ability, rarity
) VALUES (
  $1, $2, $3, $4, $5, $6
)
RETURNING *;

-- name: GetAllAbilityDetails :many
SELECT * FROM ability_details;

-- name: GetAllAbilityDetailsByRoleID :many
SELECT * FROM ability_details
WHERE role_id = $1;

-- name: GetAllAbilityDetailsByCategoryID :many
SELECT * FROM ability_details
WHERE category_ids = $1;

-- name: GetAllAbilityDetailsByAnyAbility :many
SELECT * FROM ability_details
WHERE any_ability = $1;

-- name: UpdateAbilityDetail :one
UPDATE ability_details
  SET name = $2,
  description = $3,
  role_id = $4,
  category_ids = $5,
  any_ability = $6
WHERE id = $1
RETURNING *;

-- name: DeleteAbilityDetail :exec
DELETE FROM ability_details
WHERE id = $1;
