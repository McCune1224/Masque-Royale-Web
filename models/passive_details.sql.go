// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: passive_details.sql

package models

import (
	"context"
)

const createPassiveDetail = `-- name: CreatePassiveDetail :one
INSERT INTO passive_details (
  name, description
) VALUES (
  $1, $2
)
RETURNING id, name, description
`

type CreatePassiveDetailParams struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (q *Queries) CreatePassiveDetail(ctx context.Context, arg CreatePassiveDetailParams) (PassiveDetail, error) {
	row := q.db.QueryRow(ctx, createPassiveDetail, arg.Name, arg.Description)
	var i PassiveDetail
	err := row.Scan(&i.ID, &i.Name, &i.Description)
	return i, err
}

const deletePassiveDetail = `-- name: DeletePassiveDetail :exec
delete from passive_details
where id = $1
`

func (q *Queries) DeletePassiveDetail(ctx context.Context, id int32) error {
	_, err := q.db.Exec(ctx, deletePassiveDetail, id)
	return err
}

const getAllPassiveDetails = `-- name: GetAllPassiveDetails :many
select id, name, description
from passive_details
`

func (q *Queries) GetAllPassiveDetails(ctx context.Context) ([]PassiveDetail, error) {
	rows, err := q.db.Query(ctx, getAllPassiveDetails)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []PassiveDetail
	for rows.Next() {
		var i PassiveDetail
		if err := rows.Scan(&i.ID, &i.Name, &i.Description); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getAllPassiveDetailsByID = `-- name: GetAllPassiveDetailsByID :many
select id, name, description
from passive_details
where id = $1
`

func (q *Queries) GetAllPassiveDetailsByID(ctx context.Context, id int32) ([]PassiveDetail, error) {
	rows, err := q.db.Query(ctx, getAllPassiveDetailsByID, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []PassiveDetail
	for rows.Next() {
		var i PassiveDetail
		if err := rows.Scan(&i.ID, &i.Name, &i.Description); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getPassiveDetails = `-- name: GetPassiveDetails :one
select id, name, description
from passive_details
where id = $1
`

func (q *Queries) GetPassiveDetails(ctx context.Context, id int32) (PassiveDetail, error) {
	row := q.db.QueryRow(ctx, getPassiveDetails, id)
	var i PassiveDetail
	err := row.Scan(&i.ID, &i.Name, &i.Description)
	return i, err
}

const getPassiveDetailsByName = `-- name: GetPassiveDetailsByName :one
select id, name, description
from passive_details
where name = $1
`

func (q *Queries) GetPassiveDetailsByName(ctx context.Context, name string) (PassiveDetail, error) {
	row := q.db.QueryRow(ctx, getPassiveDetailsByName, name)
	var i PassiveDetail
	err := row.Scan(&i.ID, &i.Name, &i.Description)
	return i, err
}

const updatePassiveDetail = `-- name: UpdatePassiveDetail :one
UPDATE passive_details
  SET name = $2,
  description = $3
WHERE id = $1
RETURNING id, name, description
`

type UpdatePassiveDetailParams struct {
	ID          int32  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (q *Queries) UpdatePassiveDetail(ctx context.Context, arg UpdatePassiveDetailParams) (PassiveDetail, error) {
	row := q.db.QueryRow(ctx, updatePassiveDetail, arg.ID, arg.Name, arg.Description)
	var i PassiveDetail
	err := row.Scan(&i.ID, &i.Name, &i.Description)
	return i, err
}
