package controller

import (
	"net/http"
	"time"

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

func (controller *ItemController) Routes(app *echo.Echo) {
	app.POST("/api/items/add", controller.Create)
}

func (controller *ItemController) Create(app echo.Context) error {
	var request models.CreateItemRequest
	err := app.Bind(&request)
	if err != nil {
		return err
	}

	request.ID = uuid.New()
	request.CreatedAt = time.Now()
	request.UpdatedAt = request.CreatedAt

	response, err := controller.ItemService.Create(request)
	if err != nil {
		return app.JSON(http.StatusBadRequest, models.WebResponse{
			Code:   http.StatusBadRequest,
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

	if queryParam.ID != uuid.Nil {
		response, err := controller.ItemService.GetByID(queryParam.ID)
		if err != nil {
			return app.JSON(http.StatusInternalServerError, models.WebResponse{
				Code:   http.StatusInternalServerError,
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
		response, err := controller.ItemService.GetByName(queryParam.Name)
		if err != nil {
			return app.JSON(http.StatusInternalServerError, models.WebResponse{
				Code:   http.StatusInternalServerError,
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

	response, err := controller.ItemService.GetAll()
	if err != nil {
		return app.JSON(http.StatusInternalServerError, models.WebResponse{
			Code:   http.StatusInternalServerError,
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
