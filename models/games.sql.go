// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.23.0
// source: games.sql

package models

import (
	"context"
)

const createGame = `-- name: CreateGame :one
INSERT INTO games (
  name, phase, round, player_ids
) VALUES (
  $1, $2, $3, $4
)
RETURNING id, name, phase, round, player_ids, created_at, updated_at
`

type CreateGameParams struct {
	Name      string  `json:"name"`
	Phase     string  `json:"phase"`
	Round     int32   `json:"round"`
	PlayerIds []int32 `json:"player_ids"`
}

func (q *Queries) CreateGame(ctx context.Context, arg CreateGameParams) (Game, error) {
	row := q.db.QueryRow(ctx, createGame,
		arg.Name,
		arg.Phase,
		arg.Round,
		arg.PlayerIds,
	)
	var i Game
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Phase,
		&i.Round,
		&i.PlayerIds,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteGame = `-- name: DeleteGame :exec
DELETE FROM games
WHERE id = $1
`

func (q *Queries) DeleteGame(ctx context.Context, id int32) error {
	_, err := q.db.Exec(ctx, deleteGame, id)
	return err
}

const getGame = `-- name: GetGame :one
SELECT id, name, phase, round, player_ids, created_at, updated_at FROM games
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetGame(ctx context.Context, id int32) (Game, error) {
	row := q.db.QueryRow(ctx, getGame, id)
	var i Game
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Phase,
		&i.Round,
		&i.PlayerIds,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getGameByName = `-- name: GetGameByName :one
SELECT id, name, phase, round, player_ids, created_at, updated_at FROM games
WHERE name = $1 LIMIT 1
`

func (q *Queries) GetGameByName(ctx context.Context, name string) (Game, error) {
	row := q.db.QueryRow(ctx, getGameByName, name)
	var i Game
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Phase,
		&i.Round,
		&i.PlayerIds,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getRandomGame = `-- name: GetRandomGame :one
SELECT id, name, phase, round, player_ids, created_at, updated_at FROM games
ORDER BY random() LIMIT 1
`

func (q *Queries) GetRandomGame(ctx context.Context) (Game, error) {
	row := q.db.QueryRow(ctx, getRandomGame)
	var i Game
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Phase,
		&i.Round,
		&i.PlayerIds,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const listGames = `-- name: ListGames :many
SELECT id, name, phase, round, player_ids, created_at, updated_at FROM games
ORDER BY name
`

func (q *Queries) ListGames(ctx context.Context) ([]Game, error) {
	rows, err := q.db.Query(ctx, listGames)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Game
	for rows.Next() {
		var i Game
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Phase,
			&i.Round,
			&i.PlayerIds,
			&i.CreatedAt,
			&i.UpdatedAt,
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

const updateGame = `-- name: UpdateGame :one
UPDATE games
  SET name = $2,
  phase = $3,
  round = $4,
  player_ids = $5
WHERE id = $1
RETURNING id, name, phase, round, player_ids, created_at, updated_at
`

type UpdateGameParams struct {
	ID        int32   `json:"id"`
	Name      string  `json:"name"`
	Phase     string  `json:"phase"`
	Round     int32   `json:"round"`
	PlayerIds []int32 `json:"player_ids"`
}

func (q *Queries) UpdateGame(ctx context.Context, arg UpdateGameParams) (Game, error) {
	row := q.db.QueryRow(ctx, updateGame,
		arg.ID,
		arg.Name,
		arg.Phase,
		arg.Round,
		arg.PlayerIds,
	)
	var i Game
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Phase,
		&i.Round,
		&i.PlayerIds,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
