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

func NewUserController(userService service.UserService) UserController {
	return UserController{
		UserService: userService,
	}
}

func (controller *UserController) Routes(app *echo.Echo) {
	app.POST("/api/user/register", controller.Create)
	app.GET("/api/user/:id", controller.GetByID)
}

func (controller *UserController) Create(app echo.Context) error {
	var request model.CreateUserRequest

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
		return app.JSON(http.StatusBadRequest, model.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "FAIL",
			Data:   err.Error(),
		})
	}
	
	userID := uuid.New()
	createdAt := time.Now()

	request.ID = userID
	request.Role = 1
	request.CreatedAt = createdAt
	request.UpdatedAt = request.CreatedAt

	response, err := controller.UserService.Create(request)
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

func (controller *UserController) GetByID(app echo.Context) error {
	var urlParam model.GetUserIDRequest

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

	err := app.Bind(&urlParam)
	if err != nil {
		return app.JSON(http.StatusBadRequest, model.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "FAIL",
			Data:   err.Error(),
		})
	}

	if err != nil {
		return app.JSON(http.StatusBadRequest, model.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "FAIL",
			Data:   err.Error(),
		})
	}

	response, err := controller.UserService.GetByID(urlParam.ID)
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
