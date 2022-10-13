package repositories

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"github.com/geshtng/go-base-backend/internal/models"
)

type ArticleRepository interface {
	GetAll() ([]*models.Article, error)
	GetByID(id int64) (*models.Article, error)
	Insert(article models.Article) (*models.Article, error)
	Update(id int64, article models.Article) (*models.Article, error)
	Delete(id int64) error
}

type articleRepository struct {
	db *gorm.DB
}

type ArticleRepoConfig struct {
	DB *gorm.DB
}

func NewArticleRepository(a *ArticleRepoConfig) ArticleRepository {
	return &articleRepository{
		db: a.DB,
	}
}

func (a *articleRepository) GetAll() ([]*models.Article, error) {
	var articles []*models.Article

	result := a.db.Find(&articles)
	if result.Error != nil {
		return nil, result.Error
	}

	return articles, nil
}

func (a *articleRepository) GetByID(id int64) (*models.Article, error) {
	var article *models.Article

	result := a.db.First(&article, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return article, nil
}

func (a *articleRepository) Insert(article models.Article) (*models.Article, error) {
	result := a.db.Create(&article)
	if result.Error != nil {
		return nil, result.Error
	}

	return &article, nil
}

func (a *articleRepository) Update(id int64, article models.Article) (*models.Article, error) {
	result := a.db.
		Clauses(clause.Returning{}).
		Where("id = ?", id).
		Updates(&article)

	if result.Error != nil {
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}

	return &article, nil
}

func (a *articleRepository) Delete(id int64) error {
	result := a.db.Delete(&models.Article{}, id)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}
