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

const getRoleAbilityAndPassiveJoin = `-- name: GetRoleAbilityAndPassiveJoin :one
SELECT role_abilites_join.role_id, role_abilites_join.ability_id, abilities.id, abilities.ability_details_id, abilities.player_inventory_id, passive_details.id, passive_details.name, passive_details.description
FROM role_abilites_join
JOIN abilities ON role_abilites_join.ability_id = abilities.id
JOIN passive_details ON role_abilites_join.passive_id = passive_details.id
`

type GetRoleAbilityAndPassiveJoinRow struct {
	RoleAbilitesJoin RoleAbilitesJoin `json:"role_abilites_join"`
	Ability          Ability          `json:"ability"`
	PassiveDetail    PassiveDetail    `json:"passive_detail"`
}

func (q *Queries) GetRoleAbilityAndPassiveJoin(ctx context.Context) (GetRoleAbilityAndPassiveJoinRow, error) {
	row := q.db.QueryRow(ctx, getRoleAbilityAndPassiveJoin)
	var i GetRoleAbilityAndPassiveJoinRow
	err := row.Scan(
		&i.RoleAbilitesJoin.RoleID,
		&i.RoleAbilitesJoin.AbilityID,
		&i.Ability.ID,
		&i.Ability.AbilityDetailsID,
		&i.Ability.PlayerInventoryID,
		&i.PassiveDetail.ID,
		&i.PassiveDetail.Name,
		&i.PassiveDetail.Description,
	)
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
TRUNCATE roles, role_abilites_join, role_passives_join, ability_details, passive_details RESTART IDENTITY CASCADE
`

func (q *Queries) NukeRoles(ctx context.Context) error {
	_, err := q.db.Exec(ctx, nukeRoles)
	return err
}
