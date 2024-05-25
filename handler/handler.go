package handler

import (
	"github.com/jackc/pgx/v5/pgxpool"
)

type Handler struct {
	Db *pgxpool.Pool
}

func NewHandler(db *pgxpool.Pool) *Handler {
	return &Handler{Db: db}
}
