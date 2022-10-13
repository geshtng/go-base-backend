package handlers

import (
	"github.com/geshtng/go-base-backend/db"
	r "github.com/geshtng/go-base-backend/internal/repositories"
	s "github.com/geshtng/go-base-backend/internal/services"
)

type Handler struct {
	articleService s.ArticleService
}

type HandlerConfig struct {
	ArticleService s.ArticleService
}

func New(c *HandlerConfig) *Handler {
	return &Handler{
		articleService: c.ArticleService,
	}
}

func InitAllHandlers() *Handler {
	db := db.Get()

	articleRepo := r.NewArticleRepository((&r.ArticleRepoConfig{DB: db}))

	articleService := s.NewArticleService(&s.ArticleServiceConfig{ArticleRepository: articleRepo})

	h := New(&HandlerConfig{
		ArticleService: articleService,
	})

	return h
}
