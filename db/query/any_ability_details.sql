-- name: GetAnyAbilityDetail :one
select *
from any_ability_details
where id = $1
;

-- name: GetAnyAbilityByName :one
select *
from any_ability_details
where name = $1
;

-- name: CreateAnyAbilityDetail :one
INSERT INTO any_ability_details (
  name, description, category_ids, shorthand, rarity, priority
) VALUES (
  $1, $2, $3, $4, $5 , $6
)
  RETURNING *;

-- name: ListAnyAbilityDetails :many
select *
from any_ability_details
;

-- name: UpdateAnyAbilityDetail :one
UPDATE any_ability_details
  SET name = $2,
  description = $3,
  category_ids = $4,
  rarity = $5,
  shorthand = $6,
  priority = $7
WHERE id = $1
RETURNING *;

-- name: DeleteAnyAbilityDetail :exec
delete from any_ability_details
where id = $1
;

