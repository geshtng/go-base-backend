package services

import (
	"github.com/geshtng/go-base-backend/internal/models"
	r "github.com/geshtng/go-base-backend/internal/repositories"
)

type ArticleService interface {
	GetAll() ([]*models.Article, error)
	GetByID(id int64) (*models.Article, error)
	Create(article models.Article) (*models.Article, error)
	Update(id int64, article models.Article) (*models.Article, error)
	Delete(id int64) error
}

type articleService struct {
	articleRepository r.ArticleRepository
}

type ArticleServiceConfig struct {
	ArticleRepository r.ArticleRepository
}

func NewArticleService(a *ArticleServiceConfig) ArticleService {
	return &articleService{
		articleRepository: a.ArticleRepository,
	}
}

func (a *articleService) GetAll() ([]*models.Article, error) {
	result, err := a.articleRepository.GetAll()
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (a *articleService) GetByID(id int64) (*models.Article, error) {
	result, err := a.articleRepository.GetByID(id)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (a *articleService) Create(article models.Article) (*models.Article, error) {
	result, err := a.articleRepository.Insert(article)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (a *articleService) Update(id int64, article models.Article) (*models.Article, error) {
	result, err := a.articleRepository.Update(id, article)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (a *articleService) Delete(id int64) error {
	err := a.articleRepository.Delete(id)
	if err != nil {
		return err
	}

	return nil
}
