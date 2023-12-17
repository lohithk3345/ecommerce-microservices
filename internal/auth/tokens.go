package auth

import (
	"ecommerce/config"
	"ecommerce/constants"
	"ecommerce/types"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	Id   string     `json:"sub"`
	Role types.Role `json:"role"`
	jwt.StandardClaims
}

func GenerateAccessToken(id string, role types.Role) (types.Token, error) {
	claims := Claims{
		Id:   id,
		Role: role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour).Unix(),
		},
	}

	signed := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	token, err := signed.SignedString([]byte(config.EnvMap[constants.TOKEN_SECRET]))
	return token, err
}

func GenerateRefreshToken(id string, role types.Role) (types.Token, error) {
	claims := Claims{
		Id:   id,
		Role: role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(8760 * time.Hour).Unix(),
		},
	}

	signed := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	return signed.SignedString([]byte(config.EnvMap[constants.TOKEN_SECRET]))
}

func ValidateToken(tokenString types.Token) (*Claims, error) {
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
