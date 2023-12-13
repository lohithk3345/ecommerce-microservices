package auth

import (
	"ecommerce/config"
	"ecommerce/constants"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	Id string `json:"sub"`
	jwt.StandardClaims
}

func GenerateAccessToken(id string) (string, error) {
	claims := Claims{
		Id: id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(5 * time.Minute).Unix(),
		},
	}

	signed := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	return signed.SignedString([]byte(config.EnvMap[constants.TOKEN_SECRET]))
}

func GenerateRefreshToken(id string) (string, error) {
	claims := Claims{
		Id: id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(8760 * time.Hour).Unix(),
		},
	}

	signed := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	return signed.SignedString([]byte(config.EnvMap[constants.TOKEN_SECRET]))
}

func ValidateToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.EnvMap[constants.TOKEN_SECRET]), nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return claims, nil
}
