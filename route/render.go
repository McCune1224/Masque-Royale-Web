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
	games := api.Group("/games")
	games.GET("/random", handler.GetRandomGame)
	games.GET("/all", handler.GetAllGames)
	games.GET("/:game_id", handler.GetGameByID)

}
