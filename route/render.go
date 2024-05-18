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

}
