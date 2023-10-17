package controller

import (
	"net/http"

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
		return app.JSON(http.StatusBadRequest, models.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "FAIL",
			Data:   err.Error(),
		})
	}

	response, err := controller.CategoryService.Create(request)
	if err != nil {
		return app.JSON(http.StatusInternalServerError, models.WebResponse{
			Code:   http.StatusInternalServerError,
			Status: "FAIL",
			Data:   err.Error(),
		})
	}

	return app.JSON(http.StatusCreated, models.WebResponse{
		Code:   http.StatusCreated,
		Status: "SUCCESS",
		Data:   response,
	})
}

func (controller *CategoryController) GetCategory(app echo.Context) error {
	var queryParam models.GetCategoryRequest
	err := app.Bind(&queryParam)
	if err != nil {
		return app.JSON(http.StatusBadRequest, models.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "FAIL",
			Data:   err.Error(),
		})
	}

	if queryParam.ID != "" {
		response, err := controller.CategoryService.GetByID(queryParam.ID)
		if err != nil {
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
		response, err := controller.CategoryService.GetByName(queryParam.Name)
		if err != nil {
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

	response, err := controller.CategoryService.GetAll()
	if err != nil {
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

func (controller *CategoryController) Update(app echo.Context) error {
	var request models.UpdateCategoryRequest
	err := app.Bind(&request)
	if err != nil {
		return app.JSON(http.StatusBadRequest, models.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "FAIL",
			Data:   err.Error(),
		})
	}

	err = controller.CategoryService.Update(request)
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
		Data:   request,
	})
}

func (controller *CategoryController) Delete(app echo.Context) error {
	var urlParam models.GetCategoryIDRequest
	err := app.Bind(&urlParam)
	if err != nil {
		return app.JSON(http.StatusBadRequest, models.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "FAIL",
			Data:   err.Error(),
		})
	}

	err = controller.CategoryService.Delete(urlParam.ID)
	if err != nil {
		return app.JSON(http.StatusInternalServerError, models.WebResponse{
			Code:   http.StatusInternalServerError,
			Status: "FAIL",
			Data:   err.Error(),
		})
	}

	return app.JSON(http.StatusOK, models.WebResponse{
		Code:   http.StatusOK,
		Status: "SUCCESS",
		Data:   "",
	})
}
