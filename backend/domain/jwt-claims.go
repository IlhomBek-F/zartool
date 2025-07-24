package domain

import "github.com/golang-jwt/jwt/v4"

type JwtClaims struct {
	Id int `json:"id"`
	jwt.StandardClaims
}
