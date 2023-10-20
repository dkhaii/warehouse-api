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

type UserController struct {
	UserService services.UserService
}

func NewUserController(userService services.UserService) UserController {
	return UserController{
		UserService: userService,
	}
}

func (controller *UserController) Create(app echo.Context) error {
	var request models.CreateUserRequest
	err := app.Bind(&request)
	if err != nil {
		return app.JSON(http.StatusBadRequest, models.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "FAIL",
			Data:   err.Error(),
		})
	}

	response, err := controller.UserService.Create(app.Request().Context(), request)
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

func (controller *UserController) GetUser(app echo.Context) error {
	var queryParam models.GetUserRequest
	err := app.Bind(&queryParam)
	if err != nil {
		return app.JSON(http.StatusBadRequest, models.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "FAIL",
			Data:   err.Error(),
		})
	}

	ctx, cancel := context.WithTimeout(app.Request().Context(), 10*time.Second)
	defer cancel()

	if queryParam.ID != uuid.Nil {
		response, err := controller.UserService.GetCompleteByID(ctx, queryParam.ID)
		if err != nil {
			if err == context.DeadlineExceeded {
				return app.JSON(http.StatusRequestTimeout, models.WebResponse{
					Code:   http.StatusRequestTimeout,
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

	if queryParam.Username != "" {
		response, err := controller.UserService.GetByUsername(ctx, queryParam.Username)
		if err != nil {
			if err != context.DeadlineExceeded {
				return app.JSON(http.StatusRequestTimeout, models.WebResponse{
					Code:   http.StatusRequestTimeout,
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

	response, err := controller.UserService.GetAll(ctx)
	if err != nil {
		if err == context.DeadlineExceeded {
			return app.JSON(http.StatusRequestTimeout, models.WebResponse{
				Code:   http.StatusRequestTimeout,
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

func (controller *UserController) Update(app echo.Context) error {
	var request models.UpdateUserRequest
	err := app.Bind(&request)
	if err != nil {
		return app.JSON(http.StatusBadRequest, models.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "FAIL",
			Data:   err.Error(),
		})
	}

	response, err := controller.UserService.Update(app.Request().Context(), request)
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

func (controller *UserController) Delete(app echo.Context) error {
	var urlParam models.GetUserIDRequest
	err := app.Bind(&urlParam)
	if err != nil {
		return app.JSON(http.StatusBadRequest, models.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "FAIL",
			Data:   err.Error(),
		})
	}

	err = controller.UserService.Delete(app.Request().Context(), urlParam.ID)
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
		Data:   nil,
	})
}

func (controller *UserController) Login(app echo.Context) error {
	var request models.LoginUserRequest

	err := app.Bind(&request)
	if err != nil {
		return app.JSON(http.StatusBadRequest, models.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "FAIL",
			Data:   err.Error(),
		})
	}

	ctx, cancel := context.WithTimeout(app.Request().Context(), 20*time.Second)
	defer cancel()

	tokenResponse, err := controller.UserService.Login(ctx, request)
	if err != nil {
		return app.JSON(http.StatusUnauthorized, models.WebResponse{
			Code:   http.StatusUnauthorized,
			Status: "FAIL",
			Data:   err.Error(),
		})
	}

	return app.JSON(http.StatusOK, models.WebResponse{
		Code:   http.StatusOK,
		Status: "SUCCESS",
		Data:   tokenResponse,
	})
}
