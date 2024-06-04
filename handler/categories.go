package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/mccune1224/betrayal-widget/models"
	"github.com/mccune1224/betrayal-widget/util"
)

func (h *Handler) GetCategoryByID(c echo.Context) error {
	categoryId, err := util.ParseInt32Param(c, "category_id")
	if err != nil {
		return util.BadRequestJson(c, "Invalid Category ID")
	}
	q := models.New(h.Db)
	category, err := q.GetCategoryByID(c.Request().Context(), int32(categoryId))
	if err != nil {
		return util.InternalServerErrorJson(c, err.Error())
	}
	return c.JSON(200, category)
}
