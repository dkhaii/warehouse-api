package controller

import (
	"context"
	"net/http"
	"time"

	"github.com/dkhaii/warehouse-api/helpers"
	"github.com/dkhaii/warehouse-api/models"
	"github.com/dkhaii/warehouse-api/services"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type ItemController struct {
	ItemService services.ItemService
}

func NewItemController(itemService services.ItemService) ItemController {
	return ItemController{
		ItemService: itemService,
	}
}

func (controller *ItemController) Create(app echo.Context) error {
	var request models.CreateItemRequest
	err := app.Bind(&request)
	if err != nil {
		return helpers.CreateResponseError(app, http.StatusBadRequest, err)
	}

	response, err := controller.ItemService.Create(app.Request().Context(), request)
	if err != nil {
		return helpers.CreateResponseError(app, http.StatusNotFound, err)
	}
	return helpers.CreateResponse(app, http.StatusFound, response)
}

func (controller *ItemController) GetItem(app echo.Context) error {
	var queryParam models.GetItemRequest
	err := app.Bind(&queryParam)
	if err != nil {
		return app.JSON(http.StatusBadRequest, models.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "FAIL",
			Data:   err.Error(),
		})
	}

	ctx, cancle := context.WithTimeout(app.Request().Context(), 10*time.Second)
	defer cancle()

	if queryParam.ID != uuid.Nil {
		response, err := controller.ItemService.GetByID(ctx, queryParam.ID)
		if err != nil {
			if err == context.DeadlineExceeded {
				return app.JSON(http.StatusRequestTimeout, models.WebResponse{
					Code:   http.StatusGatewayTimeout,
					Status: "FAIL",
					Data:   helpers.ErrRequestTimedOut,
				})
			}
			return app.JSON(http.StatusNotFound, models.WebResponse{
				Code:   http.StatusNotFound,
				Status: "FAIL",
				Data:   err.Error(),
			})
		}
		return app.JSON(http.StatusFound, models.WebResponse{
			Code:   http.StatusFound,
			Status: "SUCCESS",
			Data:   response,
		})
	}

	if queryParam.Name != "" {
		response, err := controller.ItemService.GetByName(ctx, queryParam.Name)
		if err != nil {
			if err == context.DeadlineExceeded {
				return app.JSON(http.StatusRequestTimeout, models.WebResponse{
					Code:   http.StatusGatewayTimeout,
					Status: "FAIL",
					Data:   helpers.ErrRequestTimedOut,
				})
			}
			return app.JSON(http.StatusNotFound, models.WebResponse{
				Code:   http.StatusNotFound,
				Status: "FAIL",
				Data:   err.Error(),
			})
		}
		return app.JSON(http.StatusFound, models.WebResponse{
			Code:   http.StatusFound,
			Status: "SUCCESS",
			Data:   response,
		})
	}

	response, err := controller.ItemService.GetAll(ctx)
	if err != nil {
		if err == context.DeadlineExceeded {
			return app.JSON(http.StatusRequestTimeout, models.WebResponse{
				Code:   http.StatusGatewayTimeout,
				Status: "FAIL",
				Data:   helpers.ErrRequestTimedOut,
			})
		}
		return app.JSON(http.StatusNotFound, models.WebResponse{
			Code:   http.StatusNotFound,
			Status: "FAIL",
			Data:   err.Error(),
		})
	}
	return app.JSON(http.StatusFound, models.WebResponse{
		Code:   http.StatusFound,
		Status: "SUCCESS",
		Data:   response,
	})
}

func (controller *ItemController) GetCompleteByID(app echo.Context) error {
	var urlParam models.GetItemByIDParamRequest
	err := app.Bind(&urlParam)
	if err != nil {
		return app.JSON(http.StatusBadRequest, models.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "FAIL",
			Data:   err.Error(),
		})
	}

	ctx, cancle := context.WithTimeout(app.Request().Context(), 10*time.Second)
	defer cancle()

	if urlParam.ID != uuid.Nil {
		response, err := controller.ItemService.GetCompleteByID(ctx, urlParam.ID)
		if err != nil {
			if err == context.DeadlineExceeded {
				return app.JSON(http.StatusRequestTimeout, models.WebResponse{
					Code:   http.StatusGatewayTimeout,
					Status: "FAIL",
					Data:   helpers.ErrRequestTimedOut,
				})
			}
			return app.JSON(http.StatusNotFound, models.WebResponse{
				Code:   http.StatusNotFound,
				Status: "FAIL",
				Data:   err.Error(),
			})
		}
		return app.JSON(http.StatusFound, models.WebResponse{
			Code:   http.StatusFound,
			Status: "SUCCESS",
			Data:   response,
		})
	}
	return app.JSON(http.StatusFound, models.WebResponse{
		Code:   http.StatusFound,
		Status: "FAIL",
		Data:   nil,
	})
}

func (controller *ItemController) Update(app echo.Context) error {
	var request models.UpdateItemRequest
	err := app.Bind(&request)
	if err != nil {
		return app.JSON(http.StatusBadRequest, models.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "FAIL",
			Data:   err.Error(),
		})
	}

	response, err := controller.ItemService.Update(app.Request().Context(), request)
	if err != nil {
		return app.JSON(http.StatusNotFound, models.WebResponse{
			Code:   http.StatusNotFound,
			Status: "FAIL",
			Data:   err.Error(),
		})
	}
	return app.JSON(http.StatusOK, models.WebResponse{
		Code:   http.StatusOK,
		Status: "SUCCESS",
		Data:   response,
	})
}

func (controller *ItemController) Delete(app echo.Context) error {
	var urlParam models.GetItemByIDParamRequest
	err := app.Bind(&urlParam)
	if err != nil {
		return app.JSON(http.StatusBadRequest, models.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "FAIL",
			Data:   err.Error(),
		})
	}

	err = controller.ItemService.Delete(app.Request().Context(), urlParam.ID)
	if err != nil {
		return app.JSON(http.StatusNotFound, models.WebResponse{
			Code:   http.StatusNotFound,
			Status: "FAIL",
			Data:   err.Error(),
		})
	}
	return app.JSON(http.StatusOK, models.WebResponse{
		Code:   http.StatusOK,
		Status: "SUCCESS",
		Data:   nil,
	})
}
