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
		return helpers.CreateResponseError(app, http.StatusInternalServerError, err)
	}

	response, err := controller.UserService.Create(app.Request().Context(), request)
	if err != nil {
		return helpers.CreateResponseError(app, http.StatusBadRequest, err)
	}
	return helpers.CreateResponse(app, http.StatusCreated, response)
}

func (controller *UserController) GetUser(app echo.Context) error {
	var queryParam models.GetUserRequest
	err := app.Bind(&queryParam)
	if err != nil {
		return helpers.CreateResponseError(app, http.StatusBadRequest, err)
	}

	ctx, cancel := context.WithTimeout(app.Request().Context(), 30*time.Second)
	defer cancel()

	if queryParam.ID != uuid.Nil {
		response, err := controller.UserService.GetCompleteByID(ctx, queryParam.ID)
		if err != nil {
			if err == context.DeadlineExceeded {
				return helpers.CreateResponseError(app, http.StatusRequestTimeout, helpers.ErrRequestTimedOut)
			}
			return helpers.CreateResponseError(app, http.StatusNotFound, err)
		}
		return helpers.CreateResponse(app, http.StatusFound, response)
	}

	if queryParam.Username != "" {
		response, err := controller.UserService.GetByUsername(ctx, queryParam.Username)
		if err != nil {
			if err != context.DeadlineExceeded {
				return helpers.CreateResponseError(app, http.StatusRequestTimeout, helpers.ErrRequestTimedOut)
			}
			return helpers.CreateResponseError(app, http.StatusNotFound, err)
		}
		return helpers.CreateResponse(app, http.StatusFound, response)
	}

	response, err := controller.UserService.GetAll(ctx)
	if err != nil {
		if err == context.DeadlineExceeded {
			return helpers.CreateResponseError(app, http.StatusRequestTimeout, helpers.ErrRequestTimedOut)
		}
		return helpers.CreateResponseError(app, http.StatusNotFound, err)
	}
	return helpers.CreateResponse(app, http.StatusFound, response)
}

func (controller *UserController) Update(app echo.Context) error {
	var request models.UpdateUserRequest
	err := app.Bind(&request)
	if err != nil {
		return helpers.CreateResponseError(app, http.StatusBadRequest, err)
	}

	response, err := controller.UserService.Update(app.Request().Context(), request)
	if err != nil {
		return helpers.CreateResponseError(app, http.StatusNotFound, err)
	}
	return helpers.CreateResponse(app, http.StatusOK, response)
}

func (controller *UserController) Delete(app echo.Context) error {
	var urlParam models.GetUserIDRequest
	err := app.Bind(&urlParam)
	if err != nil {
		return helpers.CreateResponseError(app, http.StatusBadRequest, err)
	}

	err = controller.UserService.Delete(app.Request().Context(), urlParam.ID)
	if err != nil {
		return helpers.CreateResponseError(app, http.StatusNotFound, err)
	}
	return helpers.CreateResponse(app, http.StatusOK, nil)
}

func (controller *UserController) Login(app echo.Context) error {
	var request models.LoginUserRequest

	err := app.Bind(&request)
	if err != nil {
		return helpers.CreateResponse(app, http.StatusBadRequest, err)
	}

	ctx, cancel := context.WithTimeout(app.Request().Context(), 20*time.Second)
	defer cancel()

	tokenResponse, err := controller.UserService.Login(ctx, request)
	if err != nil {
		return helpers.CreateResponseError(app, http.StatusUnauthorized, err)
	}

	// cookie := helpers.CreateCookie(tokenResponse.Token, 24)
	// app.SetCookie(cookie)

	return helpers.CreateResponse(app, http.StatusOK, tokenResponse)
}
