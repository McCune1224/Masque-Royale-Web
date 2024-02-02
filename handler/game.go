package handler

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

func (h *Handler) CreateGame(c echo.Context) error {
	gameID := c.QueryParam("game_id")
	if gameID == "" {
		return c.HTML(400, "game_id is required")
	}

	_, err := h.models.Games.InsertGame(gameID)
	if err != nil {
		return c.HTML(500, "<div>Error inserting game</div>")
	}

	c.SetCookie(&http.Cookie{
		Name:     "game_id",
		Value:    gameID,
		Path:     "/",
		SameSite: http.SameSiteNoneMode,
		Secure:   true,
	})

	return c.Redirect(302, "/games/dashboard")
}

func (h *Handler) JoinGame(c echo.Context) error {
	gameID := c.QueryParam("game_id")
	if gameID == "" {
		log.Println("hit within gameID check")
		return c.HTML(400, "game_id is required")
	}

	log.Println("hit")
	_, err := h.models.Games.GetByGameID(gameID)
	if err != nil {
		log.Println(err)
		return c.HTML(500, "<div>Error getting game</div>")
	}
	c.SetCookie(&http.Cookie{
		Name:     "game_id",
		Value:    gameID,
		Path:     "/",
		SameSite: http.SameSiteNoneMode,
		Secure:   true,
	})

	log.Println("hit2")
	return c.Redirect(302, "/games/dashboard")
}

func (h *Handler) Test(c echo.Context) error {
	someData := time.Now().Unix()
	msg := fmt.Sprintf(" <div> %v </div>", someData)

	return c.HTML(200, msg)
}

// psql statement to constrain game_id to be unique
// ALTER TABLE games ADD CONSTRAINT game_id_unique UNIQUE (game_id);
