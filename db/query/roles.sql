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

-- name: GetRolePassivesAggregate :one
SELECT ARRAY_AGG(passive_details.name) AS passive_names, ARRAY_AGG(passive_details.description) AS passive_descriptions FROM roles
JOIN role_passives_join ON role_passives_join.role_id = roles.id
JOIN passive_details ON passive_details.id = role_passives_join.passive_id
WHERE roles.id = $1 GROUP BY roles.id;
;



-- name: NukeRoles :exec
TRUNCATE roles, role_abilities_join, role_passives_join, ability_details, passive_details RESTART IDENTITY CASCADE;

