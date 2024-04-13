package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/mccune1224/betrayal-widget/util"
	"github.com/mccune1224/betrayal-widget/view/page"
)

var CurrentGameRoles = []string{
	"Agent", "Detective", "Gunman", "Lawyer", "Nurse", "Seraph", "Empress", "Succubus", "Wraith", "Actress", "Assassin", "Highwayman", "Jester", "Sommelier", "Witchdoctor",
}

func (h *Handler) AdminDashboardPage(c echo.Context) error {
	game, _ := util.GetGame(c)
	players, err := h.models.Players.GetAllComplexByGameID(game.GameID)
	if err != nil {
		return TemplRender(c, page.Error500(c, err))
	}

	return TemplRender(c, page.AdminDashboard(c, players, CurrentGameRoles))
}
