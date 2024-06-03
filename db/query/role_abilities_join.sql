
-- name: GetRoleAbilityJoin :one
SELECT sqlc.embed(role_abilities_join), sqlc.embed(abilities)
FROM role_abilities_join
JOIN abilities ON role_abilities_join.ability_id = abilities.id
;

-- name: CreateRoleAbilityJoin :one
INSERT INTO role_abilities_join (
  role_id, ability_id
) VALUES (
  $1, $2
) RETURNING *;


-- name: GetAssociatedRoleAbilities :many
SELECT ab.*
FROM role_abilities_join raj
JOIN ability_details ab ON raj.ability_id = ab.id
WHERE raj.role_id = $1;
