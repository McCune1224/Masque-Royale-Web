-- name: GetRoleAbilityJoin :one
select sqlc.embed(role_abilities_join), sqlc.embed(abilities)
from role_abilities_join
join abilities on role_abilities_join.ability_id = abilities.id
;

-- name: CreateRoleAbilityJoin :one
INSERT INTO role_abilities_join (
  role_id, ability_id
) VALUES (
  $1, $2
) RETURNING *;


-- name: GetAssociatedRoleAbilities :many
select ab.*
from role_abilities_join raj
join ability_details ab on raj.ability_id = ab.id
where raj.role_id = $1
;



-- name: GetRoleFromAbilityID :one
SELECT r.*
FROM roles r
JOIN role_abilities_Join ra ON r.id = ra.role_id
WHERE ra.ability_id = $1;
;

-- name: NukeAnyAbilities :exec
TRUNCATE  any_ability_details RESTART IDENTITY CASCADE;



