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

	roles := api.Group("/roles")
	roles.GET("", handler.GetAllRoles)
	roles.GET("/complete", handler.GetAllCompleteRoles)
	roles.GET("/:role_id", handler.GetRoleByID)
	roles.GET("/:role_id/complete", handler.GetCompleteRole)
	roles.GET("/:role_id/abilities", handler.GetRoleAbilities)
	roles.GET("/:role_id/passives", handler.GetRolePassives)

	games := api.Group("/games")
	games.GET("/random", handler.GetRandomGame)
	games.POST("", handler.InsertGame)
	games.GET("", handler.GetAllGames)
	games.GET("/:game_id", handler.GetGameByID)
	games.PUT("/:game_id", handler.UpdateGame)
	games.DELETE("/:game_id", handler.DeleteGame)

	gamesPlayers := games.Group("/:game_id/players")
	gamesPlayers.GET("", handler.GetAllGamePlayers)
	gamesPlayers.POST("", handler.InsertPlayer)
	gamesPlayers.GET("/:player_id", handler.GetPlayerByID)
	gamesPlayers.PUT("/:player_id", handler.UpdatePlayer)
	gamesPlayers.GET("/:player_id/notes", handler.GetPlayerNotes)
	gamesPlayers.PUT("/:player_id/notes", handler.UpdatePlayerNotes)
	gamesPlayers.GET("/:player_id/abilities", handler.GetPlayerAbilities)
	gamesPlayers.GET("/:player_id/abilities/:ability_id", handler.GetPlayerAbility)
	gamesPlayers.POST("/:player_id/abilities/:ability_id", handler.CreatePlayerAbility)
	gamesPlayers.PUT("/:player_id/abilities/:ability_id", handler.UpdatePlayerAbility)
	gamesPlayers.DELETE("/:player_id", handler.DeletePlayer)

	categories := api.Group("/categories")
	// categories.GET("", handler.GetAllCategories)
	categories.GET("/:category_id", handler.GetCategoryByID)

	abilities := api.Group("/abilities")
	abilities.GET("", handler.GetAllAbilities)
	abilities.GET("/:ability_id", handler.GetAbilityByID)
	abilities.GET("/:ability_id/role", handler.GetRoleForAbility)
	abilities.GET("/search", handler.GetAbilityByName)

	gameAdmin := games.Group("/:game_id/admin")
	gameAdmin.POST("/sync-roles-csv", handler.SyncRolesCsv)
	gameAdmin.POST("/sync-any-abilities-csv", handler.SyncStatusDetailsCSV)

	statuses := api.Group("/statuses")
	statuses.GET("", handler.GetAllStatuses)
}
