package apiHandlers

import (
	middleware "ecommerce/api/middlewares/user"
	"ecommerce/cache"
	"ecommerce/internal/auth"
	services "ecommerce/services/user"
	"ecommerce/types"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func SetupUserRouter(u *UserAPIHandlers) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	// router.Use()
	// router.GET("/")
	router.POST("/register", middleware.ApiKeyCheck(), u.registerUser)
	router.POST("/login", middleware.ApiKeyCheck(), u.login)
	router.GET("/protected", middleware.CheckAccessTokenAuth(), u.protected)
	router.GET("/refresh", middleware.CheckRefreshTokenAuth(), u.refresh)
	router.GET("/logout")
	return router
}

type UserAPIHandlers struct {
	service *services.UserServices
}

func NewUserApiHandler(database *mongo.Database) *UserAPIHandlers {
	return &UserAPIHandlers{services.NewUserService(database)}
}

func (u UserAPIHandlers) registerUser(ctx *gin.Context) {
	var user *types.UserRequest
	ctx.BindJSON(&user)

	newUser := user.Convert()
	hash, errPass := auth.HashPassword([]byte(user.Password))
	if errPass != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}
	newUser.AddHash(string(hash))

	_, err := u.service.CreateUser(newUser)
	if err != nil {
		ctx.JSON(409, gin.H{"error": "The Email Already Exists"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "User Created"})
	return
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (u UserAPIHandlers) login(ctx *gin.Context) {
	var data *LoginRequest
	ctx.BindJSON(&data)
	user, err := u.service.FindUserByFilter(data.Email)
	if err != nil {
		log.Println(err)
	}
	errAuth := auth.VerifyHash([]byte(user.Hash), data.Password)
	if errAuth != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "The Email Or Password Is Wrong. Please Check Again"})
		ctx.Abort()
		return
	}
	accessToken, errToken := auth.GenerateAccessToken(user.Id)
	if errToken != nil {
		ctx.Abort()
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}
	refreshToken, errRToken := auth.GenerateRefreshToken(user.Id)
	if errRToken != nil {
		ctx.Abort()
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}
	cache.RefreshTokenMap[refreshToken] = true
	ctx.Header("Authorization", "Bearer "+accessToken)
	ctx.JSON(http.StatusOK, gin.H{"accessToken": accessToken, "refreshToken": refreshToken})
}

// func logout() {}

func (u UserAPIHandlers) protected(ctx *gin.Context) {
	id, _ := ctx.Get("userId")
	log.Println(id)
	ctx.JSON(200, gin.H{"id": id})
	return
}

func (u UserAPIHandlers) refresh(ctx *gin.Context) {
	id, _ := ctx.Get("userId")
	log.Println(id)
	accessToken, errToken := auth.GenerateAccessToken(id.(string))
	if errToken != nil {
		ctx.Abort()
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}
	refreshToken, errRToken := auth.GenerateRefreshToken(id.(string))
	if errRToken != nil {
		ctx.Abort()
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	cache.RefreshTokenMap[refreshToken] = true
	ctx.Header("Authorization", "Bearer "+accessToken)
	ctx.JSON(http.StatusOK, gin.H{"accessToken": accessToken, "refreshToken": refreshToken})
}
