-- name: GetStatusDetail :one
select *
from status_details
where id = $1
;

-- name: CreateStatusDetail :one
INSERT INTO status_details (
  name, description
) VALUES (
  $1, $2
)
RETURNING *;

-- name: ListStatusDetails :many
select *
from status_details
;

-- name: GetStatusDetailByID :one
select *
from status_details
where id = $1
;

-- name: GetStatusDetailByName :one
select *
from status_details
where name = $1
;
