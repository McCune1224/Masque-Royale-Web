package route

import (
	"time"

	"github.com/labstack/echo/v4"
	"github.com/mccune1224/betrayal-widget/handler"
)

func Routes(app *echo.Echo, handler *handler.Handler) {
	app.GET("healthcheck", func(c echo.Context) error {
		start := time.Now()
		return c.JSON(200,
			echo.Map{
				"response_time": time.Since(start),
			},
		)
	})

	// auth := app.Group("/auth")

	api := app.Group("/api")

	room := api.Group("/rooms")
	room.GET("", handler.GetAllRooms)
	room.GET("/:room_id", handler.GetRoomByID)
	room.POST("", handler.GetRoomByName)

	games := api.Group("/games")
	games.GET("/random", handler.GetRandomGame)
	games.POST("", handler.InsertGame)
	games.GET("", handler.GetAllGames)
	games.GET("/:game_id", handler.GetGameByID)

	roles := api.Group("/roles")
	roles.GET("", handler.GetAllRoles)
	// roles.GET("/:role_id", handler.GetRoleByID)
	// roles.POST("", handler.InsertRole)

	players := games.Group("/:game_id/players")
	players.GET("", handler.GetAllPlayers)
	players.POST("", handler.InsertPlayer)
	players.GET("/:player_id", handler.GetPlayerByID)
	players.PUT("/:player_id", handler.UpdatePlayer)
	players.PATCH("/:player_id", handler.UpdatePlayer)

	admin := games.Group("/:game_id/admin")
	admin.POST("/sync-roles-csv", handler.SyncRolesCsv)
}
