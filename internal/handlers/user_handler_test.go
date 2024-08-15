package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"

	"github.com/geshtng/go-base-backend/internal/helpers"
	smock "github.com/geshtng/go-base-backend/internal/mocks/services"
	"github.com/geshtng/go-base-backend/internal/models"
)

func TestHandler_ShowProfileHandler(t *testing.T) {
	tests := []struct {
		name        string
		userService *smock.UserService
		mock        func(*smock.UserService)
		want        helpers.Response
	}{
		{
			name:        "Success",
			userService: smock.NewUserService(t),
			mock: func(us *smock.UserService) {
				expectedUser := models.User{
					ID:       1,
					Username: "username",
					Password: "password",
				}

				us.On("GetByID", mock.Anything).Return(&expectedUser, nil)
			},

			want: helpers.Response{
				Data: map[string]interface{}{
					"id":       float64(1),
					"username": "username",
				},
				Code:    200,
				Message: "OK",
				Error:   false,
			},
		},
		{
			name:        "Error copier",
			userService: smock.NewUserService(t),
			mock: func(us *smock.UserService) {
				us.On("GetByID", mock.Anything).Return(nil, nil)
			},

			want: helpers.Response{
				Code:    500,
				Message: "copy from is invalid",
				Error:   true,
			},
		},
		{
			name:        "Error Service.GetByID record not found",
			userService: smock.NewUserService(t),
			mock: func(us *smock.UserService) {
				us.On("GetByID", mock.Anything).Return(nil, gorm.ErrRecordNotFound)
			},

			want: helpers.Response{
				Code:    404,
				Message: "record not found",
				Error:   true,
			},
		},
		{
			name:        "Error Service.GetByID",
			userService: smock.NewUserService(t),
			mock: func(us *smock.UserService) {
				us.On("GetByID", mock.Anything).Return(nil, errors.New("something error"))
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
				userService: tt.userService,
			}

			tt.mock((*smock.UserService)(tt.userService))

			r := gin.Default()
			endpoint := "/profiles"

			r.GET(endpoint, MiddlewareMockToken, h.ShowProfileHandler)

			req, _ := http.NewRequest(http.MethodGet, endpoint, nil)

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
