package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/geshtng/go-base-backend/internal/handlers"
)

func InitAllRoutes(r *gin.Engine, h *handlers.Handler) {
	// Articles
	r.GET("/articles", h.GetAllArticles)
	r.GET("/articles/:id", h.GetArticleByID)
	r.POST("/articles", h.CreateArticle)
	r.PUT("/articles/:id", h.UpdateArticle)
	r.DELETE("/articles/:id", h.DeleteArticle)
}
