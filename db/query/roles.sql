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
select
    array_agg(passive_details.name) as passive_names,
    array_agg(passive_details.description) as passive_descriptions
from roles
join role_passives_join on role_passives_join.role_id = roles.id
join passive_details on passive_details.id = role_passives_join.passive_id
where roles.id = $1
group by roles.id
;
;


-- name: NukeRoles :exec
TRUNCATE roles, role_abilities_join, role_passives_join, ability_details, any_ability_details, passive_details RESTART IDENTITY CASCADE;

