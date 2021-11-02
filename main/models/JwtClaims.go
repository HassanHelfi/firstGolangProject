package models

import "github.com/dgrijalva/jwt-go"

type JwtClaims struct {
	Name	string
	jwt.StandardClaims
}
