package services

import (
	"github.com/geshtng/go-base-backend/internal/models"
	r "github.com/geshtng/go-base-backend/internal/repositories"
)

type UserService interface {
	GetByID(id int64) (*models.User, error)
}

type userService struct {
	userRepository r.UserRepository
}

type UserServiceConfig struct {
	UserRepository r.UserRepository
}

func NewUserService(u *UserServiceConfig) UserService {
	return &userService{
		userRepository: u.UserRepository,
	}
}

func (u *userService) GetByID(id int64) (*models.User, error) {
	result, err := u.userRepository.GetByID(id)
	if err != nil {
		return nil, err
	}

	return result, nil
}
