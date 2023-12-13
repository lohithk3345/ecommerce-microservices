package middleware

import (
	"ecommerce/cache"
	"ecommerce/config"
	"ecommerce/constants"
	"ecommerce/internal/auth"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ApiKeyCheck() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		key := ctx.GetHeader(constants.API_KEY)
		if key == "" && key != config.EnvMap["API_KEY"] {
			log.Println("Not")
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			ctx.Abort()
			return
		}

		ctx.Next()
	}
}

func CheckAccessTokenAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("Authorization")
		log.Println(token)
		bearerFix := "Bearer "
		if token == "" || len(token) < len(bearerFix) || token[:len(bearerFix)] != bearerFix {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Missing Token"})
			ctx.Abort()
			return
		}
		actualToken := token[len(bearerFix):]

		claims, err := auth.ValidateToken(actualToken)
		if err != nil {
			log.Println(err)
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			ctx.Abort()
			return
		}

		ctx.Set("userId", claims.Id)

		ctx.Next()
	}
}

func CheckRefreshTokenAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("Authorization")
		log.Println(token)
		bearerFix := "Bearer "
		if token == "" || len(token) < len(bearerFix) || token[:len(bearerFix)] != bearerFix {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Missing Token"})
			ctx.Abort()
			return
		}
		actualToken := token[len(bearerFix):]

		if !cache.RefreshTokenMap[actualToken] {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			ctx.Abort()
			return
		}

		claims, err := auth.ValidateToken(actualToken)
		if err != nil {
			log.Println(err)
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			ctx.Abort()
			return
		}

		ctx.Set("userId", claims.Id)

		delete(cache.RefreshTokenMap, actualToken)
		ctx.Next()
	}
}
