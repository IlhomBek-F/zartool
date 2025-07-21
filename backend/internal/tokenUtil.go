package internal

import (
	"time"
	"zartool/models"

	"github.com/golang-jwt/jwt/v4"
)

func GeneretaAccessToken(user models.Owner, secret string, expiry int) (token string, err error) {
	exp := time.Now().Add(time.Hour * time.Duration(expiry)).Unix()

	claims := models.JwtClaims{
		Id: int(user.ID),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: exp,
		},
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err = jwtToken.SignedString([]byte(secret))
	return token, err
}
