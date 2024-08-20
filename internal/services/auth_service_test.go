package services

import (
	"errors"
	"os"
	"path"
	"reflect"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/geshtng/go-base-backend/internal/helpers"
	rmock "github.com/geshtng/go-base-backend/internal/mocks/repositories"
	"github.com/geshtng/go-base-backend/internal/models"
)

func TestNewAuthService(t *testing.T) {
	userRepo := rmock.NewUserRepository(t)
	authService := NewAuthService(&AuthServiceConfig{
		UserRepository: userRepo,
	})

	assert.NotNil(t, authService)
}

func Test_authService_Register(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		_, filename, _, _ := runtime.Caller(0)
		dir := path.Join(path.Dir(filename), "../../")
		err := os.Chdir(dir)
		if err != nil {
			panic(err)
		}

		encryptedPassword, _ := helpers.Encrypt("password")

		type args struct {
			auth models.User
		}
		test := struct {
			args    args
			want    *models.User
			wantErr bool
		}{
			args: args{
				auth: models.User{
					Username: "username",
					Password: "password",
				},
			},
			want: &models.User{
				Username: "username",
				Password: encryptedPassword,
			},
			wantErr: false,
		}

		userRepo := rmock.NewUserRepository(t)
		s := &authService{
			userRepository: userRepo,
		}

		expectedUser := models.User{
			Username: "username",
			Password: encryptedPassword,
		}

		userRepo.On("Insert", mock.Anything).Return(&expectedUser, nil)

		got, _, err := s.Register(test.args.auth)
		if (err != nil) != test.wantErr {
			t.Errorf("authService.Register() error = %v, wantErr %v", err, test.wantErr)
			return
		}

		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("authService.Register() = %v, want %v", got, test.want)
		}
	})

	t.Run("Error Repository.Insert", func(t *testing.T) {
		_, filename, _, _ := runtime.Caller(0)
		dir := path.Join(path.Dir(filename), "../../")
		err := os.Chdir(dir)
		if err != nil {
			panic(err)
		}

		type args struct {
			auth models.User
		}
		test := struct {
			args    args
			want    *models.User
			wantErr bool
		}{
			args: args{
				auth: models.User{
					Username: "username",
					Password: "password",
				},
			},
			want:    nil,
			wantErr: true,
		}

		userRepo := rmock.NewUserRepository(t)
		s := &authService{
			userRepository: userRepo,
		}

		userRepo.On("Insert", mock.Anything).Return(nil, errors.New("something error"))

		got, _, err := s.Register(test.args.auth)
		if (err != nil) != test.wantErr {
			t.Errorf("authService.Register() error = %v, wantErr %v", err, test.wantErr)
			return
		}

		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("authService.Register() = %v, want %v", got, test.want)
		}
	})
}

func Test_authService_Login(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		_, filename, _, _ := runtime.Caller(0)
		dir := path.Join(path.Dir(filename), "../../")
		err := os.Chdir(dir)
		if err != nil {
			panic(err)
		}

		encryptedPassword, _ := helpers.Encrypt("password")

		type args struct {
			username string
			password string
		}
		test := struct {
			args    args
			want    *models.User
			wantErr bool
		}{
			args: args{
				username: "username",
				password: "password",
			},
			want: &models.User{
				Username: "username",
				Password: encryptedPassword,
			},
			wantErr: false,
		}

		userRepo := rmock.NewUserRepository(t)
		s := &authService{
			userRepository: userRepo,
		}

		expectedUser := models.User{
			Username: "username",
			Password: encryptedPassword,
		}

		userRepo.On("GetByUsername", mock.Anything).Return(&expectedUser, nil)

		got, _, err := s.Login(test.args.username, test.args.password)
		if (err != nil) != test.wantErr {
			t.Errorf("authService.Register() error = %v, wantErr %v", err, test.wantErr)
			return
		}

		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("authService.Register() = %v, want %v", got, test.want)
		}
	})

	t.Run("Error Repository.GetByUsername", func(t *testing.T) {
		_, filename, _, _ := runtime.Caller(0)
		dir := path.Join(path.Dir(filename), "../../")
		err := os.Chdir(dir)
		if err != nil {
			panic(err)
		}

		type args struct {
			username string
			password string
		}
		test := struct {
			args    args
			want    *models.User
			wantErr bool
		}{
			args: args{
				username: "username",
				password: "password",
			},
			want:    nil,
			wantErr: true,
		}

		userRepo := rmock.NewUserRepository(t)
		s := &authService{
			userRepository: userRepo,
		}

		userRepo.On("GetByUsername", mock.Anything).Return(nil, errors.New("something error"))

		got, _, err := s.Login(test.args.username, test.args.password)
		if (err != nil) != test.wantErr {
			t.Errorf("authService.Register() error = %v, wantErr %v", err, test.wantErr)
			return
		}

		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("authService.Register() = %v, want %v", got, test.want)
		}
	})

	t.Run("Error CompareEncrypt", func(t *testing.T) {
		_, filename, _, _ := runtime.Caller(0)
		dir := path.Join(path.Dir(filename), "../../")
		err := os.Chdir(dir)
		if err != nil {
			panic(err)
		}

		encryptedPassword, _ := helpers.Encrypt("password")

		type args struct {
			username string
			password string
		}
		test := struct {
			args    args
			want    *models.User
			wantErr bool
		}{
			args: args{
				username: "username",
				password: "wrongpassword",
			},
			want:    nil,
			wantErr: true,
		}

		userRepo := rmock.NewUserRepository(t)
		s := &authService{
			userRepository: userRepo,
		}

		expectedUser := models.User{
			Username: "username",
			Password: encryptedPassword,
		}

		userRepo.On("GetByUsername", mock.Anything).Return(&expectedUser, nil)

		got, _, err := s.Login(test.args.username, test.args.password)
		if (err != nil) != test.wantErr {
			t.Errorf("authService.Register() error = %v, wantErr %v", err, test.wantErr)
			return
		}

		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("authService.Register() = %v, want %v", got, test.want)
		}
	})
}
