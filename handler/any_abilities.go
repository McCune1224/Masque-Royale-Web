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

func (h *Handler) GetAnyAbilityByID(c echo.Context) error {
	anyAbilityId, err := util.ParseInt32Param(c, "any_ability_id")
	if err != nil {
		return util.BadRequestJson(c, "Invalid Any Ability ID")
	}

	q := models.New(h.Db)
	anyAbility, err := q.GetAnyAbilityDetail(c.Request().Context(), int32(anyAbilityId))
	if err != nil {
		return util.InternalServerErrorJson(c, err.Error())
	}
	return c.JSON(200, anyAbility)
}

func (h *Handler) GetAnyAbilityByName(c echo.Context) error {
	q := models.New(h.Db)
	anyAbility, err := q.GetAnyAbilityByName(c.Request().Context(), c.QueryParam("name"))
	if err != nil {
		return util.InternalServerErrorJson(c, err.Error())
	}
	return c.JSON(200, anyAbility)
}
