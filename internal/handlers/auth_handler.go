package handlers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/geshtng/go-base-backend/internal/dtos"
	errn "github.com/geshtng/go-base-backend/internal/errors"
	"github.com/geshtng/go-base-backend/internal/helpers"
	"github.com/geshtng/go-base-backend/internal/models"
)

func (h *Handler) RegisterHandler(c *gin.Context) {
	var request dtos.RegisterRequestDTO
	var response dtos.RegisterResponseDTO

	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, helpers.BuildErrorResponse(http.StatusBadRequest, err.Error()))
		return
	}

	auth := models.User{
		Username: request.Username,
		Password: request.Password,
	}

	newAuth, token, err := h.authService.Register(auth)
	if err != nil {
		c.JSON(http.StatusInternalServerError, helpers.BuildErrorResponse(http.StatusInternalServerError, err.Error()))
		return
	}

	response = dtos.RegisterResponseDTO{
		ID:       newAuth.ID,
		Username: newAuth.Username,
		Token:    *token,
	}

	c.JSON(http.StatusCreated, helpers.BuildSuccessResponse(http.StatusCreated, http.StatusText(http.StatusCreated), response))
}

func (h *Handler) LoginHandler(c *gin.Context) {
	var request dtos.LoginRequestDTO
	var response dtos.LoginResponseDTO

	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, helpers.BuildErrorResponse(http.StatusBadRequest, err.Error()))
		return
	}

	auth, token, err := h.authService.Login(request.Username, request.Password)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusUnauthorized, helpers.BuildErrorResponse(http.StatusUnauthorized, errn.ErrWrongUsernameOrPassword.Error()))
			return
		}

		c.JSON(http.StatusInternalServerError, helpers.BuildErrorResponse(http.StatusInternalServerError, err.Error()))
		return
	}

	response = dtos.LoginResponseDTO{
		ID:       auth.ID,
		Username: auth.Username,
		Token:    *token,
	}

	c.JSON(http.StatusOK, helpers.BuildSuccessResponse(http.StatusOK, http.StatusText(http.StatusOK), response))
}
