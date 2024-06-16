-- name: GetAbilityDetail :one
select *
from ability_details
where id = $1
;

-- name: GetAbilityDetailsByName :one
select *
from ability_details
where name = $1
;

-- name: CreateAbilityDetail :one
INSERT INTO ability_details (
  name, description, default_charges, any_ability, rarity
) VALUES (
  $1, $2, $3, $4, $5 
)
RETURNING *;

-- name: ListAbilityDetails :many
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
  any_ability = $5
WHERE id = $1
RETURNING *;

-- name: DeleteAbilityDetail :exec
delete from ability_details
where id = $1
;

-- name: GetAnyAbilityDetailsMarkedAnyAbility :many
select *
from ability_details
where any_ability = true
;

