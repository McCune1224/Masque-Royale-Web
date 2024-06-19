--
-- name: GetPlayer :one
select *
from players
where id = $1
limit 1
;

-- name: CreatePlayer :one
insert into players
  ( name, game_id, role_id, alive, alignment, room_id)
values ( $1, $2, $3, $4, $5, $6 )
returning *
;


-- name: ListPlayers :many
select *
from players
;

-- name: ListPlayersByGame :many
select *
from players
where game_id = $1
;

-- name: GetPlayerByID :one
select *
from players
where id = $1
;

-- name: GetAllPlayers :many
select *
from players
;

-- name: GetPlayerByName :one
select *
from players
where name = $1
;

-- name: UpdatePlayer :one
update players
set name = $2,
game_id = $3,
role_id = $4,
alive = $5,
alignment = $6,
room_id = $7
where id = $1
returning *;

-- name: UpdatePlayerAlive :one
update players
set alive = $2
where id = $1
returning *;


-- name: UpdatePlayerRole :one
update players
set role_id = $2
where id = $1
returning *;

-- name: UpdatePlayerAlignment :one
update players
set alignment = $2
where id = $1
returning *;

-- name: UpdatePlayerRoom :one
update players
set room_id = $2
where id = $1
returning *;

-- name: GetPlayerNote :one
select *
from player_notes
where player_id = $1
;

-- name: UpsertPlayerNote :one
insert into player_notes
  ( player_id, note )
values ( $1, $2 )
on conflict (player_id) do update
  set note = $2
returning *
;

-- name: DeletePlayerNote :exec
delete from player_notes
where player_id = $1
;

-- name: ListPlayerAbilites :one
select *
from player_abilities
where player_id = $1
;

-- name: CreatePlayerAbility :one
insert into player_abilities
  ( player_id, ability_details_id, charges )
values ( $1, $2, $3 )
returning *
;

-- name: UpdatePlayerAbility :one
update player_abilities
set charges = $3
where player_id = $1
and ability_details_id = $2
returning *
;

-- name: DeletePlayerAbility :exec
delete from player_abilities
where player_id = $1
and ability_details_id = $2
;

-- name: CreatePlayerStatus :one
insert into player_statuses
  ( player_id, status_id, stack, round_given )
values ( $1, $2, $3, $4 )
returning *
;



-- name: DeletePlayer :exec
delete from players
where id = $1;



-- name: ListPlayerAbilitiesJoin :many
select  
  ability_details_id, 
  charges, 
  name, 
  description, 
  rarity, 
  any_ability
from player_abilities pa 
inner join ability_details ad on ad.id = pa.ability_details_id
where player_id = $1
;

-- name: GetPlayerAbility :one
select  
  ability_details_id, 
  charges, 
  name, 
  description, 
  rarity, 
  any_ability
from player_abilities pa 
inner join ability_details ad on ad.id = pa.ability_details_id
where player_id = $1 and ability_details_id = $2
;

