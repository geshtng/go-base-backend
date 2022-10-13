package handlers

import (
	"github.com/geshtng/go-base-backend/db"
	r "github.com/geshtng/go-base-backend/internal/repositories"
	s "github.com/geshtng/go-base-backend/internal/services"
)

type Handler struct {
	articleService s.ArticleService
	authService    s.AuthService
	userService    s.UserService
}

type HandlerConfig struct {
	ArticleService s.ArticleService
	AuthService    s.AuthService
	UserService    s.UserService
}

func New(c *HandlerConfig) *Handler {
	return &Handler{
		articleService: c.ArticleService,
		authService:    c.AuthService,
		userService:    c.UserService,
	}
}

func InitAllHandlers() *Handler {
	db := db.Get()

	articleRepo := r.NewArticleRepository((&r.ArticleRepoConfig{DB: db}))
	userRepo := r.NewUserRepository((&r.UserRepoConfig{DB: db}))

	articleService := s.NewArticleService(&s.ArticleServiceConfig{ArticleRepository: articleRepo})
	authService := s.NewAuthService(&s.AuthServiceConfig{UserRepository: userRepo})
	userService := s.NewUserService(&s.UserServiceConfig{UserRepository: userRepo})

	h := New(&HandlerConfig{
		ArticleService: articleService,
		AuthService:    authService,
		UserService:    userService,
	})

	return h
}
