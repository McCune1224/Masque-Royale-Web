package route

import (
	"github.com/labstack/echo/v4"
	"github.com/mccune1224/betrayal-widget/handler"
)

func ViewRoutes(app *echo.Echo, handler *handler.Handler) {
	app.GET("/", handler.IndexPage)

	app.GET("/new", handler.NewGamesPage)

	app.GET("/join/:game_id", handler.JoinGamePage)

	app.GET("/games/:game_id", handler.GameDashboardPage)
	app.GET("/games/:game_id/players/:player_name", handler.PlayerDashboardPage)
}
