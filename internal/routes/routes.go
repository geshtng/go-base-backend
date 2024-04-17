package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/geshtng/go-base-backend/internal/handlers"
	"github.com/geshtng/go-base-backend/internal/middlewares"
)

func InitAllRoutes(r *gin.Engine, h *handlers.Handler) {
	// Articles
	r.GET("/articles", h.GetAllArticlesHandler)
	r.GET("/articles/:id", h.GetArticleByIDHandler)
	r.POST("/articles", h.CreateArticleHandler)
	r.PUT("/articles/:id", h.UpdateArticleHandler)
	r.DELETE("/articles/:id", h.DeleteArticleHandler)

	// Auth
	r.POST("/login", h.LoginHandler)
	r.POST("/register", h.RegisterHandler)

	// API endpoints below will use authorization
	r.Use(middlewares.AuthorizeJWT)

	r.GET("/profiles", h.ShowProfileHandler)
}
