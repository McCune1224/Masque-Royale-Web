package route

import (
	"github.com/labstack/echo/v4"
	"github.com/mccune1224/betrayal-widget/handler"
)

func ViewRoutes(app *echo.Echo, handler *handler.Handler) {
	app.GET("/", handler.IndexPage)
}
