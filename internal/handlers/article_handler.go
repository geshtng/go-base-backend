package handlers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"

	"github.com/geshtng/go-base-backend/internal/dtos"
	"github.com/geshtng/go-base-backend/internal/helpers"
	"github.com/geshtng/go-base-backend/internal/models"
)

func (h *Handler) GetAllArticles(c *gin.Context) {
	result, err := h.articleService.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, helpers.BuildErrorResponse(http.StatusInternalServerError, err.Error()))
		return
	}

	c.JSON(http.StatusOK, helpers.BuildSuccessResponse(http.StatusOK, http.StatusText(http.StatusOK), result))
}

func (h *Handler) GetArticleByID(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, helpers.BuildErrorResponse(http.StatusBadRequest, err.Error()))
		return
	}

	result, err := h.articleService.GetByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, helpers.BuildErrorResponse(http.StatusNotFound, err.Error()))
			return
		}

		c.JSON(http.StatusInternalServerError, helpers.BuildErrorResponse(http.StatusInternalServerError, err.Error()))
		return
	}

	c.JSON(http.StatusOK, helpers.BuildSuccessResponse(http.StatusOK, http.StatusText(http.StatusOK), result))
}

func (h *Handler) CreateArticle(c *gin.Context) {
	var request dtos.CreateArticleRequestDTO
	var response dtos.CreateArticleResponseDTO

	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, helpers.BuildErrorResponse(http.StatusBadRequest, err.Error()))
		return
	}

	article := models.Article{
		Title:       request.Title,
		Description: request.Description,
	}

	newArticle, err := h.articleService.Create(article)
	if err != nil {
		c.JSON(http.StatusInternalServerError, helpers.BuildErrorResponse(http.StatusInternalServerError, err.Error()))
		return
	}

	err = copier.Copy(&response, &newArticle)
	if err != nil {
		c.JSON(http.StatusInternalServerError, helpers.BuildErrorResponse(http.StatusInternalServerError, err.Error()))
		return
	}

	c.JSON(http.StatusCreated, helpers.BuildSuccessResponse(http.StatusCreated, http.StatusText(http.StatusCreated), newArticle))
}

func (h *Handler) UpdateArticle(c *gin.Context) {
	var request dtos.UpdateArticleRequestDTO
	var response dtos.UpdateArticleResponseDTO

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, helpers.BuildErrorResponse(http.StatusBadRequest, err.Error()))
		return
	}

	err = c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, helpers.BuildErrorResponse(http.StatusBadRequest, err.Error()))
		return
	}

	article := models.Article{
		Title:       request.Title,
		Description: request.Description,
	}

	updatedArticle, err := h.articleService.Update(id, article)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, helpers.BuildErrorResponse(http.StatusNotFound, err.Error()))
			return
		}

		c.JSON(http.StatusInternalServerError, helpers.BuildErrorResponse(http.StatusInternalServerError, err.Error()))
		return
	}

	err = copier.Copy(&response, &updatedArticle)
	if err != nil {
		c.JSON(http.StatusInternalServerError, helpers.BuildErrorResponse(http.StatusInternalServerError, err.Error()))
		return
	}

	c.JSON(http.StatusOK, helpers.BuildSuccessResponse(http.StatusOK, http.StatusText(http.StatusOK), updatedArticle))
}

func (h *Handler) DeleteArticle(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, helpers.BuildErrorResponse(http.StatusBadRequest, err.Error()))
		return
	}

	err = h.articleService.Delete(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, helpers.BuildErrorResponse(http.StatusNotFound, err.Error()))
			return
		}

		c.JSON(http.StatusInternalServerError, helpers.BuildErrorResponse(http.StatusInternalServerError, err.Error()))
		return
	}

	c.JSON(http.StatusOK, helpers.BuildSuccessResponse(http.StatusOK, http.StatusText(http.StatusOK), id))
}
