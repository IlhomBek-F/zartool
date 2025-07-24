package internal

import (
	"time"
	"zartool/domain"

	"github.com/golang-jwt/jwt/v4"
)

func GeneretaAccessToken(user domain.Owner, secret string, expiry int) (token string, err error) {
	exp := time.Now().Add(time.Hour * time.Duration(expiry)).Unix()

	claims := domain.JwtClaims{
		Id: int(user.ID),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: exp,
		},
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err = jwtToken.SignedString([]byte(secret))
	return token, err
}
