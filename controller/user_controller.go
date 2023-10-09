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
	app.POST("/api/users/register", controller.Create)
	app.GET("/api/users", controller.GetWithOptions)
	app.PUT("/api/users/:id", controller.Update)
}

func (controller *UserController) Create(app echo.Context) error {
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

	var request model.CreateUserRequest
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

func (controller *UserController) GetWithOptions(app echo.Context) error {
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

	var queryParam model.GetUserRequest
	err := app.Bind(&queryParam)
	if err != nil {
		return app.JSON(http.StatusBadRequest, model.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "FAIL",
			Data:   err.Error(),
		})
	}

	if queryParam.ID != uuid.Nil {
		response, err := controller.UserService.GetByID(queryParam.ID)
		if err != nil {
			return app.JSON(http.StatusNotFound, model.WebResponse{
				Code:   http.StatusNotFound,
				Status: "FAIL",
				Data:   err.Error(),
			})
		}

		return app.JSON(http.StatusFound, model.WebResponse{
			Code:   http.StatusFound,
			Status: "SUCCESS",
			Data:   response,
		})
	}

	if queryParam.Username != "" {
		response, err := controller.UserService.GetByUsername(queryParam.Username)
		if err != nil {
			return app.JSON(http.StatusNotFound, model.WebResponse{
				Code:   http.StatusNotFound,
				Status: "FAIL",
				Data:   err.Error(),
			})
		}

		return app.JSON(http.StatusFound, model.WebResponse{
			Code:   http.StatusFound,
			Status: "SUCCESS",
			Data:   response,
		})
	}

	response, err := controller.UserService.GetAll()
	if err != nil {
		return app.JSON(http.StatusNotFound, model.WebResponse{
			Code:   http.StatusNotFound,
			Status: "FAIL",
			Data:   err.Error(),
		})
	}

	return app.JSON(http.StatusFound, model.WebResponse{
		Code:   http.StatusFound,
		Status: "SUCCESS",
		Data:   response,
	})
}

func (controller *UserController) Update(app echo.Context) error {
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

	var request model.UpdateUserRequest
	err := app.Bind(&request)
	if err != nil {
		return app.JSON(http.StatusBadRequest, model.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "FAIL",
			Data:   err.Error(),
		})
	}

	updatedAt := time.Now()
	request.UpdatedAt = updatedAt

	err = controller.UserService.Update(request)
	if err != nil {
		return app.JSON(http.StatusNotFound, model.WebResponse{
			Code:   http.StatusNotFound,
			Status: "FAIL",
			Data:   err.Error(),
		})
	}

	return app.JSON(http.StatusOK, model.WebResponse{
		Code:   http.StatusOK,
		Status: "SUCCESS",
		Data:   request,
	})
}
