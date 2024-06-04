-- name: GetAbilityDetail :one
select *
from ability_details
where id = $1
;

-- name: CreateAbilityDetail :one
INSERT INTO ability_details (
  name, description, default_charges, category_ids, any_ability, priority, rarity
) VALUES (
  $1, $2, $3, $4, $5, $6, $7
)
RETURNING *;

-- name: GetAllAbilityDetails :many
select *
from ability_details
;
--
-- name: GetAllAbilityDetailsByAnyAbility :many
select *
from ability_details
where any_ability = $1
;

-- name: UpdateAbilityDetail :one
UPDATE ability_details
  SET name = $2,
  description = $3,
  default_charges = $4,
  category_ids = $5,
  priority = $6,
  any_ability = $7
WHERE id = $1
RETURNING *;

-- name: DeleteAbilityDetail :exec
delete from ability_details
where id = $1
;

