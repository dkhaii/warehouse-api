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

type userControllerImpl struct {
	userService services.UserService
}

func NewUserController(userService services.UserService) UserController {
	return &userControllerImpl{
		userService: userService,
	}
}

func (controller *userControllerImpl) Create(app echo.Context) error {
	var request models.CreateUserRequest
	err := app.Bind(&request)
	if err != nil {
		return helpers.CreateResponseError(app, http.StatusInternalServerError, err)
	}

	response, err := controller.userService.Create(app.Request().Context(), request)
	if err != nil {
		return helpers.CreateResponseError(app, http.StatusBadRequest, err)
	}
	return helpers.CreateResponse(app, http.StatusCreated, response)
}

func (controller *userControllerImpl) GetUser(app echo.Context) error {
	var queryParam models.GetUserRequest
	err := app.Bind(&queryParam)
	if err != nil {
		return helpers.CreateResponseError(app, http.StatusBadRequest, err)
	}

	ctx, cancel := context.WithTimeout(app.Request().Context(), 30*time.Second)
	defer cancel()

	if queryParam.ID != uuid.Nil {
		response, err := controller.userService.GetCompleteByID(ctx, queryParam.ID)
		if err != nil {
			if err == context.DeadlineExceeded {
				return helpers.CreateResponseError(app, http.StatusRequestTimeout, helpers.ErrRequestTimedOut)
			}
			return helpers.CreateResponseError(app, http.StatusNotFound, err)
		}
		return helpers.CreateResponse(app, http.StatusFound, response)
	}

	if queryParam.Username != "" {
		response, err := controller.userService.GetByUsername(ctx, queryParam.Username)
		if err != nil {
			if err != context.DeadlineExceeded {
				return helpers.CreateResponseError(app, http.StatusRequestTimeout, helpers.ErrRequestTimedOut)
			}
			return helpers.CreateResponseError(app, http.StatusNotFound, err)
		}
		return helpers.CreateResponse(app, http.StatusFound, response)
	}

	response, err := controller.userService.GetAll(ctx)
	if err != nil {
		if err == context.DeadlineExceeded {
			return helpers.CreateResponseError(app, http.StatusRequestTimeout, helpers.ErrRequestTimedOut)
		}
		return helpers.CreateResponseError(app, http.StatusNotFound, err)
	}
	return helpers.CreateResponse(app, http.StatusFound, response)
}

func (controller *userControllerImpl) Update(app echo.Context) error {
	var request models.UpdateUserRequest
	err := app.Bind(&request)
	if err != nil {
		return helpers.CreateResponseError(app, http.StatusBadRequest, err)
	}

	response, err := controller.userService.Update(app.Request().Context(), request)
	if err != nil {
		return helpers.CreateResponseError(app, http.StatusNotFound, err)
	}
	return helpers.CreateResponse(app, http.StatusOK, response)
}

func (controller *userControllerImpl) Delete(app echo.Context) error {
	var urlParam models.GetUserIDRequest
	err := app.Bind(&urlParam)
	if err != nil {
		return helpers.CreateResponseError(app, http.StatusBadRequest, err)
	}

	err = controller.userService.Delete(app.Request().Context(), urlParam.ID)
	if err != nil {
		return helpers.CreateResponseError(app, http.StatusNotFound, err)
	}
	return helpers.CreateResponse(app, http.StatusOK, nil)
}

func (controller *userControllerImpl) Login(app echo.Context) error {
	var request models.LoginUserRequest

	err := app.Bind(&request)
	if err != nil {
		return helpers.CreateResponse(app, http.StatusBadRequest, err)
	}

	ctx, cancel := context.WithTimeout(app.Request().Context(), 20*time.Second)
	defer cancel()

	tokenResponse, err := controller.userService.Login(ctx, request)
	if err != nil {
		return helpers.CreateResponseError(app, http.StatusUnauthorized, err)
	}

	return helpers.CreateResponse(app, http.StatusOK, tokenResponse)
}
