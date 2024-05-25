// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.23.0

package models

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type Game struct {
	ID        int32            `json:"id"`
	Name      string           `json:"name"`
	Phase     string           `json:"phase"`
	Round     int32            `json:"round"`
	PlayerIds []int32          `json:"player_ids"`
	CreatedAt pgtype.Timestamp `json:"created_at"`
	UpdatedAt pgtype.Timestamp `json:"updated_at"`
}
