package services

import (
	"crypto/subtle"
	// "fmt"

	"github.com/dkhaii/warehouse-api/config"
	"github.com/dkhaii/warehouse-api/entity"
	"github.com/dkhaii/warehouse-api/internal/tokenutil"
	"github.com/dkhaii/warehouse-api/models"
	"github.com/dkhaii/warehouse-api/repositories"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type userServiceImpl struct {
	userRepository repositories.UserRepository
}

func NewUserService(userRepository repositories.UserRepository) UserService {
	return &userServiceImpl{
		userRepository: userRepository,
	}
}

func (service *userServiceImpl) Create(request models.CreateUserRequest) (models.CreateUserResponse, error) {
	user := entity.User{
		ID:        request.ID,
		Username:  request.Username,
		Password:  request.Password,
		Contact:   request.Contact,
		Role:      request.Role,
		CreatedAt: request.CreatedAt,
		UpdatedAt: request.UpdatedAt,
	}

	_, err := service.userRepository.Insert(&user)
	if err != nil {
		return models.CreateUserResponse{}, err
	}

	response := models.CreateUserResponse{
		ID:        user.ID,
		Username:  user.Username,
		Contact:   user.Contact,
		Role:      user.Role,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	return response, nil
}

func (service *userServiceImpl) GetAll() ([]models.GetUserResponse, error) {
	users, err := service.userRepository.FindAll()
	if err != nil {
		return nil, err
	}

	responses := make([]models.GetUserResponse, len(users))

	for key, user := range users {
		responses[key] = models.GetUserResponse{
			ID:        user.ID,
			Username:  user.Username,
			Contact:   user.Contact,
			Role:      user.Role,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		}
	}

	return responses, nil
}

func (service *userServiceImpl) GetByID(usrID uuid.UUID) (models.GetUserResponse, error) {
	user, err := service.userRepository.FindByID(usrID)
	if err != nil {
		return models.GetUserResponse{}, err
	}

	response := models.GetUserResponse{
		ID:        user.ID,
		Username:  user.Username,
		Contact:   user.Contact,
		Role:      user.Role,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	return response, nil
}

func (service *userServiceImpl) GetByUsername(name string) ([]models.GetUserResponse, error) {
	users, err := service.userRepository.GetByUsername(name)
	if err != nil {
		return nil, err
	}

	responses := make([]models.GetUserResponse, len(users))

	for key, user := range users {
		responses[key] = models.GetUserResponse{
			ID:       user.ID,
			Username: user.Username,
			Contact:  user.Contact,
			Role:     user.Role,
		}
	}

	return responses, nil
}

func (service *userServiceImpl) Update(request models.UpdateUserRequest) error {
	isUser, err := service.userRepository.FindByID(request.ID)
	if err != nil {
		return err
	}

	updatedUser := entity.User{
		ID:        isUser.ID,
		Username:  request.Username,
		Password:  request.Password,
		Contact:   request.Contact,
		Role:      request.Role,
		CreatedAt: isUser.CreatedAt,
		UpdatedAt: request.UpdatedAt,
	}

	err = service.userRepository.Update(&updatedUser)
	if err != nil {
		return err
	}

	return nil
}

func (service *userServiceImpl) Delete(usrID uuid.UUID) error {
	user, err := service.userRepository.FindByID(usrID)
	if err != nil {
		return err
	}

	err = service.userRepository.Delete(user.ID)
	if err != nil {
		return err
	}

	return nil
}

func (service *userServiceImpl) Login(request models.LoginUserRequest) (models.TokenResponse, error) {
	user, err := service.userRepository.FindByUsername(request.Username)
	if err != nil {
		return models.TokenResponse{}, err
	}

	if subtle.ConstantTimeCompare([]byte(request.Username), []byte(user.Username)) == 1 && subtle.ConstantTimeCompare([]byte(request.Password), []byte(user.Password)) == 1 {
		config, err := config.New()
		if err != nil {
			return models.TokenResponse{}, err
		}

		jwtSecret := config.Get("JWT_SECRET")

		token, err := tokenutil.CreateAccessToken(user, jwtSecret, 24)
		if err != nil {
			return models.TokenResponse{}, err
		}

		return models.TokenResponse{
			Token: token,
		}, nil
	}

	return models.TokenResponse{}, jwt.ErrTokenMalformed
}
