package user

import (
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

func (service *UserServiceImpl) GetByID(id uuid.UUID) (response model.GetUserResponse, error) {
	user, err := uc.userRepository.FindByID(id)
	if err != nil {
		return nil, err
	}

	return response = model.GetUserResponse{
		ID: user.,
	}, nil
}

func (uc *UseCase) GetByName(name string) (*entity.User, error) {
	user, err := uc.userRepository.FindByName(name)
	if err != nil {
		return nil, err
	}

	return user, nil
}
