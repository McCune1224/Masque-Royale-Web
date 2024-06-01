
-- name: GetRoleAbilityJoin :one
SELECT sqlc.embed(role_abilites_join), sqlc.embed(abilities)
FROM role_abilites_join
JOIN abilities ON role_abilites_join.ability_id = abilities.id
;

-- name: CreateRoleAbilityJoin :one
INSERT INTO role_abilites_join (
  role_id, ability_id
) VALUES (
  $1, $2
) RETURNING *;

