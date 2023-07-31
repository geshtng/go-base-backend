package handlers

import (
	"github.com/geshtng/go-base-backend/db"
	"google.golang.org/grpc"

	article_grpc "github.com/geshtng/go-base-backend/internal/handlers/article"
	r "github.com/geshtng/go-base-backend/internal/repositories"
	s "github.com/geshtng/go-base-backend/internal/services"
)

func InitAllHandlers(server *grpc.Server) {
	db := db.Get()

	articleRepo := r.NewArticleRepository((&r.ArticleRepoConfig{DB: db}))

	articleService := s.NewArticleService(&s.ArticleServiceConfig{ArticleRepository: articleRepo})

	article_grpc.NewArticleServerGrpc(server, articleService)
}
