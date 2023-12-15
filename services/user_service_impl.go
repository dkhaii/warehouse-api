package services

import (
	"context"
	"crypto/subtle"
	"database/sql"
	"time"

	"github.com/dkhaii/warehouse-api/config"
	"github.com/dkhaii/warehouse-api/entity"
	"github.com/dkhaii/warehouse-api/helpers"
	"github.com/dkhaii/warehouse-api/models"
	"github.com/dkhaii/warehouse-api/repositories"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type userServiceImpl struct {
	userRepository repositories.UserRepository
	database       *sql.DB
}

func NewUserService(userRepository repositories.UserRepository, database *sql.DB) UserService {
	return &userServiceImpl{
		userRepository: userRepository,
		database:       database,
	}
}

func (service *userServiceImpl) Create(ctx context.Context, request models.CreateUserRequest) (models.CreateUserResponse, error) {
	err := helpers.ValidateRequest(request)
	if err != nil {
		return models.CreateUserResponse{}, err
	}

	tx, err := service.database.Begin()
	if err != nil {
		return models.CreateUserResponse{}, err
	}
	defer helpers.CommitOrRollBack(tx)

	userID := uuid.New()
	createdAt := time.Now()

	request.ID = userID
	request.CreatedAt = createdAt
	request.UpdatedAt = request.CreatedAt

	hashedPassword, err := helpers.HashPassword(request.Password)
	if err != nil {
		return models.CreateUserResponse{}, err
	}
	request.Password = hashedPassword

	user := entity.User{
		ID:        request.ID,
		Username:  request.Username,
		Password:  request.Password,
		Contact:   request.Contact,
		RoleID:    request.RoleID,
		CreatedAt: request.CreatedAt,
		UpdatedAt: request.UpdatedAt,
		Role:      nil,
	}

	createdUser, err := service.userRepository.Insert(ctx, tx, &user)
	if err != nil {
		return models.CreateUserResponse{}, err
	}

	response := models.CreateUserResponse{
		ID:        createdUser.ID,
		Username:  createdUser.Username,
		Contact:   createdUser.Contact,
		RoleID:    createdUser.RoleID,
		CreatedAt: createdUser.CreatedAt,
		UpdatedAt: createdUser.UpdatedAt,
	}

	return response, nil
}

func (service *userServiceImpl) GetAll(ctx context.Context) ([]models.GetUserResponse, error) {
	rows, err := service.userRepository.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	users := make([]models.GetUserResponse, len(rows))

	for key, user := range rows {
		users[key] = models.GetUserResponse{
			ID:        user.ID,
			Username:  user.Username,
			Contact:   user.Contact,
			RoleID:    user.RoleID,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		}
	}

	return users, nil
}

func (service *userServiceImpl) GetCompleteByID(ctx context.Context, usrID uuid.UUID) (models.GetCompleteUserResponse, error) {
	user, err := service.userRepository.FindCompleteByID(ctx, usrID)
	if err != nil {
		return models.GetCompleteUserResponse{}, err
	}

	response := models.GetCompleteUserResponse{
		ID:        user.ID,
		Username:  user.Username,
		Contact:   user.Contact,
		RoleID:    user.RoleID,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Role:      user.Role,
	}

	return response, nil
}

func (service *userServiceImpl) GetByUsername(ctx context.Context, name string) ([]models.GetUserResponse, error) {
	rows, err := service.userRepository.GetByUsername(ctx, name)
	if err != nil {
		return nil, err
	}

	users := make([]models.GetUserResponse, len(rows))

	for key, user := range rows {
		users[key] = models.GetUserResponse{
			ID:       user.ID,
			Username: user.Username,
			Contact:  user.Contact,
			RoleID:   user.RoleID,
		}
	}

	return users, nil
}

func (service *userServiceImpl) Update(ctx context.Context, request models.UpdateUserRequest) (models.CreateUserResponse, error) {
	err := helpers.ValidateRequest(request)
	if err != nil {
		return models.CreateUserResponse{}, err
	}

	user, err := service.userRepository.FindByID(ctx, request.ID)
	if err != nil {
		return models.CreateUserResponse{}, err
	}

	tx, err := service.database.Begin()
	if err != nil {
		return models.CreateUserResponse{}, err
	}
	defer helpers.CommitOrRollBack(tx)

	request.UpdatedAt = time.Now()

	updatedUser := entity.User{
		ID:        user.ID,
		Username:  request.Username,
		Password:  request.Password,
		Contact:   request.Contact,
		RoleID:    request.RoleID,
		CreatedAt: user.CreatedAt,
		UpdatedAt: request.UpdatedAt,
	}

	userData, err := service.userRepository.Update(ctx, tx, &updatedUser)
	if err != nil {
		return models.CreateUserResponse{}, err
	}

	response := models.CreateUserResponse{
		ID:        userData.ID,
		Username:  userData.Username,
		Contact:   userData.Contact,
		RoleID:    userData.RoleID,
		CreatedAt: userData.CreatedAt,
		UpdatedAt: userData.UpdatedAt,
	}

	return response, nil
}

func (service *userServiceImpl) Delete(ctx context.Context, usrID uuid.UUID) error {
	user, err := service.userRepository.FindByID(ctx, usrID)
	if err != nil {
		return err
	}

	tx, err := service.database.Begin()
	if err != nil {
		return err
	}
	defer helpers.CommitOrRollBack(tx)

	err = service.userRepository.Delete(ctx, tx, user.ID)
	if err != nil {
		return err
	}

	return nil
}

func (service *userServiceImpl) Login(ctx context.Context, request models.LoginUserRequest) (models.LoginUserResponse, error) {
	err := helpers.ValidateRequest(request)
	if err != nil {
		return models.LoginUserResponse{}, err
	}

	user, err := service.userRepository.FindByUsername(ctx, request.Username)
	if err != nil {
		return models.LoginUserResponse{}, err
	}

	if subtle.ConstantTimeCompare([]byte(request.Username), []byte(user.Username)) == 1 && helpers.ComparePassword(user.Password, request.Password) {
		config, err := config.Init()
		if err != nil {
			return models.LoginUserResponse{}, err
		}

		jwtSecret := config.GetString("JWT_SECRET")

		token, err := helpers.CreateAccessToken(user, jwtSecret, 2)
		if err != nil {
			return models.LoginUserResponse{}, err
		}

		userClaim, err := helpers.GetUserClaimsFromToken(token, jwtSecret)
		if err != nil {
			return models.LoginUserResponse{}, err
		}

		return models.LoginUserResponse{
			Username: userClaim.Username,
			RoleID:   userClaim.RoleID,
			Token:    token,
		}, nil
	}

	return models.LoginUserResponse{}, jwt.ErrTokenMalformed
}
