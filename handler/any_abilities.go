package handler

import (
	"github.com/labstack/echo/v4"
)

func (h *Handler) GetAllAnyAbilities(c echo.Context) error {
	// q := models.New(h.Db)
	// anyAbilities, err := q.ListAnyAbilityDetails(c.Request().Context())
	// if err != nil {
	// 	log.Println(anyAbilities)
	// 	return util.InternalServerErrorJson(c, err.Error())
	// }
	return c.JSON(500, "Not Implemented")
}
