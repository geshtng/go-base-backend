package handlers

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"

	"github.com/geshtng/go-base-backend/internal/dtos"
	"github.com/geshtng/go-base-backend/internal/helpers"
	smock "github.com/geshtng/go-base-backend/internal/mocks/services"
	"github.com/geshtng/go-base-backend/internal/models"
)

func TestHandler_GetAllArticlesHandler(t *testing.T) {
	tests := []struct {
		name           string
		articleService *smock.ArticleService
		mock           func(*smock.ArticleService)
		want           helpers.Response
	}{
		{
			name:           "Success",
			articleService: smock.NewArticleService(t),
			mock: func(ms *smock.ArticleService) {
				expectedArticles := []*models.Article{
					{
						ID:          1,
						Title:       "title",
						Description: "description",
					},
				}

				ms.On("GetAll").Return(expectedArticles, nil)
			},
			want: helpers.Response{
				Data: []interface{}{
					map[string]interface{}{
						"ID":          float64(1),
						"Title":       "title",
						"Description": "description",
					},
				},
				Code:    200,
				Message: "OK",
				Error:   false,
			},
		},
		{
			name:           "Error Service.GetAll",
			articleService: smock.NewArticleService(t),
			mock: func(ms *smock.ArticleService) {
				ms.On("GetAll").Return(nil, errors.New("something error"))
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
				articleService: tt.articleService,
			}

			tt.mock(tt.articleService)

			r := gin.Default()
			endpoint := "/articles"

			r.GET(endpoint, h.GetAllArticlesHandler)

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

func TestHandler_GetArticleByIDHandler(t *testing.T) {
	tests := []struct {
		name           string
		articleService *smock.ArticleService
		mock           func(*smock.ArticleService)
		want           helpers.Response
	}{
		{
			name:           "Success",
			articleService: smock.NewArticleService(t),
			mock: func(ms *smock.ArticleService) {
				expectedArticle := &models.Article{
					ID:          1,
					Title:       "title",
					Description: "description",
				}

				ms.On("GetByID", mock.Anything).Return(expectedArticle, nil)
			},
			want: helpers.Response{
				Data: map[string]interface{}{
					"ID":          float64(1),
					"Title":       "title",
					"Description": "description",
				},
				Code:    200,
				Message: "OK",
				Error:   false,
			},
		},
		{
			name:           "Error Parse articleID",
			articleService: smock.NewArticleService(t),
			mock:           func(ms *smock.ArticleService) {},
			want: helpers.Response{
				Code:    400,
				Message: "strconv.ParseInt: parsing \"xx\": invalid syntax",
				Error:   true,
			},
		},
		{
			name:           "Error Service.ID record not found",
			articleService: smock.NewArticleService(t),
			mock: func(ms *smock.ArticleService) {
				ms.On("GetByID", mock.Anything).Return(nil, gorm.ErrRecordNotFound)
			},
			want: helpers.Response{
				Code:    404,
				Message: "record not found",
				Error:   true,
			},
		},
		{
			name:           "Error Service.GetByID",
			articleService: smock.NewArticleService(t),
			mock: func(ms *smock.ArticleService) {
				ms.On("GetByID", mock.Anything).Return(nil, errors.New("something error"))
			},
			want: helpers.Response{
				Code:    500,
				Message: "something error",
				Error:   true,
			},
		},
	}
	for _, tt := range tests {
		h := &Handler{
			articleService: tt.articleService,
		}

		tt.mock(tt.articleService)

		r := gin.Default()
		endpoint := "/articles"

		if tt.name == "Error Parse articleID" {
			r.GET(endpoint, MiddlewareMockID("xx"), h.GetArticleByIDHandler)
		} else {
			r.GET(endpoint, MiddlewareMockID("1"), h.GetArticleByIDHandler)
		}

		req, _ := http.NewRequest(http.MethodGet, endpoint, nil)

		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		var response helpers.Response
		err := json.Unmarshal(w.Body.Bytes(), &response)
		require.NoError(t, err)

		assert.Equal(t, tt.want.Code, w.Code)
		assert.Equal(t, tt.want, response)
	}
}

func TestHandler_CreateArticleHandler(t *testing.T) {
	tests := []struct {
		name           string
		body           io.Reader
		articleService *smock.ArticleService
		mock           func(*smock.ArticleService)
		want           helpers.Response
	}{
		{
			name: "Success",
			body: MakeRequestBody(dtos.CreateArticleRequestDTO{
				Title:       "title",
				Description: "description",
			}),
			articleService: smock.NewArticleService(t),
			mock: func(ms *smock.ArticleService) {
				expectedArticle := &models.Article{
					ID:          1,
					Title:       "title",
					Description: "description",
				}

				ms.On("Create", mock.Anything).Return(expectedArticle, nil)
			},
			want: helpers.Response{
				Data: map[string]interface{}{
					"ID":          float64(1),
					"Title":       "title",
					"Description": "description",
				},
				Code:    201,
				Message: "Created",
				Error:   false,
			},
		},
		{
			name: "Error ShouldBindJSON",
			body: MakeRequestBody(dtos.CreateArticleRequestDTO{
				Title: "title",
			}),
			articleService: smock.NewArticleService(t),
			mock:           func(ms *smock.ArticleService) {},
			want: helpers.Response{
				Code:    400,
				Message: "Key: 'CreateArticleRequestDTO.Description' Error:Field validation for 'Description' failed on the 'required' tag",
				Error:   true,
			},
		},
		{
			name: "Error Service.Create",
			body: MakeRequestBody(dtos.CreateArticleRequestDTO{
				Title:       "title",
				Description: "description",
			}),
			articleService: smock.NewArticleService(t),
			mock: func(ms *smock.ArticleService) {
				ms.On("Create", mock.Anything).Return(nil, errors.New("something error"))
			},
			want: helpers.Response{
				Code:    500,
				Message: "something error",
				Error:   true,
			},
		},
		{
			name: "Error copier",
			body: MakeRequestBody(dtos.CreateArticleRequestDTO{
				Title:       "title",
				Description: "description",
			}),
			articleService: smock.NewArticleService(t),
			mock: func(ms *smock.ArticleService) {
				ms.On("Create", mock.Anything).Return(nil, nil)
			},
			want: helpers.Response{
				Code:    500,
				Message: "copy from is invalid",
				Error:   true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &Handler{
				articleService: tt.articleService,
			}

			tt.mock(tt.articleService)

			r := gin.Default()
			endpoint := "/articles"

			r.POST(endpoint, h.CreateArticleHandler)

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

func TestHandler_UpdateArticleHandler(t *testing.T) {
	tests := []struct {
		name           string
		body           io.Reader
		articleService *smock.ArticleService
		mock           func(*smock.ArticleService)
		want           helpers.Response
	}{
		{
			name: "Success",
			body: MakeRequestBody(dtos.UpdateArticleRequestDTO{
				Title:       "title update",
				Description: "description update",
			}),
			articleService: smock.NewArticleService(t),
			mock: func(ms *smock.ArticleService) {
				expectedArticle := &models.Article{
					ID:          1,
					Title:       "title update",
					Description: "description update",
				}

				ms.On("Update", mock.Anything, mock.Anything).Return(expectedArticle, nil)
			},
			want: helpers.Response{
				Data: map[string]interface{}{
					"ID":          float64(1),
					"Title":       "title update",
					"Description": "description update",
				},
				Code:    200,
				Message: "OK",
				Error:   false,
			},
		},
		{
			name: "Error Parse articleID",
			body: MakeRequestBody(dtos.UpdateArticleRequestDTO{
				Title:       "title update",
				Description: "description update",
			}),
			articleService: smock.NewArticleService(t),
			mock:           func(ms *smock.ArticleService) {},
			want: helpers.Response{
				Code:    400,
				Message: "strconv.ParseInt: parsing \"xx\": invalid syntax",
				Error:   true,
			},
		},
		{
			name:           "Error ShouldBindJSON",
			body:           nil,
			articleService: smock.NewArticleService(t),
			mock:           func(ms *smock.ArticleService) {},
			want: helpers.Response{
				Code:    400,
				Message: "invalid request",
				Error:   true,
			},
		},
		{
			name: "Error Service.Update",
			body: MakeRequestBody(dtos.UpdateArticleRequestDTO{
				Title:       "title update",
				Description: "description update",
			}),
			articleService: smock.NewArticleService(t),
			mock: func(ms *smock.ArticleService) {
				ms.On("Update", mock.Anything, mock.Anything).Return(nil, errors.New("something error"))
			},
			want: helpers.Response{
				Code:    500,
				Message: "something error",
				Error:   true,
			},
		},
		{
			name: "Error Service.Update record not found",
			body: MakeRequestBody(dtos.UpdateArticleRequestDTO{
				Title:       "title update",
				Description: "description update",
			}),
			articleService: smock.NewArticleService(t),
			mock: func(ms *smock.ArticleService) {
				ms.On("Update", mock.Anything, mock.Anything).Return(nil, gorm.ErrRecordNotFound)
			},
			want: helpers.Response{
				Code:    404,
				Message: "record not found",
				Error:   true,
			},
		},
		{
			name: "Error copier",
			body: MakeRequestBody(dtos.UpdateArticleRequestDTO{
				Title:       "title update",
				Description: "description update",
			}),
			articleService: smock.NewArticleService(t),
			mock: func(ms *smock.ArticleService) {
				ms.On("Update", mock.Anything, mock.Anything).Return(nil, nil)
			},
			want: helpers.Response{
				Code:    500,
				Message: "copy from is invalid",
				Error:   true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &Handler{
				articleService: tt.articleService,
			}

			tt.mock(tt.articleService)

			r := gin.Default()
			endpoint := "/articles"

			if tt.name == "Error Parse articleID" {
				r.PUT(endpoint, MiddlewareMockID("xx"), h.UpdateArticleHandler)
			} else {
				r.PUT(endpoint, MiddlewareMockID("1"), h.UpdateArticleHandler)
			}

			req, _ := http.NewRequest(http.MethodPut, endpoint, tt.body)

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

func TestHandler_DeleteArticleHandler(t *testing.T) {
	tests := []struct {
		name           string
		articleService *smock.ArticleService
		mock           func(*smock.ArticleService)
		want           helpers.Response
	}{
		{
			name:           "Success",
			articleService: smock.NewArticleService(t),
			mock: func(ms *smock.ArticleService) {
				ms.On("Delete", mock.Anything).Return(nil)
			},
			want: helpers.Response{
				Data:    float64(1),
				Code:    200,
				Message: "OK",
				Error:   false,
			},
		},
		{
			name:           "Error Parse articleID",
			articleService: smock.NewArticleService(t),
			mock:           func(ms *smock.ArticleService) {},
			want: helpers.Response{
				Code:    400,
				Message: "strconv.ParseInt: parsing \"xx\": invalid syntax",
				Error:   true,
			},
		},
		{
			name:           "Error Service.Delete",
			articleService: smock.NewArticleService(t),
			mock: func(ms *smock.ArticleService) {
				ms.On("Delete", mock.Anything).Return(errors.New("something error"))
			},
			want: helpers.Response{
				Code:    500,
				Message: "something error",
				Error:   true,
			},
		},
		{
			name:           "Error Service.Delete record not found",
			articleService: smock.NewArticleService(t),
			mock: func(ms *smock.ArticleService) {
				ms.On("Delete", mock.Anything).Return(gorm.ErrRecordNotFound)
			},
			want: helpers.Response{
				Code:    404,
				Message: "record not found",
				Error:   true,
			},
		},
	}
	for _, tt := range tests {
		h := &Handler{
			articleService: tt.articleService,
		}

		tt.mock(tt.articleService)

		r := gin.Default()
		endpoint := "/articles"

		if tt.name == "Error Parse articleID" {
			r.DELETE(endpoint, MiddlewareMockID("xx"), h.DeleteArticleHandler)
		} else {
			r.DELETE(endpoint, MiddlewareMockID("1"), h.DeleteArticleHandler)
		}

		req, _ := http.NewRequest(http.MethodDelete, endpoint, nil)

		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		var response helpers.Response
		err := json.Unmarshal(w.Body.Bytes(), &response)
		require.NoError(t, err)

		assert.Equal(t, tt.want.Code, w.Code)
		assert.Equal(t, tt.want, response)
	}
}
