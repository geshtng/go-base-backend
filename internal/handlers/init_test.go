package handlers

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

	"github.com/geshtng/go-base-backend/internal/dtos"
	smock "github.com/geshtng/go-base-backend/internal/mocks/services"
)

var MockTokenized dtos.JwtData = dtos.JwtData{
	ID:       1,
	Username: "username",
}

func MakeRequestBody(dto interface{}) *strings.Reader {
	payload, _ := json.Marshal(dto)
	return strings.NewReader(string(payload))
}

// To mock endpoint param
func MiddlewareMockID(mockID string) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Params = []gin.Param{
			{
				Key:   "id",
				Value: mockID,
			},
		}
	}
}

// To mock token
func MiddlewareMockToken(ctx *gin.Context) {
	ctx.Set("user", MockTokenized)
	ctx.Next()
}

func TestNew(t *testing.T) {
	articleService := smock.NewArticleService(t)

	handlerConfig := HandlerConfig{
		ArticleService: articleService,
	}

	New(&handlerConfig)
}

func TestInitAllHandlers(t *testing.T) {
	tests := []struct {
		name string
		want *Handler
	}{
		{
			name: "Pass",
			want: &Handler{
				articleService: smock.NewArticleService(t),
				authService:    smock.NewAuthService(t),
				userService:    smock.NewUserService(t),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := InitAllHandlers()
			assert.NotNil(t, got, tt.want)
		})
	}
}
