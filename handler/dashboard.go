package handler

import (
	"log"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/mccune1224/betrayal-widget/views"
)

func (h *Handler) Dashboard(c echo.Context) error {
	game_id := c.Param("game_id")
	game, err := h.models.Games.GetByGameID(game_id)
	if err != nil || game == nil {
		log.Println(err)
		return c.Redirect(302, "/")
	}
	// players, err := h.models.Players.GetByGameID(game_id)
	// if err != nil {
	// 	log.Println(err)
	// 	return c.Redirect(302, "/")
	// }
	// Set value within the context for the game_id
	c.Set("game_id", game_id)
	return TemplRender(c, views.Home(c, game.PlayerCount, rotateCSSGenerator(game.PlayerCount)))
}

func rotateCSSGenerator(playerCount int) []string {
	res := []string{}

	for i := 0; i < playerCount; i++ {
		foo := strconv.Itoa(360 / playerCount * i)
		bar := "absolute transform h-64 w-1  bg-transparent rotate-[" + foo + "deg]"
		res = append(res, bar)
	}
	return res
}

// res = append(res, "h-8 w-8 bg-orange-500 rotate-["+strconv.Itoa(360/playerCount*i)+"deg]")
// res = append(res, "transform h-8 w-8 bg-orange-500 rounded-full rotate-[315deg] translate-x-"+strconv.Itoa(i))
// res = append(res, fmt.Sprintf("h-8 w-8 bg-orange-500 rotate-[%ddeg]", 360/playerCount*i))
