package services

import (
	"gorm.io/gorm"

	"github.com/geshtng/go-base-backend/internal/helpers"
	"github.com/geshtng/go-base-backend/internal/models"
	r "github.com/geshtng/go-base-backend/internal/repositories"
)

type AuthService interface {
	Register(auth models.User) (*models.User, *string, error)
	Login(username, password string) (*models.User, *string, error)
}

type authService struct {
	userRepository r.UserRepository
}

type AuthServiceConfig struct {
	UserRepository r.UserRepository
}

func NewAuthService(a *AuthServiceConfig) AuthService {
	return &authService{
		userRepository: a.UserRepository,
	}
}

func (a *authService) Register(auth models.User) (*models.User, *string, error) {
	encryptedPassword, err := helpers.Encrypt(auth.Password)
	if err != nil {
		return nil, nil, err
	}

	auth.Password = encryptedPassword

	newAuth, err := a.userRepository.Insert(auth)
	if err != nil {
		return nil, nil, err
	}

	dataToJwt := models.User{
		ID:       newAuth.ID,
		Username: newAuth.Username,
	}

	token, err := helpers.GenerateJwtToken(dataToJwt)
	if err != nil {
		return nil, nil, err
	}

	return newAuth, token, nil
}

func (a *authService) Login(username, password string) (*models.User, *string, error) {
	auth, err := a.userRepository.GetByUsername(username)
	if err != nil {
		return nil, nil, err
	}

	if !helpers.CompareEncrypt(password, auth.Password) {
		return nil, nil, gorm.ErrRecordNotFound
	}

	dataToJwt := models.User{
		ID:       auth.ID,
		Username: auth.Username,
	}

	token, err := helpers.GenerateJwtToken(dataToJwt)
	if err != nil {
		return nil, nil, err
	}

	return auth, token, nil
}
