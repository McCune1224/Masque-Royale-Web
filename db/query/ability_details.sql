-- name: GetAbilityDetail :one
select *
from ability_details
where id = $1
;

-- name: CreateAbilityDetail :one
INSERT INTO ability_details (
  name, description, role_id, category_ids, any_ability
) VALUES (
  $1, $2, $3, $4, $5
)
RETURNING *;

