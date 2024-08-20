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

func TestNewArticleService(t *testing.T) {
	articleRepo := rmock.NewArticleRepository(t)
	articleService := NewArticleService(&ArticleServiceConfig{
		ArticleRepository: articleRepo,
	})

	assert.NotNil(t, articleService)
}

func Test_articleService_GetAll(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		tt := struct {
			want    []*models.Article
			wantErr bool
		}{
			want: []*models.Article{
				{
					ID:          1,
					Title:       "title",
					Description: "description",
				},
			},
			wantErr: false,
		}

		articleRepo := rmock.NewArticleRepository(t)
		s := &articleService{
			articleRepository: articleRepo,
		}

		result := []*models.Article{
			{
				ID:          1,
				Title:       "title",
				Description: "description",
			},
		}

		articleRepo.On("GetAll").Return(result, nil)

		got, err := s.GetAll()
		if (err != nil) != tt.wantErr {
			t.Errorf("articleService.GetAll() error = %v, wantErr %v", err, tt.wantErr)
			return
		}

		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("articleService.GetAll() = %v, want %v", got, tt.want)
		}
	})

	t.Run("Error Repository.GetAll", func(t *testing.T) {
		tt := struct {
			want    []*models.Article
			wantErr bool
		}{
			want:    nil,
			wantErr: true,
		}

		articleRepo := rmock.NewArticleRepository(t)
		s := &articleService{
			articleRepository: articleRepo,
		}

		articleRepo.On("GetAll").Return(nil, errors.New("something error"))

		got, err := s.GetAll()
		if (err != nil) != tt.wantErr {
			t.Errorf("articleService.GetAll() error = %v, wantErr %v", err, tt.wantErr)
			return
		}

		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("articleService.GetAll() = %v, want %v", got, tt.want)
		}
	})
}

func Test_articleService_GetByID(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		type args struct {
			id int64
		}
		tt := struct {
			args    args
			want    *models.Article
			wantErr bool
		}{
			args: args{
				id: 1,
			},
			want: &models.Article{
				ID:          1,
				Title:       "title",
				Description: "description",
			},
			wantErr: false,
		}

		articleRepo := rmock.NewArticleRepository(t)
		s := &articleService{
			articleRepository: articleRepo,
		}

		result := &models.Article{
			ID:          1,
			Title:       "title",
			Description: "description",
		}

		articleRepo.On("GetByID", mock.Anything).Return(result, nil)

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
			want    *models.Article
			wantErr bool
		}{
			args: args{
				id: 1,
			},
			want:    nil,
			wantErr: true,
		}

		articleRepo := rmock.NewArticleRepository(t)
		s := &articleService{
			articleRepository: articleRepo,
		}

		articleRepo.On("GetByID", mock.Anything).Return(nil, errors.New("something error"))

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

func Test_articleService_Create(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		type args struct {
			menu models.Article
		}
		tt := struct {
			args    args
			want    *models.Article
			wantErr bool
		}{
			args: args{
				menu: models.Article{
					Title:       "title",
					Description: "description",
				},
			},
			want: &models.Article{
				ID:          1,
				Title:       "title",
				Description: "description",
			},
			wantErr: false,
		}

		articleRepo := rmock.NewArticleRepository(t)
		s := &articleService{
			articleRepository: articleRepo,
		}

		result := &models.Article{
			ID:          1,
			Title:       "title",
			Description: "description",
		}

		articleRepo.On("Insert", mock.Anything).Return(result, nil)

		got, err := s.Create(tt.args.menu)
		if (err != nil) != tt.wantErr {
			t.Errorf("articleService.GetAll() error = %v, wantErr %v", err, tt.wantErr)
			return
		}

		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("articleService.GetAll() = %v, want %v", got, tt.want)
		}
	})

	t.Run("Error Repository.Insert", func(t *testing.T) {
		type args struct {
			menu models.Article
		}
		tt := struct {
			args    args
			want    *models.Article
			wantErr bool
		}{
			args: args{
				menu: models.Article{
					Title:       "title",
					Description: "description",
				},
			},
			want:    nil,
			wantErr: true,
		}

		articleRepo := rmock.NewArticleRepository(t)
		s := &articleService{
			articleRepository: articleRepo,
		}

		articleRepo.On("Insert", mock.Anything).Return(nil, errors.New("something error"))

		got, err := s.Create(tt.args.menu)
		if (err != nil) != tt.wantErr {
			t.Errorf("articleService.GetAll() error = %v, wantErr %v", err, tt.wantErr)
			return
		}

		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("articleService.GetAll() = %v, want %v", got, tt.want)
		}
	})
}

func Test_articleService_Update(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		type args struct {
			id   int64
			menu models.Article
		}
		tt := struct {
			args    args
			want    *models.Article
			wantErr bool
		}{
			args: args{
				id: 1,
				menu: models.Article{
					Title:       "title update",
					Description: "description update",
				},
			},
			want: &models.Article{
				ID:          1,
				Title:       "title update",
				Description: "description update",
			},
			wantErr: false,
		}

		articleRepo := rmock.NewArticleRepository(t)
		s := &articleService{
			articleRepository: articleRepo,
		}

		result := &models.Article{
			ID:          1,
			Title:       "title update",
			Description: "description update",
		}

		articleRepo.On("Update", mock.Anything, mock.Anything).Return(result, nil)

		got, err := s.Update(tt.args.id, tt.args.menu)
		if (err != nil) != tt.wantErr {
			t.Errorf("articleService.GetAll() error = %v, wantErr %v", err, tt.wantErr)
			return
		}

		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("articleService.GetAll() = %v, want %v", got, tt.want)
		}
	})

	t.Run("Error Repository.Update", func(t *testing.T) {
		type args struct {
			id   int64
			menu models.Article
		}
		tt := struct {
			args    args
			want    *models.Article
			wantErr bool
		}{
			args: args{
				id: 1,
				menu: models.Article{
					Title:       "title update",
					Description: "description update",
				},
			},
			want:    nil,
			wantErr: true,
		}

		articleRepo := rmock.NewArticleRepository(t)
		s := &articleService{
			articleRepository: articleRepo,
		}

		articleRepo.On("Update", mock.Anything, mock.Anything).Return(nil, errors.New("something error"))

		got, err := s.Update(tt.args.id, tt.args.menu)
		if (err != nil) != tt.wantErr {
			t.Errorf("articleService.GetAll() error = %v, wantErr %v", err, tt.wantErr)
			return
		}

		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("articleService.GetAll() = %v, want %v", got, tt.want)
		}
	})
}

func Test_articleService_Delete(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		type args struct {
			id int64
		}
		tt := struct {
			args    args
			wantErr bool
		}{
			args: args{
				id: 1,
			},
			wantErr: false,
		}

		articleRepo := rmock.NewArticleRepository(t)
		s := &articleService{
			articleRepository: articleRepo,
		}

		articleRepo.On("Delete", mock.Anything).Return(nil)

		if err := s.Delete(tt.args.id); (err != nil) != tt.wantErr {
			t.Errorf("articleService.GetAll() error = %v, wantErr %v", err, tt.wantErr)
			return
		}
	})

	t.Run("Error Repository.Delete", func(t *testing.T) {
		type args struct {
			id int64
		}
		tt := struct {
			args    args
			wantErr bool
		}{
			args: args{
				id: 1,
			},
			wantErr: true,
		}

		articleRepo := rmock.NewArticleRepository(t)
		s := &articleService{
			articleRepository: articleRepo,
		}

		articleRepo.On("Delete", mock.Anything).Return(errors.New("something error"))

		if err := s.Delete(tt.args.id); (err != nil) != tt.wantErr {
			t.Errorf("articleService.GetAll() error = %v, wantErr %v", err, tt.wantErr)
			return
		}
	})
}
