package middleware

import (
	"log"

	"github.com/labstack/echo/v4"
)

func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		pw, _ := c.Cookie("password")
		log.Println(pw.Name)
		if pw != nil && pw.Value != "sussy" {
			return c.Redirect(302, "/")
		}
		return next(c)
	}
}
