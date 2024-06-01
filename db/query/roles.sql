-- name: GetRole :one
select *
from roles
where id = $1
limit 1
;

-- name: CreateRole :one
INSERT INTO roles (
 name, alignment 
  )
VALUES ( $1, $2 ) RETURNING *;


-- name: ListRoles :many
select *
from roles
;

-- name: GetRoleAbilityAndPassiveJoin :one
SELECT sqlc.embed(role_abilites_join), sqlc.embed(abilities), sqlc.embed(passive_details)
FROM role_abilites_join
JOIN abilities ON role_abilites_join.ability_id = abilities.id
JOIN passive_details ON role_abilites_join.passive_id = passive_details.id
;

-- name: NukeRoles :exec
TRUNCATE roles, role_abilites_join, role_passives_join, ability_details, passive_details RESTART IDENTITY CASCADE;

