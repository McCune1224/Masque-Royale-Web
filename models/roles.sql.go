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
 name, alignment 
  )
VALUES ( $1, $2 ) RETURNING id, name, alignment
`

type CreateRoleParams struct {
	Name      string    `json:"name"`
	Alignment Alignment `json:"alignment"`
}

func (q *Queries) CreateRole(ctx context.Context, arg CreateRoleParams) (Role, error) {
	row := q.db.QueryRow(ctx, createRole, arg.Name, arg.Alignment)
	var i Role
	err := row.Scan(&i.ID, &i.Name, &i.Alignment)
	return i, err
}

const getRole = `-- name: GetRole :one
select id, name, alignment
from roles
where id = $1
limit 1
`

func (q *Queries) GetRole(ctx context.Context, id int32) (Role, error) {
	row := q.db.QueryRow(ctx, getRole, id)
	var i Role
	err := row.Scan(&i.ID, &i.Name, &i.Alignment)
	return i, err
}

const getRolePassivesAggregate = `-- name: GetRolePassivesAggregate :one
select
    array_agg(passive_details.name) as passive_names,
    array_agg(passive_details.description) as passive_descriptions
from roles
join role_passives_join on role_passives_join.role_id = roles.id
join passive_details on passive_details.id = role_passives_join.passive_id
where roles.id = $1
group by roles.id
`

type GetRolePassivesAggregateRow struct {
	PassiveNames        interface{} `json:"passive_names"`
	PassiveDescriptions interface{} `json:"passive_descriptions"`
}

func (q *Queries) GetRolePassivesAggregate(ctx context.Context, id int32) (GetRolePassivesAggregateRow, error) {
	row := q.db.QueryRow(ctx, getRolePassivesAggregate, id)
	var i GetRolePassivesAggregateRow
	err := row.Scan(&i.PassiveNames, &i.PassiveDescriptions)
	return i, err
}

const listRoles = `-- name: ListRoles :many
select id, name, alignment
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
		if err := rows.Scan(&i.ID, &i.Name, &i.Alignment); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const nukeRoles = `-- name: NukeRoles :exec
TRUNCATE roles, role_abilities_join, role_passives_join, ability_details, passive_details, ability_details_categories_join RESTART IDENTITY CASCADE
`

func (q *Queries) NukeRoles(ctx context.Context) error {
	_, err := q.db.Exec(ctx, nukeRoles)
	return err
}
