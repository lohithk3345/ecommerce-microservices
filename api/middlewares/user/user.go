package middleware

import (
	"ecommerce/cache"
	"ecommerce/config"
	"ecommerce/constants"
	"ecommerce/internal/auth"
	"ecommerce/internal/helpers"
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
		token, err := helpers.H.GetToken(ctx)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Missing Token"})
			ctx.Abort()
			return
		}

		claims, err := auth.ValidateToken(token)
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
		token, err := helpers.H.GetToken(ctx)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Missing Token"})
			ctx.Abort()
			return
		}

		claims, err := auth.ValidateToken(token)
		if err != nil {
			log.Println(err)
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			ctx.Abort()
			return
		}

		if cache.RefreshTokenMap[claims.Id] != token {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			ctx.Abort()
			return
		}

		ctx.Set("userId", claims.Id)

		delete(cache.RefreshTokenMap, claims.Id)
		ctx.Next()
	}
}
