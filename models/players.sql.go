// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: players.sql

package models

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createPlayer = `-- name: CreatePlayer :one
INSERT INTO players (
 name, game_id, role_id, alive, alignment_override, notes, room_id
  )
VALUES ( $1, $2, $3, $4, $5, $6, $7 ) RETURNING id, name, game_id, role_id, alive, alignment_override, notes, room_id
`

type CreatePlayerParams struct {
	Name              string      `json:"name"`
	GameID            pgtype.Int4 `json:"game_id"`
	RoleID            pgtype.Int4 `json:"role_id"`
	Alive             bool        `json:"alive"`
	AlignmentOverride pgtype.Text `json:"alignment_override"`
	Notes             string      `json:"notes"`
	RoomID            pgtype.Int4 `json:"room_id"`
}

func (q *Queries) CreatePlayer(ctx context.Context, arg CreatePlayerParams) (Player, error) {
	row := q.db.QueryRow(ctx, createPlayer,
		arg.Name,
		arg.GameID,
		arg.RoleID,
		arg.Alive,
		arg.AlignmentOverride,
		arg.Notes,
		arg.RoomID,
	)
	var i Player
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.GameID,
		&i.RoleID,
		&i.Alive,
		&i.AlignmentOverride,
		&i.Notes,
		&i.RoomID,
	)
	return i, err
}

const deletePlayer = `-- name: DeletePlayer :exec
delete from players
where id = $1
`

func (q *Queries) DeletePlayer(ctx context.Context, id int32) error {
	_, err := q.db.Exec(ctx, deletePlayer, id)
	return err
}

const getAllPlayers = `-- name: GetAllPlayers :many
select id, name, game_id, role_id, alive, alignment_override, notes, room_id
from players
`

func (q *Queries) GetAllPlayers(ctx context.Context) ([]Player, error) {
	rows, err := q.db.Query(ctx, getAllPlayers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Player
	for rows.Next() {
		var i Player
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.GameID,
			&i.RoleID,
			&i.Alive,
			&i.AlignmentOverride,
			&i.Notes,
			&i.RoomID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getPlayer = `-- name: GetPlayer :one

select id, name, game_id, role_id, alive, alignment_override, notes, room_id
from players
where id = $1
limit 1
`

// CREATE TABLE IF NOT EXISTS players(
// id serial PRIMARY KEY,
// name VARCHAR(64) UNIQUE NOT NULL,
// game_id INT REFERENCES games (id),
// role_id INT REFERENCES roles (id),
// alive bool NOT NULL,
// alignment_override VARCHAR(64),
// notes TEXT NOT NULL,
// room_id INT REFERENCES rooms (id)
// );
//
// CREATE TABLE IF NOT EXISTS player_inventories(
// player_id serial UNIQUE NOT NULL ,
// ability_name VARCHAR(64) UNIQUE NOT NULL,
// ability_quantity int,
// PRIMARY KEY(player_id, ability_name)
// );
//
// CREATE TABLE IF NOT EXISTS abilities(
// id serial PRIMARY KEY,
// ability_details_id int REFERENCES ability_details (id),
// player_inventory_id int REFERENCES player_inventories (player_id)
// );
func (q *Queries) GetPlayer(ctx context.Context, id int32) (Player, error) {
	row := q.db.QueryRow(ctx, getPlayer, id)
	var i Player
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.GameID,
		&i.RoleID,
		&i.Alive,
		&i.AlignmentOverride,
		&i.Notes,
		&i.RoomID,
	)
	return i, err
}

const getPlayerByID = `-- name: GetPlayerByID :one
select id, name, game_id, role_id, alive, alignment_override, notes, room_id
from players
where id = $1
`

func (q *Queries) GetPlayerByID(ctx context.Context, id int32) (Player, error) {
	row := q.db.QueryRow(ctx, getPlayerByID, id)
	var i Player
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.GameID,
		&i.RoleID,
		&i.Alive,
		&i.AlignmentOverride,
		&i.Notes,
		&i.RoomID,
	)
	return i, err
}

const getPlayerByName = `-- name: GetPlayerByName :one
select id, name, game_id, role_id, alive, alignment_override, notes, room_id
from players
where name = $1
`

func (q *Queries) GetPlayerByName(ctx context.Context, name string) (Player, error) {
	row := q.db.QueryRow(ctx, getPlayerByName, name)
	var i Player
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.GameID,
		&i.RoleID,
		&i.Alive,
		&i.AlignmentOverride,
		&i.Notes,
		&i.RoomID,
	)
	return i, err
}

const listPlayers = `-- name: ListPlayers :many
select id, name, game_id, role_id, alive, alignment_override, notes, room_id
from players
`

func (q *Queries) ListPlayers(ctx context.Context) ([]Player, error) {
	rows, err := q.db.Query(ctx, listPlayers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Player
	for rows.Next() {
		var i Player
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.GameID,
			&i.RoleID,
			&i.Alive,
			&i.AlignmentOverride,
			&i.Notes,
			&i.RoomID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updatePlayer = `-- name: UpdatePlayer :one
UPDATE players
  SET name = $2,
  game_id = $3,
  role_id = $4,
  alive = $5,
  alignment_override = $6,
  notes = $7,
  room_id = $8
WHERE id = $1
RETURNING id, name, game_id, role_id, alive, alignment_override, notes, room_id
`

type UpdatePlayerParams struct {
	ID                int32       `json:"id"`
	Name              string      `json:"name"`
	GameID            pgtype.Int4 `json:"game_id"`
	RoleID            pgtype.Int4 `json:"role_id"`
	Alive             bool        `json:"alive"`
	AlignmentOverride pgtype.Text `json:"alignment_override"`
	Notes             string      `json:"notes"`
	RoomID            pgtype.Int4 `json:"room_id"`
}

func (q *Queries) UpdatePlayer(ctx context.Context, arg UpdatePlayerParams) (Player, error) {
	row := q.db.QueryRow(ctx, updatePlayer,
		arg.ID,
		arg.Name,
		arg.GameID,
		arg.RoleID,
		arg.Alive,
		arg.AlignmentOverride,
		arg.Notes,
		arg.RoomID,
	)
	var i Player
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.GameID,
		&i.RoleID,
		&i.Alive,
		&i.AlignmentOverride,
		&i.Notes,
		&i.RoomID,
	)
	return i, err
}
