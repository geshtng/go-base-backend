package services

import (
	"errors"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	rmock "github.com/geshtng/go-base-backend/internal/mocks/repositories"
	"github.com/geshtng/go-base-backend/internal/models"
)

func TestNewUserService(t *testing.T) {
	userRepo := rmock.NewUserRepository(t)
	userService := NewUserService(&UserServiceConfig{
		UserRepository: userRepo,
	})

	assert.NotNil(t, userService)
}

func Test_userService_GetByID(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		type args struct {
			id int64
		}
		tt := struct {
			args    args
			want    *models.User
			wantErr bool
		}{
			args: args{
				id: 1,
			},
			want: &models.User{
				ID:       1,
				Username: "username",
			},
			wantErr: false,
		}

		userRepo := rmock.NewUserRepository(t)
		s := &userService{
			userRepository: userRepo,
		}

		result := &models.User{
			ID:       1,
			Username: "username",
		}

		userRepo.On("GetByID", mock.Anything).Return(result, nil)

		got, err := s.GetByID(tt.args.id)
		if (err != nil) != tt.wantErr {
			t.Errorf("articleService.GetAll() error = %v, wantErr %v", err, tt.wantErr)
			return
		}

		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("articleService.GetAll() = %v, want %v", got, tt.want)
		}
	})

	t.Run("Error Repository.GetByID", func(t *testing.T) {
		type args struct {
			id int64
		}
		tt := struct {
			args    args
			want    *models.User
			wantErr bool
		}{
			args: args{
				id: 1,
			},
			want:    nil,
			wantErr: true,
		}

		userRepo := rmock.NewUserRepository(t)
		s := &userService{
			userRepository: userRepo,
		}

		userRepo.On("GetByID", mock.Anything).Return(nil, errors.New("something error"))

		got, err := s.GetByID(tt.args.id)
		if (err != nil) != tt.wantErr {
			t.Errorf("articleService.GetAll() error = %v, wantErr %v", err, tt.wantErr)
			return
		}

		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("articleService.GetAll() = %v, want %v", got, tt.want)
		}
	})
}
