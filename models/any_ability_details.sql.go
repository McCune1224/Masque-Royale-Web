// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: any_ability_details.sql

package models

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createAnyAbilityDetail = `-- name: CreateAnyAbilityDetail :one
INSERT INTO any_ability_details (
  name, description, category_ids, shorthand, rarity, priority
) VALUES (
  $1, $2, $3, $4, $5 , $6
)
  RETURNING id, name, shorthand, description, category_ids, rarity, priority
`

type CreateAnyAbilityDetailParams struct {
	Name        string      `json:"name"`
	Description string      `json:"description"`
	CategoryIds []int32     `json:"category_ids"`
	Shorthand   string      `json:"shorthand"`
	Rarity      Rarity      `json:"rarity"`
	Priority    pgtype.Int4 `json:"priority"`
}

// CREATE TABLE IF NOT EXISTS any_ability_details (
// id serial PRIMARY KEY,
// name VARCHAR(64) UNIQUE NOT NULL,
// description TEXT NOT NULL,
// category_ids INT[] DEFAULT '{}',
// rarity rarity NOT NULL,
// priority int
// );
func (q *Queries) CreateAnyAbilityDetail(ctx context.Context, arg CreateAnyAbilityDetailParams) (AnyAbilityDetail, error) {
	row := q.db.QueryRow(ctx, createAnyAbilityDetail,
		arg.Name,
		arg.Description,
		arg.CategoryIds,
		arg.Shorthand,
		arg.Rarity,
		arg.Priority,
	)
	var i AnyAbilityDetail
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Shorthand,
		&i.Description,
		&i.CategoryIds,
		&i.Rarity,
		&i.Priority,
	)
	return i, err
}

const deleteAnyAbilityDetail = `-- name: DeleteAnyAbilityDetail :exec
delete from any_ability_details
where id = $1
`

func (q *Queries) DeleteAnyAbilityDetail(ctx context.Context, id int32) error {
	_, err := q.db.Exec(ctx, deleteAnyAbilityDetail, id)
	return err
}

const getAnyAbilityDetail = `-- name: GetAnyAbilityDetail :one
select id, name, shorthand, description, category_ids, rarity, priority
from any_ability_details
where id = $1
`

func (q *Queries) GetAnyAbilityDetail(ctx context.Context, id int32) (AnyAbilityDetail, error) {
	row := q.db.QueryRow(ctx, getAnyAbilityDetail, id)
	var i AnyAbilityDetail
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Shorthand,
		&i.Description,
		&i.CategoryIds,
		&i.Rarity,
		&i.Priority,
	)
	return i, err
}

const getAnyAbilityDetailsByID = `-- name: GetAnyAbilityDetailsByID :many
select id, name, shorthand, description, category_ids, rarity, priority
from any_ability_details
where id = $1
`

func (q *Queries) GetAnyAbilityDetailsByID(ctx context.Context, id int32) ([]AnyAbilityDetail, error) {
	rows, err := q.db.Query(ctx, getAnyAbilityDetailsByID, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []AnyAbilityDetail
	for rows.Next() {
		var i AnyAbilityDetail
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Shorthand,
			&i.Description,
			&i.CategoryIds,
			&i.Rarity,
			&i.Priority,
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

const listAnyAbilityDetails = `-- name: ListAnyAbilityDetails :many
select id, name, shorthand, description, category_ids, rarity, priority
from any_ability_details
`

func (q *Queries) ListAnyAbilityDetails(ctx context.Context) ([]AnyAbilityDetail, error) {
	rows, err := q.db.Query(ctx, listAnyAbilityDetails)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []AnyAbilityDetail
	for rows.Next() {
		var i AnyAbilityDetail
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Shorthand,
			&i.Description,
			&i.CategoryIds,
			&i.Rarity,
			&i.Priority,
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

const updateAnyAbilityDetail = `-- name: UpdateAnyAbilityDetail :one
UPDATE any_ability_details
  SET name = $2,
  description = $3,
  category_ids = $4,
  rarity = $5,
  shorthand = $6,
  priority = $7
WHERE id = $1
RETURNING id, name, shorthand, description, category_ids, rarity, priority
`

type UpdateAnyAbilityDetailParams struct {
	ID          int32       `json:"id"`
	Name        string      `json:"name"`
	Description string      `json:"description"`
	CategoryIds []int32     `json:"category_ids"`
	Rarity      Rarity      `json:"rarity"`
	Shorthand   string      `json:"shorthand"`
	Priority    pgtype.Int4 `json:"priority"`
}

func (q *Queries) UpdateAnyAbilityDetail(ctx context.Context, arg UpdateAnyAbilityDetailParams) (AnyAbilityDetail, error) {
	row := q.db.QueryRow(ctx, updateAnyAbilityDetail,
		arg.ID,
		arg.Name,
		arg.Description,
		arg.CategoryIds,
		arg.Rarity,
		arg.Shorthand,
		arg.Priority,
	)
	var i AnyAbilityDetail
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Shorthand,
		&i.Description,
		&i.CategoryIds,
		&i.Rarity,
		&i.Priority,
	)
	return i, err
}
