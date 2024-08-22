package repositories

import (
	"errors"
	"reflect"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/geshtng/go-base-backend/internal/models"
)

func TestNewUserRepository(t *testing.T) {
	dbConn, _, _ := sqlmock.New()
	dia := postgres.New(postgres.Config{
		DriverName: "postgres",
		Conn:       dbConn,
	})

	db, _ := gorm.Open(dia)
	userRepo := NewUserRepository(&UserRepoConfig{DB: db})

	assert.NotNil(t, userRepo)
}

func Test_userRepository_GetByID(t *testing.T) {
	dbConn, mock, _ := sqlmock.New()
	dia := postgres.New(postgres.Config{
		DriverName: "postgres",
		Conn:       dbConn,
	})

	db, _ := gorm.Open(dia)

	type args struct {
		id int64
	}
	tests := []struct {
		name    string
		u       *userRepository
		mock    func()
		args    args
		want    *models.User
		wantErr bool
	}{
		{
			name: "Success",
			u: &userRepository{
				db: db,
			},
			mock: func() {
				mock.ExpectQuery("SELECT (.+)").
					WithArgs(1).
					WillReturnRows(sqlmock.NewRows([]string{"id", "username"}).
						AddRow(1, "username"))
			},
			args: args{
				id: 1,
			},
			want: &models.User{
				ID:       1,
				Username: "username",
			},
			wantErr: false,
		},
		{
			name: "Error DB First",
			u: &userRepository{
				db: db,
			},
			mock: func() {
				mock.ExpectQuery("SELECT (.+)").
					WithArgs(1).
					WillReturnError(errors.New("something error"))
			},
			args: args{
				id: 1,
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()

			got, err := tt.u.GetByID(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("userRepository.GetByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("userRepository.GetByID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_userRepository_GetByUsername(t *testing.T) {
	dbConn, mock, _ := sqlmock.New()
	dia := postgres.New(postgres.Config{
		DriverName: "postgres",
		Conn:       dbConn,
	})

	db, _ := gorm.Open(dia)

	type args struct {
		username string
	}
	tests := []struct {
		name    string
		u       *userRepository
		mock    func()
		args    args
		want    *models.User
		wantErr bool
	}{
		{
			name: "Success",
			u: &userRepository{
				db: db,
			},
			mock: func() {
				mock.ExpectQuery("SELECT (.+)").
					WithArgs("username").
					WillReturnRows(sqlmock.NewRows([]string{"id", "username"}).
						AddRow(1, "username"))
			},
			args: args{
				username: "username",
			},
			want: &models.User{
				ID:       1,
				Username: "username",
			},
			wantErr: false,
		},
		{
			name: "Error DB First",
			u: &userRepository{
				db: db,
			},
			mock: func() {
				mock.ExpectQuery("SELECT (.+)").
					WithArgs("username").
					WillReturnError(errors.New("something error"))
			},
			args: args{
				username: "username",
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()

			got, err := tt.u.GetByUsername(tt.args.username)
			if (err != nil) != tt.wantErr {
				t.Errorf("userRepository.GetByUsername() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("userRepository.GetByUsername() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_userRepository_Insert(t *testing.T) {
	dbConn, mock, _ := sqlmock.New()
	dia := postgres.New(postgres.Config{
		DriverName: "postgres",
		Conn:       dbConn,
	})

	db, _ := gorm.Open(dia)

	type args struct {
		user models.User
	}
	tests := []struct {
		name    string
		u       *userRepository
		mock    func()
		args    args
		want    *models.User
		wantErr bool
	}{
		{
			name: "Success",
			u: &userRepository{
				db: db,
			},
			mock: func() {
				mock.ExpectBegin()
				mock.ExpectQuery(`INSERT INTO "users" (.+) VALUES (.+)`).
					WithArgs("username", "password").
					WillReturnRows(sqlmock.NewRows([]string{"id", "username", "password"}).
						AddRow(1, "username", "password"))
				mock.ExpectCommit()
			},
			args: args{
				user: models.User{
					Username: "username",
					Password: "password",
				},
			},
			want: &models.User{
				ID:       1,
				Username: "username",
				Password: "password",
			},
			wantErr: false,
		},
		{
			name: "Error DB Create",
			u: &userRepository{
				db: db,
			},
			mock: func() {
				mock.ExpectBegin()
				mock.ExpectQuery(`INSERT INTO "users" (.+) VALUES (.+)`).
					WithArgs("username", "password").
					WillReturnError(errors.New("something error"))
				mock.ExpectCommit()
			},
			args: args{
				user: models.User{
					Username: "username",
					Password: "password",
				},
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()

			got, err := tt.u.Insert(tt.args.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("userRepository.Insert() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("userRepository.Insert() = %v, want %v", got, tt.want)
			}
		})
	}
}
