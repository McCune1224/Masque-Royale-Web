package util

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// c.JSON(400 echo.Map{"message": ...} shorthand
func BadRequestJson(c echo.Context, msg string) error {
	return c.JSON(http.StatusBadRequest, echo.Map{"message": msg})
}

// c.JSON(401 echo.Map{"message": ...} shorthand
func UnauthorizedJson(c echo.Context, msg string) error {
	return c.JSON(http.StatusUnauthorized, echo.Map{"message": msg})
}

// c.JSON(404 echo.Map{"message": ...} shorthand
func NotFoundJson(c echo.Context, msg string) error {
	return c.JSON(http.StatusNotFound, echo.Map{"message": msg})
}

// c.JSON(500 echo.Map{"error": ...} shorthand
func InternalServerErrorJson(c echo.Context, msg string) error {
	return c.JSON(http.StatusInternalServerError, echo.Map{"error": msg})
}
