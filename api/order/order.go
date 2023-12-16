package orderHandler

import (
	middleware "ecommerce/api/middlewares/user"
	orderServices "ecommerce/services/order"
	"ecommerce/types"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func SetupOrderRouter(o *OrderApiHandler) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	router.Use(middleware.ApiKeyCheck(), middleware.CheckAccessTokenAuth())

	router.POST("/order", o.orderProduct)
	// router.POST("/order", o)
	router.GET("/order", o.getOrdersByUserId)
	// router.GET("/order/:id", middleware.CheckRefreshTokenAuth(), o)
	// router.GET("/logout", logout)

	return router
}

type OrderRequest struct {
	ProductId types.ProductID
	// Quantity  int32
}

type OrderApiHandler struct {
	service *orderServices.OrderServices
}

func NewOrderApiHandler(database *mongo.Database) *OrderApiHandler {
	return &OrderApiHandler{service: orderServices.NewOrderService(database)}
}

func (o OrderApiHandler) orderProduct(ctx *gin.Context) {
	userId := ctx.MustGet("userId").(types.UserID)
	var orderReq *OrderRequest
	ctx.BindJSON(&orderReq)
	log.Println(userId, orderReq)
	_, err := o.service.CreateOrder(userId, orderReq.ProductId)
	if err != nil {
		ctx.JSON(http.StatusNonAuthoritativeInfo, "Internal Server Error")
		return
	}

	ctx.JSON(http.StatusCreated, "Order Successful")
	return
}

func (o OrderApiHandler) getOrdersByUserId(ctx *gin.Context) {
	userId := ctx.MustGet("userId").(types.UserID)
	result, err := o.service.FindByUserId(userId)
	if err != nil {
		ctx.JSON(http.StatusNonAuthoritativeInfo, "Order Not Found")
		return
	}

	ctx.JSON(http.StatusOK, result)
	return
}
