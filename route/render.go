package route

import (
	"github.com/labstack/echo/v4"
	"github.com/mccune1224/betrayal-widget/handler"
)

func Routes(app *echo.Echo, handler *handler.Handler) {
	app.GET("/", handler.IndexPage)

	app.GET("/join/:game_id", handler.JoinGamePage)
	app.POST("/create", handler.CreateGame)
	app.DELETE("/delete/:game_id", handler.DeleteGame)

	games := app.Group("/games/:game_id")
	games.GET("", handler.GameDashboardPage)

	games.POST("/players", handler.AddPlayerToGame)
	games.DELETE("/players/:player_id", handler.DeletePlayerFromGame)
	games.POST("/players/:player_id/notes", handler.UpdatePlayerNotes)

	gamesPlayer := games.Group("/players/:player_id")
	gamesPlayer.GET("", handler.PlayerDashboardPage)
	gamesPlayer.GET("/flashcard", handler.PlayerFlashcard)
	gamesPlayer.POST("/flashcard/roles", handler.SearchRole)
	gamesPlayer.POST("/flashcard/abilities", handler.SearchAbility)
	gamesPlayer.POST("/death", handler.MarkPlayerDead)
	gamesPlayer.POST("/actions", handler.SubmitPlayerAction)
	gamesPlayer.DELETE("/actions/:action_id", handler.DeletePlayerAction)
	gamesPlayer.POST("/actions/:action_id/approve", handler.ApprovePlayerAction)
	gamesPlayer.POST("/actions/:action_id/note", handler.UpdatePlayerActionNote)

	admin := games.Group("/admin")
	admin.GET("", handler.AdminDashboardPage)
	admin.GET("/history", handler.AdminActionHistoryPage)
	admin.DELETE("", handler.AdminDashboardPage)
	admin.POST("/cycle/increment", handler.GamePhaseIncrement)
	admin.POST("/cycle/decrement", handler.GamePhaseDecrement)
}
