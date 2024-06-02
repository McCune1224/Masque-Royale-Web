-- name: GetRoom :one
select *
from rooms
where id = $1
limit 1
;

-- name: ListRooms :many
select *
from rooms
;

-- name: GetRoomByID :one
select *
from rooms
where id = $1
;

-- name: GetRoomBynName :one
select *
from rooms
where name = $1
;
