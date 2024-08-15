package handlers

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgconn"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"

	"github.com/geshtng/go-base-backend/internal/dtos"
	"github.com/geshtng/go-base-backend/internal/helpers"
	smock "github.com/geshtng/go-base-backend/internal/mocks/services"
	"github.com/geshtng/go-base-backend/internal/models"
)

func TestHandler_RegisterHandler(t *testing.T) {
	sampleToken := "sample"

	tests := []struct {
		name        string
		body        io.Reader
		authService *smock.AuthService
		mock        func(*smock.AuthService)
		want        helpers.Response
	}{
		{
			name: "Success",
			body: MakeRequestBody(dtos.RegisterRequestDTO{
				Username: "username",
				Password: "password",
			}),
			authService: smock.NewAuthService(t),
			mock: func(as *smock.AuthService) {
				expectedUser := models.User{
					ID:       1,
					Username: "username",
					Password: "password",
				}

				as.On("Register", mock.Anything).Return(&expectedUser, &sampleToken, nil)
			},
			want: helpers.Response{
				Data: map[string]interface{}{
					"id":       float64(1),
					"username": "username",
					"token":    "sample",
				},
				Code:    201,
				Message: "Created",
				Error:   false,
			},
		},
		{
			name: "Error ShouldBindJSON",
			body: MakeRequestBody(dtos.RegisterRequestDTO{
				Username: "username",
			}),
			authService: smock.NewAuthService(t),
			mock:        func(as *smock.AuthService) {},
			want: helpers.Response{
				Code:    400,
				Message: "Key: 'RegisterRequestDTO.Password' Error:Field validation for 'Password' failed on the 'required' tag",
				Error:   true,
			},
		},
		{
			name: "Error Service.Register unique violation",
			body: MakeRequestBody(dtos.RegisterRequestDTO{
				Username: "username",
				Password: "password",
			}),
			authService: smock.NewAuthService(t),
			mock: func(as *smock.AuthService) {
				as.On("Register", mock.Anything).Return(nil, nil, &pgconn.PgError{Code: "23505"})
			},
			want: helpers.Response{
				Code:    401,
				Message: "username is already exist",
				Error:   true,
			},
		},
		{
			name: "Error Service.Register",
			body: MakeRequestBody(dtos.RegisterRequestDTO{
				Username: "username",
				Password: "password",
			}),
			authService: smock.NewAuthService(t),
			mock: func(as *smock.AuthService) {
				as.On("Register", mock.Anything).Return(nil, nil, &pgconn.PgError{Message: "something error"})
			},
			want: helpers.Response{
				Code:    500,
				Message: ": something error (SQLSTATE )",
				Error:   true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &Handler{
				authService: tt.authService,
			}

			tt.mock(tt.authService)

			r := gin.Default()
			endpoint := "/register"

			r.POST(endpoint, h.RegisterHandler)
			req, _ := http.NewRequest(http.MethodPost, endpoint, tt.body)

			w := httptest.NewRecorder()
			r.ServeHTTP(w, req)

			var response helpers.Response
			err := json.Unmarshal(w.Body.Bytes(), &response)
			require.NoError(t, err)

			assert.Equal(t, tt.want.Code, w.Code)
			assert.Equal(t, tt.want, response)
		})
	}
}

func TestHandler_LoginHandler(t *testing.T) {
	sampleToken := "sample"

	tests := []struct {
		name        string
		body        io.Reader
		authService *smock.AuthService
		mock        func(*smock.AuthService)
		want        helpers.Response
	}{
		{
			name: "Success",
			body: MakeRequestBody(dtos.LoginRequestDTO{
				Username: "username",
				Password: "password",
			}),
			authService: smock.NewAuthService(t),
			mock: func(as *smock.AuthService) {
				expectedUser := models.User{
					ID:       1,
					Username: "username",
					Password: "password",
				}

				as.On("Login", mock.Anything, mock.Anything).Return(&expectedUser, &sampleToken, nil)
			},
			want: helpers.Response{
				Data: map[string]interface{}{
					"id":       float64(1),
					"username": "username",
					"token":    "sample",
				},
				Code:    200,
				Message: "OK",
				Error:   false,
			},
		},
		{
			name: "Error ShouldBindJSON",
			body: MakeRequestBody(dtos.LoginRequestDTO{
				Username: "username",
			}),
			authService: smock.NewAuthService(t),
			mock:        func(as *smock.AuthService) {},
			want: helpers.Response{
				Code:    400,
				Message: "Key: 'LoginRequestDTO.Password' Error:Field validation for 'Password' failed on the 'required' tag",
				Error:   true,
			},
		},
		{
			name: "Error Service.Login record not found",
			body: MakeRequestBody(dtos.LoginRequestDTO{
				Username: "username",
				Password: "password",
			}),
			authService: smock.NewAuthService(t),
			mock: func(as *smock.AuthService) {
				as.On("Login", mock.Anything, mock.Anything).Return(nil, nil, gorm.ErrRecordNotFound)
			},
			want: helpers.Response{
				Code:    401,
				Message: "wrong username or password",
				Error:   true,
			},
		},
		{
			name: "Error Service.Login",
			body: MakeRequestBody(dtos.LoginRequestDTO{
				Username: "username",
				Password: "password",
			}),
			authService: smock.NewAuthService(t),
			mock: func(as *smock.AuthService) {
				as.On("Login", mock.Anything, mock.Anything).Return(nil, nil, errors.New("something error"))
			},
			want: helpers.Response{
				Code:    500,
				Message: "something error",
				Error:   true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &Handler{
				authService: tt.authService,
			}

			tt.mock(tt.authService)

			r := gin.Default()
			endpoint := "/login"

			r.POST(endpoint, h.LoginHandler)
			req, _ := http.NewRequest(http.MethodPost, endpoint, tt.body)

			w := httptest.NewRecorder()
			r.ServeHTTP(w, req)

			var response helpers.Response
			err := json.Unmarshal(w.Body.Bytes(), &response)
			require.NoError(t, err)

			assert.Equal(t, tt.want.Code, w.Code)
			assert.Equal(t, tt.want, response)
		})
	}
}
