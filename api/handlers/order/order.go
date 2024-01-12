package orderHandler

import (
	middleware "ecommerce/api/middlewares/user"
	"ecommerce/cache"
	"ecommerce/constants"
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
	router.DELETE("/order", o.cancelOrder)
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
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, "Dealer is not authorized to create order")
		return
	}

	ctx.JSON(http.StatusCreated, "Order Successful")
	return
}

func (o OrderApiHandler) getOrdersByUserId(ctx *gin.Context) {
	userId := ctx.MustGet("userId").(types.UserID)
	result, err := o.service.FindByUserId(userId)
	if err != nil {
		ctx.JSON(http.StatusNoContent, "Order Not Found")
		return
	}

	ctx.JSON(http.StatusOK, result)
	return
}

func (o OrderApiHandler) cancelOrder(ctx *gin.Context) {
	var order *types.Order
	ctx.BindJSON(&order)

	userId := ctx.MustGet("userId").(types.UserID)
	log.Println("OrderId", order.Id)

	result, err := o.service.FindById(order.Id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, "Order Not Found")
		return
	}

	if userId != result.UserId {
		ctx.JSON(http.StatusUnauthorized, "User Unauthorized")
		delete(cache.RefreshTokenMap, userId)
		ctx.Abort()
		return
	}

	if result.Status == constants.OrderCancelled {
		ctx.JSON(http.StatusBadRequest, "Already Cancelled")
		ctx.Abort()
		return
	}

	errC := o.service.CancelOneById(order.Id, result.ProductId)
	if errC != nil {
		ctx.JSON(http.StatusInternalServerError, "Internal Server Error")
		return
	}

	ctx.JSON(http.StatusOK, gin.H{})
	return
}
