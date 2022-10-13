package handlers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"

	"github.com/geshtng/go-base-backend/internal/dtos"
	"github.com/geshtng/go-base-backend/internal/helpers"
)

func (h *Handler) ShowProfileHandler(c *gin.Context) {
	userContext := c.MustGet("user").(dtos.JwtData)

	var response dtos.ShowProfileResponseDTO

	user, err := h.userService.GetByID(userContext.ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, helpers.BuildErrorResponse(http.StatusNotFound, err.Error()))
			return
		}

		c.JSON(http.StatusInternalServerError, helpers.BuildErrorResponse(http.StatusInternalServerError, err.Error()))
		return
	}

	err = copier.Copy(&response, &user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, helpers.BuildErrorResponse(http.StatusInternalServerError, err.Error()))
		return
	}

	c.JSON(http.StatusOK, helpers.BuildSuccessResponse(http.StatusOK, http.StatusText(http.StatusOK), response))
}
