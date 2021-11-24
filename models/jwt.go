package models

import "github.com/dgrijalva/jwt-go"

type MyClaims struct {
	Id   int64  `json:"id"`
	Role string `json:"role"`
	jwt.StandardClaims
}
