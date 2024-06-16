-- name: CreateAbilityDetailsCategoriesJoin :one
INSERT INTO ability_details_categories_join (
  ability_details_id, categories_id
) VALUES (
  $1, $2
) RETURNING *;

-- name: GetCategoriesForAbilityByID :many
select categories.*
from categories
join
    ability_details_categories_join
    on categories.id = ability_details_categories_join.categories_id
where ability_details_categories_join.ability_details_id = 1
;

