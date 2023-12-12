package apiHandlers

import (
	"ecommerce/config"
	"ecommerce/constants"
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
	router.POST("/register", u.registerUser)
	router.POST("/login")
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
	headers := ctx.Request.Header
	key := headers.Get(constants.API_KEY)
	if key == "" && key != config.EnvMap["API_KEY"] {
		log.Println("Not")
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	ctx.BindJSON(&user)
	newUser := user.Convert()
	_, err := u.service.CreateUser(newUser)
	if err != nil {
		ctx.JSON(409, gin.H{"error": "The Email Already Exists"})
		return
	}

	ctx.JSON(200, user)
	return
}
