package middleware

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mccune1224/betrayal-widget/handler"
	"github.com/mccune1224/betrayal-widget/views/errors"
)

// Wrapper for echo.DefaultHTTPErrorHandler
// Automatically renders the templ error page for 404 and 500 errors
func TemplHTTPErrorHandler(err error, c echo.Context) {
	code := http.StatusInternalServerError
	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
	}
	c.Logger().Error("[TEMPL]:", err)

	switch code {
	case http.StatusNotFound:
		handler.TemplRender(c, errors.Error404())
	case http.StatusInternalServerError:
		handler.TemplRender(c, errors.Error404())

	default:
		c.Echo().DefaultHTTPErrorHandler(err, c)
	}
}
