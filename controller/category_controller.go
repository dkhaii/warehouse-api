package controller

import (
	"context"
	"net/http"
	"time"

	"github.com/dkhaii/warehouse-api/helpers"
	"github.com/dkhaii/warehouse-api/models"
	"github.com/dkhaii/warehouse-api/services"
	"github.com/labstack/echo/v4"
)

type CategoryController struct {
	CategoryService services.CategoryService
}

func NewCategoryController(categoryService services.CategoryService) CategoryController {
	return CategoryController{
		CategoryService: categoryService,
	}
}

func (controller *CategoryController) Create(app echo.Context) error {
	var request models.CreateCategoryRequest
	err := app.Bind(&request)
	if err != nil {
		return helpers.CreateResponseError(app, http.StatusBadRequest, err)
	}

	response, err := controller.CategoryService.Create(app.Request().Context(), request)
	if err != nil {
		return helpers.CreateResponseError(app, http.StatusInternalServerError, err)
	}
	return helpers.CreateResponse(app, http.StatusCreated, response)
}

func (controller *CategoryController) GetCategory(app echo.Context) error {
	var queryParam models.GetCategoryRequest
	err := app.Bind(&queryParam)
	if err != nil {
		return helpers.CreateResponseError(app, http.StatusBadRequest, err)
	}

	ctx, cancel := context.WithTimeout(app.Request().Context(), 10*time.Second)
	defer cancel()

	if queryParam.ID != "" {
		response, err := controller.CategoryService.GetByID(ctx, queryParam.ID)
		if err != nil {
			if err == context.DeadlineExceeded {
				return helpers.CreateResponseError(app, http.StatusRequestTimeout, helpers.ErrRequestTimedOut)
			}
			return helpers.CreateResponseError(app, http.StatusNotFound, err)
		}
		return helpers.CreateResponse(app, http.StatusFound, response)
	}

	if queryParam.Name != "" {
		response, err := controller.CategoryService.GetByName(ctx, queryParam.Name)
		if err != nil {
			if err == context.DeadlineExceeded {
				return helpers.CreateResponseError(app, http.StatusRequestTimeout, helpers.ErrRequestTimedOut)
			}
			return helpers.CreateResponseError(app, http.StatusNotFound, err)
		}
		return helpers.CreateResponse(app, http.StatusFound, response)
	}

	response, err := controller.CategoryService.GetAll(ctx)
	if err != nil {
		if err == context.DeadlineExceeded {
			return helpers.CreateResponseError(app, http.StatusRequestTimeout, helpers.ErrRequestTimedOut)
		}
		return helpers.CreateResponseError(app, http.StatusNotFound, err)
	}
	return helpers.CreateResponse(app, http.StatusFound, response)
}

func (controller *CategoryController) Update(app echo.Context) error {
	var request models.UpdateCategoryRequest
	err := app.Bind(&request)
	if err != nil {
		return helpers.CreateResponseError(app, http.StatusBadRequest, err)
	}

	response, err := controller.CategoryService.Update(app.Request().Context(), request)
	if err != nil {
		return helpers.CreateResponseError(app, http.StatusNotFound, err)
	}
	return helpers.CreateResponse(app, http.StatusOK, response)
}

func (controller *CategoryController) Delete(app echo.Context) error {
	var urlParam models.GetCategoryIDRequest
	err := app.Bind(&urlParam)
	if err != nil {
		return helpers.CreateResponse(app, http.StatusBadRequest, err)
	}

	err = controller.CategoryService.Delete(app.Request().Context(), urlParam.ID)
	if err != nil {
		return helpers.CreateResponseError(app, http.StatusNotFound, err)
	}
	return helpers.CreateResponse(app, http.StatusOK, nil)
}
