package handler

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/mccune1224/betrayal-widget/models"
	"github.com/mccune1224/betrayal-widget/util"
)

func (h *Handler) GetAllAbilities(c echo.Context) error {
	q := models.New(h.Db)

	abilities, err := q.ListAbilityDetails(c.Request().Context())
	if err != nil {
		return echo.NewHTTPError(500, err.Error())
	}

	return c.JSON(200, abilities)
}

func (h *Handler) GetAbilityByID(c echo.Context) error {
	abilityId, err := util.ParseInt32Param(c, "ability_id")
	if err != nil {
		return util.BadRequestJson(c, "Invalid Ability ID")
	}
	q := models.New(h.Db)
	ability, err := q.GetAbilityDetail(c.Request().Context(), int32(abilityId))
	if err != nil {
		return util.InternalServerErrorJson(c, err.Error())
	}
	return c.JSON(200, ability)
}

func (h *Handler) GetAbilityByName(c echo.Context) error {
	q := models.New(h.Db)
	ability, err := q.GetAbilityByName(c.Request().Context(), c.QueryParam("name"))
	if err != nil {
		return util.InternalServerErrorJson(c, err.Error())
	}
	return c.JSON(200, ability)
}

func (h *Handler) GetRoleForAbility(c echo.Context) error {
	abilityId, err := util.ParseInt32Param(c, "ability_id")
	if err != nil {
		log.Println(err)
		return util.InternalServerErrorJson(c, err.Error())
	}
	q := models.New(h.Db)
	role, err := q.GetRoleFromAbilityID(c.Request().Context(), abilityId)
	if err != nil {
		log.Println(err)
		return util.InternalServerErrorJson(c, err.Error())
	}
	return c.JSON(200, role)
}
