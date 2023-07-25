package helpers

import (
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v4"

	"github.com/geshtng/go-base-backend/config"
)

type IdTokenClaims struct {
	jwt.RegisteredClaims
	Data interface{} `json:"data"`
}

func GenerateJwtToken(data interface{}) (*string, error) {
	var idExp int64

	jwtEnv := config.InitJwt()

	expired, _ := strconv.ParseInt(jwtEnv.Expired, 10, 64)

	idExp = expired * 60
	unixTime := time.Now().Unix()
	tokenExp := unixTime + idExp

	claims := &IdTokenClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer: jwtEnv.Issuer,
			ExpiresAt: &jwt.NumericDate{
				Time: time.Unix(tokenExp, 0),
			},
			IssuedAt: &jwt.NumericDate{
				Time: time.Now(),
			},
		},
		Data: data,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(jwtEnv.Secret))
	if err != nil {
		return nil, err
	}

	return &tokenString, nil
}
