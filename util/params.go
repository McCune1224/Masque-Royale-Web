package util

import (
	"strconv"

	"github.com/labstack/echo/v4"
)

func QueryParamInt(c echo.Context, name string, def int) int {
	param := c.QueryParam(name)
	result, err := strconv.Atoi(param)
	if err != nil {
		return def
	}
	return result
}

func ParamInt(c echo.Context, name string, def int) int {
	param := c.Param(name)
	result, err := strconv.Atoi(param)
	if err != nil {
		return def
	}
	return result
}
