package middleware

import (
	"ecommerce/cache"
	"ecommerce/config"
	"ecommerce/constants"
	"ecommerce/internal/auth"
	"ecommerce/internal/helpers"
	"ecommerce/types"
	"log"
	"net/http"
	"slices"

	"github.com/gin-gonic/gin"
)

func ApiKeyCheck() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		key := ctx.GetHeader(constants.API_KEY)
		if key == types.EmptyString && key != config.EnvMap["API_KEY"] {
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
		ctx.Set("role", claims.Role)

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

func CheckIfDealerRole() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		iRole := ctx.MustGet("role")
		role := iRole.(types.Role)

		if role != constants.Dealer || !slices.Contains(constants.Roles, role) {
			log.Println("NOT A DEALER")
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			ctx.Abort()
			return
		}

		ctx.Next()
	}
}

func CheckIfCustomerRole() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		iRole := ctx.MustGet("role")
		role := iRole.(types.Role)

		if role != constants.Customer || !slices.Contains(constants.Roles, role) {
			log.Println("NOT A CUSTOMER")
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			ctx.Abort()
			return
		}

		ctx.Next()
	}
}
