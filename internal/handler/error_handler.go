package handler

import (
	"net/http"

	"rename-service-name-here/internal/util"

	"github.com/labstack/echo/v4"
)

func errorHandler(c echo.Context, err error) error {
	switch e := err.(type) {
	case util.AppErrors:
		c.JSON(e.Errors[0].Status, e)
	case error:
		c.JSON(http.StatusInternalServerError, e)
	}

	return nil
}
