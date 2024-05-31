// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: roles.sql

package models

import (
	"context"
)

const createRole = `-- name: CreateRole :one
INSERT INTO roles (
 name, alignment, ability_ids, passive_ids
  )
VALUES ( $1, $2, $3, $4) RETURNING id, name, alignment, ability_ids, passive_ids
`

type CreateRoleParams struct {
	Name       string    `json:"name"`
	Alignment  Alignment `json:"alignment"`
	AbilityIds []int32   `json:"ability_ids"`
	PassiveIds []int32   `json:"passive_ids"`
}

func (q *Queries) CreateRole(ctx context.Context, arg CreateRoleParams) (Role, error) {
	row := q.db.QueryRow(ctx, createRole,
		arg.Name,
		arg.Alignment,
		arg.AbilityIds,
		arg.PassiveIds,
	)
	var i Role
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Alignment,
		&i.AbilityIds,
		&i.PassiveIds,
	)
	return i, err
}

const getRole = `-- name: GetRole :one
select id, name, alignment, ability_ids, passive_ids
from roles
where id = $1
limit 1
`

func (q *Queries) GetRole(ctx context.Context, id int32) (Role, error) {
	row := q.db.QueryRow(ctx, getRole, id)
	var i Role
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Alignment,
		&i.AbilityIds,
		&i.PassiveIds,
	)
	return i, err
}

const listRoles = `-- name: ListRoles :many
select id, name, alignment, ability_ids, passive_ids
from roles
`

func (q *Queries) ListRoles(ctx context.Context) ([]Role, error) {
	rows, err := q.db.Query(ctx, listRoles)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Role
	for rows.Next() {
		var i Role
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Alignment,
			&i.AbilityIds,
			&i.PassiveIds,
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
