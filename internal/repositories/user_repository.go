package repositories

import (
	"gorm.io/gorm"

	"github.com/geshtng/go-base-backend/internal/models"
)

type UserRepository interface {
	GetByID(id int64) (*models.User, error)
	GetByUsername(username string) (*models.User, error)
	Insert(user models.User) (*models.User, error)
}

type userRepository struct {
	db *gorm.DB
}

type UserRepoConfig struct {
	DB *gorm.DB
}

func NewUserRepository(u *UserRepoConfig) UserRepository {
	return &userRepository{
		db: u.DB,
	}
}

func (u *userRepository) GetByID(id int64) (*models.User, error) {
	var user *models.User

	result := u.db.
		Where("id = ?", id).
		First(&user)

	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}

func (u *userRepository) GetByUsername(username string) (*models.User, error) {
	var user *models.User

	result := u.db.
		Where("username = ?", username).
		First(&user)

	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}

func (u *userRepository) Insert(user models.User) (*models.User, error) {
	result := u.db.Create(&user)
	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}
