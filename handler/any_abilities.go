package handler

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/mccune1224/betrayal-widget/models"
	"github.com/mccune1224/betrayal-widget/util"
)

func (h *Handler) GetAllAnyAbilities(c echo.Context) error {
	q := models.New(h.Db)
	anyAbilities, err := q.ListAnyAbilityDetails(c.Request().Context())
	if err != nil {
		log.Println(anyAbilities)
		return util.InternalServerErrorJson(c, err.Error())
	}
	return c.JSON(200, anyAbilities)
}
