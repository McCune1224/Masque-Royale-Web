-- name: CreateRolePassiveJoin :one
INSERT INTO role_passives_join (
  role_id, passive_id
) VALUES (
  $1, $2
) RETURNING *;


-- name: GetRolePassiveJoin :one
SELECT sqlc.embed(role_passives_join), sqlc.embed(passive_details)
FROM role_passives_join
JOIN passive_details ON role_passives_join.passive_id = passive_details.id
;

