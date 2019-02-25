package utils

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/go-chi/jwtauth"
)

var TokenAuth *jwtauth.JWTAuth

func init() {
	TokenAuth = jwtauth.New("HS256", []byte("secretkeyforjwt"), nil)
}

func GenerateUserToken(jwtClaim jwt.MapClaims) (string, error) {
	_, tokenString, _ := TokenAuth.Encode(jwtClaim)
	return tokenString, nil
}
