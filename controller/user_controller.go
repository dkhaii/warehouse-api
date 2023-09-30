package controller

import (
	"net/http"
	"time"

	"github.com/dkhaii/warehouse-api/model"
	"github.com/dkhaii/warehouse-api/service"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type UserController struct {
	UserService service.UserService
}

func NewUserController(userService *service.UserService) UserController {
	return UserController{
		UserService: *userService,
	}
}

func (controller *UserController) Routes(app *echo.Echo) {
	app.POST("/api/user", controller.Create)
}

func (controller *UserController) Create(app echo.Context) error {
	var request model.CreateUserRequest

	err := app.Bind(&request)
	if err != nil {
		return err
	}

	request.ID = uuid.New()
	request.Role = 1
	request.CreatedAt = time.Now()
	request.UpdatedAt = request.CreatedAt

	response, err := controller.UserService.Create(request)
	if err != nil {
		return app.JSON(http.StatusBadRequest, model.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "FAIL",
			Data:   err,
		})
	}

	return app.JSON(http.StatusOK, model.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   response,
	})
}
