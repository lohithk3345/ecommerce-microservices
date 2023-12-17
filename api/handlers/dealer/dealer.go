package dealerHandlers

// import (
// 	middleware "ecommerce/api/middlewares/user"
// 	"ecommerce/cache"
// 	"ecommerce/internal/auth"
// 	"ecommerce/internal/helpers"
// 	dealerService "ecommerce/services/dealer"
// 	"ecommerce/types"
// 	"log"
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// 	"go.mongodb.org/mongo-driver/mongo"
// )

// func SetupUserRouter(u *UserAPIHandlers) *gin.Engine {
// 	gin.SetMode(gin.ReleaseMode)
// 	router := gin.Default()

// 	router.Use(middleware.ApiKeyCheck())

// 	router.POST("/register", u.registerUser)
// 	router.POST("/login", u.login)
// 	router.GET("/protected", middleware.CheckAccessTokenAuth(), u.protected)
// 	router.GET("/refresh", middleware.CheckRefreshTokenAuth(), u.refresh)
// 	router.GET("/logout", logout)

// 	return router
// }

// type UserAPIHandlers struct {
// 	service *dealerService.DealerServices
// }

// func NewUserApiHandler(database *mongo.Database) *UserAPIHandlers {
// 	return &UserAPIHandlers{service: dealerService.NewDealerService(database)}
// }

// func (u UserAPIHandlers) registerUser(ctx *gin.Context) {
// 	var user *types.UserRequest
// 	ctx.BindJSON(&user)

// 	newUser := user.Convert()
// 	hash, errPass := auth.HashPassword([]byte(user.Password))
// 	if errPass != nil {
// 		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
// 		return
// 	}
// 	newUser.AddHash(string(hash))

// 	_, err := u.service.CreateDealer(newUser)
// 	if err != nil {
// 		ctx.JSON(409, gin.H{"error": "The Email Already Exists"})
// 		return
// 	}

// 	ctx.JSON(http.StatusCreated, gin.H{"message": "User Created"})
// 	return
// }

// type LoginRequest struct {
// 	Email    string `json:"email"`
// 	Password string `json:"password"`
// }

// func (u UserAPIHandlers) login(ctx *gin.Context) {
// 	var data *LoginRequest
// 	ctx.BindJSON(&data)
// 	user, err := u.service.FindUserByFilter(data.Email)
// 	if err != nil {
// 		log.Println(err)
// 	}
// 	errAuth := auth.VerifyHash([]byte(user.Hash), data.Password)
// 	if errAuth != nil {
// 		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "The Email Or Password Is Wrong. Please Check Again"})
// 		ctx.Abort()
// 		return
// 	}
// 	accessToken, errToken := auth.GenerateAccessToken(user.Id)
// 	if errToken != nil {
// 		ctx.Abort()
// 		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
// 		return
// 	}
// 	refreshToken, errRToken := auth.GenerateRefreshToken(user.Id)
// 	if errRToken != nil {
// 		ctx.Abort()
// 		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
// 		return
// 	}
// 	cache.RefreshTokenMap[user.Id] = refreshToken
// 	ctx.Header("Authorization", "Bearer "+accessToken)
// 	ctx.JSON(http.StatusOK, gin.H{"accessToken": accessToken, "refreshToken": refreshToken})
// }

// func logout(ctx *gin.Context) {
// 	token, err := helpers.H.GetToken(ctx)
// 	if err != nil {
// 		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Missing Token"})
// 		ctx.Abort()
// 		return
// 	}

// 	claims, err := auth.ValidateToken(token)
// 	if err != nil {
// 		log.Println(err)
// 		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
// 		ctx.Abort()
// 		return
// 	}
// 	delete(cache.RefreshTokenMap, claims.Id)
// }

// func (u UserAPIHandlers) protected(ctx *gin.Context) {
// 	id, _ := ctx.Get("userId")
// 	ctx.JSON(200, gin.H{"id": id})
// 	return
// }

// func (u UserAPIHandlers) refresh(ctx *gin.Context) {
// 	id := ctx.MustGet("userId").(types.ID)
// 	log.Println(id)
// 	accessToken, errToken := auth.GenerateAccessToken(id)
// 	if errToken != nil {
// 		ctx.Abort()
// 		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
// 		return
// 	}
// 	refreshToken, errRToken := auth.GenerateRefreshToken(id)
// 	if errRToken != nil {
// 		ctx.Abort()
// 		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
// 		return
// 	}

// 	cache.RefreshTokenMap[id] = refreshToken
// 	ctx.Header("Authorization", "Bearer "+accessToken)
// 	ctx.JSON(http.StatusOK, gin.H{"accessToken": accessToken, "refreshToken": refreshToken})
// }
