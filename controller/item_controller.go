package controller

import (
	"net/http"
	"time"

	"github.com/dkhaii/warehouse-api/model"
	"github.com/dkhaii/warehouse-api/service"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type ItemController struct {
	ItemService service.ItemService
}

func NewItemController(itemService service.ItemService) ItemController {
	return ItemController{
		ItemService: itemService,
	}
}

func (controller *ItemController) Routes(app *echo.Echo) {
	app.POST("/api/item", controller.Create)
}

func (controller *ItemController) Create(app echo.Context) error {
	var request model.CreateItemRequest

	defer func() {
		err := recover()
		if err != nil {
			app.JSON(http.StatusInternalServerError, model.WebResponse{
				Code:   http.StatusInternalServerError,
				Status: "FAIL",
				Data:   err,
			})
		}
	}()

	err := app.Bind(&request)
	if err != nil {
		return err
	}

	request.ID = uuid.New()
	request.CreatedAt = time.Now()
	request.UpdatedAt = request.CreatedAt

	response, err := controller.ItemService.Create(request)
	if err != nil {
		return app.JSON(http.StatusBadRequest, model.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "FAIL",
			Data:   err.Error(),
		})
	}

	return app.JSON(http.StatusOK, model.WebResponse{
		Code:   http.StatusOK,
		Status: "SUCCESS",
		Data:   response,
	})
}
