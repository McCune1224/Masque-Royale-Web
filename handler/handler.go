package handler

import (
	"github.com/jmoiron/sqlx"
	"github.com/mccune1224/betrayal-widget/data"
)

type Handler struct {
	models data.Models
}

func NewHandler(db *sqlx.DB) *Handler {
	return &Handler{
		models: *data.NewModels(db),
	}
}
