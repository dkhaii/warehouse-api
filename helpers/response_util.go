package helpers

import (
	"github.com/dkhaii/warehouse-api/models"
	"github.com/labstack/echo/v4"
)

func CreateResponse(app echo.Context, statusCode int, data interface{}) error {
	return app.JSON(statusCode, models.WebResponse{
		Code:   statusCode,
		Status: "SUCCESS",
		Data:   data,
	})
}

func CreateResponseError(app echo.Context, statusCode int, err error) error {
	return app.JSON(statusCode, models.WebResponse{
		Code:   statusCode,
		Status: "FAIL",
		Data:   err.Error(),
	})
}
