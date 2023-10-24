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

	currentUserToken, err := helpers.GetSplitedToken(app)
	if err != nil {
		return helpers.CreateResponseError(app, http.StatusUnauthorized, err)
	}

	response, err := controller.ItemService.Create(app.Request().Context(), request, currentUserToken)
	if err != nil {
		return helpers.CreateResponseError(app, http.StatusInternalServerError, err)
	}
	return helpers.CreateResponse(app, http.StatusCreated, response)
}

func (controller *ItemController) GetItem(app echo.Context) error {
	var queryParam models.GetItemRequest
	err := app.Bind(&queryParam)
	if err != nil {
		return helpers.CreateResponseError(app, http.StatusBadRequest, err)
	}

	ctx, cancle := context.WithTimeout(app.Request().Context(), 10*time.Second)
	defer cancle()

	if queryParam.ID != uuid.Nil {
		response, err := controller.ItemService.GetByID(ctx, queryParam.ID)
		if err != nil {
			if err == context.DeadlineExceeded {
				return helpers.CreateResponseError(app, http.StatusRequestTimeout, helpers.ErrRequestTimedOut)
			}
			return helpers.CreateResponseError(app, http.StatusNotFound, err)
		}
		return helpers.CreateResponse(app, http.StatusFound, response)
	}

	if queryParam.Name != "" {
		response, err := controller.ItemService.GetByName(ctx, queryParam.Name)
		if err != nil {
			if err == context.DeadlineExceeded {
				return helpers.CreateResponseError(app, http.StatusRequestTimeout, helpers.ErrRequestTimedOut)
			}
			return helpers.CreateResponseError(app, http.StatusNotFound, err)
		}
		return helpers.CreateResponse(app, http.StatusFound, response)
	}

	response, err := controller.ItemService.GetAll(ctx)
	if err != nil {
		if err == context.DeadlineExceeded {
			return helpers.CreateResponseError(app, http.StatusRequestTimeout, helpers.ErrRequestTimedOut)
		}
		return helpers.CreateResponseError(app, http.StatusNotFound, err)
	}
	return helpers.CreateResponse(app, http.StatusFound, response)
}

func (controller *ItemController) GetCompleteByID(app echo.Context) error {
	var urlParam models.GetItemByIDParamRequest
	err := app.Bind(&urlParam)
	if err != nil {
		return helpers.CreateResponseError(app, http.StatusBadRequest, err)
	}

	ctx, cancle := context.WithTimeout(app.Request().Context(), 10*time.Second)
	defer cancle()

	response, err := controller.ItemService.GetCompleteByID(ctx, urlParam.ID)
	if err != nil {
		if err == context.DeadlineExceeded {
			return helpers.CreateResponseError(app, http.StatusRequestTimeout, helpers.ErrRequestTimedOut)
		}
		return helpers.CreateResponseError(app, http.StatusNotFound, err)
	}
	return helpers.CreateResponse(app, http.StatusFound, response)
}

func (controller *ItemController) Update(app echo.Context) error {
	var request models.UpdateItemRequest
	err := app.Bind(&request)
	if err != nil {
		return helpers.CreateResponseError(app, http.StatusBadRequest, err)
	}

	response, err := controller.ItemService.Update(app.Request().Context(), request)
	if err != nil {
		return helpers.CreateResponseError(app, http.StatusNotFound, err)
	}
	return helpers.CreateResponse(app, http.StatusOK, response)
}

func (controller *ItemController) Delete(app echo.Context) error {
	var urlParam models.GetItemByIDParamRequest
	err := app.Bind(&urlParam)
	if err != nil {
		return helpers.CreateResponseError(app, http.StatusBadRequest, err)
	}

	err = controller.ItemService.Delete(app.Request().Context(), urlParam.ID)
	if err != nil {
		helpers.CreateResponseError(app, http.StatusNotFound, err)
	}
	return helpers.CreateResponse(app, http.StatusOK, nil)
}
