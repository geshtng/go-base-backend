package middlewares

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"

	"github.com/geshtng/go-base-backend/config"
	"github.com/geshtng/go-base-backend/internal/dtos"
	"github.com/geshtng/go-base-backend/internal/helpers"
)

func validateToken(encodedToken string) (*jwt.Token, error) {
	jwtEnv := config.InitJwt()

	return jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid token")
		}

		return []byte(jwtEnv[2]), nil
	})
}

func AuthorizeJWT(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")

	if authHeader == "" {
		c.AbortWithStatusJSON(
			http.StatusUnauthorized,
			helpers.BuildErrorResponse(http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized)),
		)
		return
	}

	tokenString := authHeader[7:]

	token, err := validateToken(tokenString)
	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusUnauthorized,
			helpers.BuildErrorResponse(http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized)),
		)
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		c.AbortWithStatusJSON(
			http.StatusUnauthorized,
			helpers.BuildErrorResponse(http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized)),
		)
		return
	}

	userJson, _ := json.Marshal(claims["data"]) // change from type map[string]interface{} to []byte
	jwtData := dtos.JwtData{}
	err = json.Unmarshal(userJson, &jwtData) // change from type []byte to models.User{}
	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusUnauthorized,
			helpers.BuildErrorResponse(http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized)),
		)
		return
	}

	c.Set("user", jwtData)
}
