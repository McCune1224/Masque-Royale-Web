package util

import (
	"errors"
	"strconv"

	"github.com/jackc/pgconn"
	"github.com/labstack/echo/v4"
)

func ParseInt32Param(c echo.Context, str string) (int32, error) {
	val := c.Param(str)
	if val == "" {
		return -1, errors.New("missing param value")
	}

	i, err := strconv.ParseInt(val, 10, 32)
	return int32(i), err
}

func ParsePgError(err error) *pgconn.PgError {
	var pgErr *pgconn.PgError
	if ok := errors.As(err, &pgErr); ok {
		return pgErr
	}
	return nil
}
