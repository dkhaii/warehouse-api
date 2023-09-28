package user

import (
	"time"

	"github.com/dkhaii/warehouse-api/domain/user"
	"github.com/dkhaii/warehouse-api/domain/user/repository"
	"github.com/dkhaii/warehouse-api/model"
	"github.com/google/uuid"
)

type UserServiceImpl struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) *UserServiceImpl {
	return &UserServiceImpl{
		userRepository: userRepository,
	}
}

func (service *UserServiceImpl) Create(request model.CreateUserRequest) (model.CreateUserResponse, error) {
	createdAt := time.Now()

	user := user.UserEntity{
		ID:        request.ID,
		Name:      request.Name,
		Contact:   request.Contact,
		Role:      request.Role,
		CreatedAt: createdAt,
		UpdatedAt: createdAt,
	}

	_, err := service.userRepository.Insert(&user)
	if err != nil {
		return model.CreateUserResponse{}, err
	}

	response := model.CreateUserResponse{
		ID:        user.ID,
		Name:      user.Name,
		Contact:   user.Contact,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	return response, nil
}

func (service *UserServiceImpl) GetByID(usrID uuid.UUID) (model.GetUserResponse, error) {
	user, err := service.userRepository.FindByID(usrID)
	if err != nil {
		return model.GetUserResponse{}, err
	}

	return model.GetUserResponse{
		ID:      user.ID,
		Name:    user.Name,
		Contact: user.Contact,
		Role:    user.Role,
	}, nil
}

func (service *UserServiceImpl) GetByName(name string) ([]model.GetUserResponse, error) {
	users, err := service.userRepository.FindByName(name)
	if err != nil {
		return nil, err
	}

	responses := make([]model.GetUserResponse, len(users))

	for key, user := range users {
		responses[key] = model.GetUserResponse{
			ID:      user.ID,
			Name:    user.Name,
			Contact: user.Contact,
			Role:    user.Role,
		}
	}

	return responses, nil
}

func (service *UserServiceImpl) Update(request model.CreateUserRequest) error {
	updatedAt := time.Now()

	user := user.UserEntity{
		Name:      request.Name,
		Contact:   request.Contact,
		Role:      request.Role,
		UpdatedAt: updatedAt,
	}

	err := service.userRepository.Update(&user)
	if err != nil {
		return err
	}

	return nil
}

func (service *UserServiceImpl) Delete(usrID uuid.UUID) error {
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
