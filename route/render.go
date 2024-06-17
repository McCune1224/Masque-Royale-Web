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
	games.PUT("/:game_id", handler.UpdateGame)
	games.PATCH("/:game_id", handler.UpdateGame)
	games.DELETE("/:game_id", handler.DeleteGame)

	roles := api.Group("/roles")
	roles.GET("", handler.GetAllRoles)
	roles.GET("/complete", handler.GetAllCompleteRoles)
	roles.GET("/:role_id", handler.GetRoleByID)
	roles.GET("/:role_id/complete", handler.GetCompleteRole)
	roles.GET("/:role_id/abilities", handler.GetRoleAbilities)
	roles.GET("/:role_id/passives", handler.GetRolePassives)

	// roles.POST("", handler.InsertRole

	categories := api.Group("/categories")
	// categories.GET("", handler.GetAllCategories)
	categories.GET("/:category_id", handler.GetCategoryByID)

	abilities := api.Group("/abilities")
	abilities.GET("", handler.GetAllAbilities)
	abilities.GET("/:ability_id", handler.GetAbilityByID)
	abilities.GET("/:ability_id/role", handler.GetRoleForAbility)
	abilities.GET("/search", handler.GetAbilityByName)

	anyAbilities := api.Group("/any_abilities")
	anyAbilities.GET("", handler.GetAllAnyAbilities)
	// anyAbilities.GET("/:any_ability_id", handler.GetAnyAbilityByID)
	// anyAbilities.GET("/search", handler.GetAnyAbilityByName)

	statuses := api.Group("/statuses")
	statuses.GET("", handler.GetAllStatuses)

	players := games.Group("/:game_id/players")
	players.GET("", handler.GetAllPlayers)
	players.POST("", handler.InsertPlayer)
	players.GET("/:player_id", handler.GetPlayerByID)
	players.PUT("/:player_id", handler.UpdatePlayer)
	players.PATCH("/:player_id", handler.UpdatePlayer)
	players.DELETE("/:player_id", handler.DeletePlayer)
	players.GET("/:player_id/actions", handler.GetPlayerActions)

	actions := games.Group("/:game_id/actions")
	actions.GET("", handler.GetAllActions)
	actions.POST("", handler.InsertAction)
	actions.PUT("/:action_id", handler.UpdateAction)
	actions.PATCH("/:action_id", handler.UpdateAction)
	// actions.GET("/:action_id", handler.GetActionByID)
	actions.DELETE("/:action_id", handler.DeleteAction)

	admin := games.Group("/:game_id/admin")
	admin.POST("/sync-roles-csv", handler.SyncRolesCsv)
	admin.POST("/sync-any-abilities-csv", handler.SyncStatusDetailsCSV)
}
