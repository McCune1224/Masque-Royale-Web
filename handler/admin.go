package handler

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/lib/pq"
	"github.com/mccune1224/betrayal-widget/data"
	"github.com/mccune1224/betrayal-widget/util"
	"github.com/mccune1224/betrayal-widget/view/page"
)

var CurrentGameRoles = []string{
	"Agent", "Detective", "Gunman", "Lawyer", "Nurse", "Seraph", "Empress", "Succubus", "Wraith", "Actress", "Assassin", "Highwayman", "Jester", "Sommelier", "Witchdoctor",
}

type ComplexAction struct {
	data.PlayerRequest
	data.Action
}

func (h *Handler) AdminDashboardPage(c echo.Context) error {
	game, _ := util.GetGame(c)
	players, err := h.models.Players.GetAllComplexByGameID(game.GameID)
	if err != nil {
		return TemplRender(c, page.Error500(c, err))
	}

	playerRequests, err := h.models.Actions.GetAllPlayerActionsForGame(game.GameID)
	if err != nil {
		return TemplRender(c, page.Error500(c, err))
	}

	targetActionIDS := make(pq.Int64Array, len(playerRequests))
	for _, a := range playerRequests {
		targetActionIDS = append(targetActionIDS, int64(a.ActionID))
	}

	actions := []data.Action{}
	err = h.models.Actions.Select(&actions, "SELECT * from actions WHERE id = ANY($1)", targetActionIDS)
	if err != nil {
		return TemplRender(c, page.Error500(c, err))
	}

	cprList := []*data.ComplexPlayerRequest{}
	for _, currRequest := range playerRequests {
		if currRequest.Approved {
			continue
		}
		cpr := &data.ComplexPlayerRequest{}
		for _, currPlayer := range players {
			if currRequest.PlayerID == currPlayer.P.ID {
				cpr.P = *currPlayer
				break
			}
		}
		cpr.R = *currRequest
		for _, action := range actions {
			if action.ID == cpr.R.ActionID {
				cpr.A = action
			}
		}
		cprList = append(cprList, cpr)
	}

	sortedCprList := sortComplexPlayerRequest(cprList)
	return TemplRender(c, page.AdminDashboard(c, game, players, CurrentGameRoles, sortedCprList))
}

func (h *Handler) ApprovePlayerAction(c echo.Context) error {
	actionID := util.ParamInt(c, "action_id", -1)
	request, err := h.models.Actions.GetPlayerRequest(actionID)
	if err != nil {
		return TemplRender(c, page.Error500(c, err))
	}

	log.Println("HIT")
	request.Approved = true
	err = h.models.Actions.ApprovePlayerRequest(request.ID)
	if err != nil {
		return TemplRender(c, page.Error500(c, err))
	}
	log.Println("HIT2")
	return nil
}
