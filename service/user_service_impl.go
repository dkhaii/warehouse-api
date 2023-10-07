package service

import (
	"github.com/dkhaii/warehouse-api/entity"
	"github.com/dkhaii/warehouse-api/model"
	"github.com/dkhaii/warehouse-api/repository"
	"github.com/google/uuid"
)

type userServiceImpl struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) UserService {
	return &userServiceImpl{
		userRepository: userRepository,
	}
}

func (service *userServiceImpl) Create(request model.CreateUserRequest) (model.CreateUserResponse, error) {
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
		return model.CreateUserResponse{}, err
	}

	response := model.CreateUserResponse{
		ID:        user.ID,
		Username:  user.Username,
		Contact:   user.Contact,
		Role:      user.Role,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	return response, nil
}

func (service *userServiceImpl) GetAll() ([]model.GetUserResponse, error) {
	users, err := service.userRepository.FindAll()
	if err != nil {
		return nil, err
	}

	responses := make([]model.GetUserResponse, len(users))

	for key, user := range users {
		responses[key] = model.GetUserResponse{
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

func (service *userServiceImpl) GetByID(usrID uuid.UUID) (model.GetUserResponse, error) {
	user, err := service.userRepository.FindByID(usrID)
	if err != nil {
		return model.GetUserResponse{}, err
	}

	response := model.GetUserResponse{
		ID:        user.ID,
		Username:  user.Username,
		Contact:   user.Contact,
		Role:      user.Role,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	return response, nil
}

func (service *userServiceImpl) GetByUsername(name string) ([]model.GetUserResponse, error) {
	users, err := service.userRepository.FindByUsername(name)
	if err != nil {
		return nil, err
	}

	responses := make([]model.GetUserResponse, len(users))

	for key, user := range users {
		responses[key] = model.GetUserResponse{
			ID:       user.ID,
			Username: user.Username,
			Contact:  user.Contact,
			Role:     user.Role,
		}
	}

	return responses, nil
}

func (service *userServiceImpl) Update(request model.CreateUserRequest) error {
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

// func (service *UserServiceImpl) Login(request model.LoginUserRequest) (model.LoginUserResponse, error) {
// }
