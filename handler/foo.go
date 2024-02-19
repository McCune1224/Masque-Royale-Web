package handler

import (
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/mccune1224/betrayal-widget/views/components"
)

func (h *Handler) Resize(c echo.Context) error {
	size := c.FormValue("size")
	if size == "" {
		size = "20"
	}

	iSize, err := strconv.Atoi(size)
	if err != nil {
		iSize = 20
	}

	return TemplRender(c, components.Box(iSize))
}
