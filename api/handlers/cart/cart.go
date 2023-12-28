package cartHandler

import (
	middleware "ecommerce/api/middlewares/user"
	"ecommerce/cache"
	cartServices "ecommerce/services/cart"
	"ecommerce/types"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func SetupCartRouter(o *CartApiHandler) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	router.Use(middleware.ApiKeyCheck(), middleware.CheckAccessTokenAuth())

	router.POST("/cart", o.addToCart)
	router.DELETE("/cart", o.removeFromCart)
	router.GET("/cart", o.getCartByUserId)
	// router.GET("/cart/:id", middleware.CheckRefreshTokenAuth(), o)
	// router.GET("/logout", logout)

	return router
}

type CartRequest struct {
	ProductId types.ProductID
}

type RemoveItemCartRequest struct {
	CartId types.CartID
}

type CartApiHandler struct {
	service *cartServices.CartServices
}

func NewCartApiHandler(database *mongo.Database) *CartApiHandler {
	return &CartApiHandler{service: cartServices.NewCartService(database)}
}

func (o CartApiHandler) addToCart(ctx *gin.Context) {
	userId := ctx.MustGet("userId").(types.UserID)
	var cartReq *CartRequest
	ctx.BindJSON(&cartReq)
	log.Println(userId, cartReq)
	_, err := o.service.AddToCart(userId, cartReq.ProductId)
	if err != nil {
		ctx.JSON(http.StatusNonAuthoritativeInfo, gin.H{"Status": "Internal Server Error"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"Status": "Added To Cart"})
	return
}

func (o CartApiHandler) getCartByUserId(ctx *gin.Context) {
	userId := ctx.MustGet("userId").(types.UserID)
	result, err := o.service.FindByUserId(userId)
	if err != nil {
		ctx.JSON(http.StatusNonAuthoritativeInfo, gin.H{"Status": "Item Not Found"})
		return
	}

	if result == nil {
		ctx.JSON(http.StatusOK, gin.H{"Status": "No Items In Your Cart"})
		return
	}

	ctx.JSON(http.StatusOK, result)
	return
}

func (o CartApiHandler) removeFromCart(ctx *gin.Context) {
	var cart *RemoveItemCartRequest
	ctx.BindJSON(&cart)

	userId := ctx.MustGet("userId").(types.UserID)
	log.Println("CartId", cart.CartId)

	result, err := o.service.FindById(cart.CartId)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"Status": "Item Not Found"})
		return
	}

	if userId != result.UserId {
		ctx.JSON(http.StatusUnauthorized, gin.H{"Status": "User Unauthorized"})
		delete(cache.RefreshTokenMap, userId)
		ctx.Abort()
		return
	}

	errI := o.service.RemoveItemByCartId(cart.CartId)
	if errI != nil {
		ctx.JSON(http.StatusNonAuthoritativeInfo, gin.H{"Status": "Internal Server Error"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"Status": "Removed From Cart"})
	return
}
